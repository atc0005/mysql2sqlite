// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

//go:generate go-winres make --product-version=git-tag --file-version=git-tag

package main

import (
	"context"
	"database/sql"
	"fmt"
	"path"
	"strings"

	"github.com/atc0005/mysql2sqlite/internal/caller"
	"github.com/atc0005/mysql2sqlite/internal/config"
	"github.com/atc0005/mysql2sqlite/internal/dbqs"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/apex/log"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info(config.Branding())

	log.Infof("Successfully loaded configuration file %s", cfg.ConfigFileUsed())

	// Create context that can be used to cancel background jobs.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mysqlDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?tls=%s",
		cfg.MySQLUsername(),
		cfg.MySQLPassword(),
		cfg.MySQLHost(),
		cfg.MySQLPort(),
		cfg.MySQLDatabase(),
		cfg.MySQLEncryption(),
	)

	mysqlDB, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		log.Error(err.Error())
		return
	}

	defer func(db string, host string) {
		if err := mysqlDB.Close(); err != nil {
			log.Errorf(
				"error closing connection to database %s on host %s: %v",
				db,
				host,
				err,
			)
		} else {
			log.Infof(
				"Successfully closed connection to database %s on host %s",
				db,
				host,
			)
		}
	}(cfg.MySQLDatabase(), cfg.MySQLHost())

	// https://github.com/go-sql-driver/mysql#important-settings
	mysqlDB.SetConnMaxLifetime(cfg.MySQLConnMaxLifetime())
	mysqlDB.SetMaxOpenConns(cfg.MySQLMaxOpenConns())
	mysqlDB.SetMaxIdleConns(cfg.MySQLMaxIdleConns())
	mysqlDB.SetConnMaxIdleTime(cfg.MySQLConnMaxIdleTime()) // Go 1.15+

	// test MySQL database connection before proceeding further
	if err = dbqs.VerifyDBConn(
		ctx,
		mysqlDB,
		cfg.ConnectionRetries(),
		cfg.ConnectionRetryDelay(),
		cfg.ConnectionTimeout(),
	); err != nil {
		log.Errorf(
			"verifying MySQL database connection failed: %v",
			err,
		)
		return
	}
	log.Infof(
		"Connection established to source database server at %s:%d",
		cfg.MySQLHost(),
		cfg.MySQLPort(),
	)

	sqliteDBFile := path.Join(cfg.SQLiteDBPath(), cfg.SQLiteDBFile())

	sqliteDSN := fmt.Sprintf(
		"%s?_journal_mode=%s&_busy_timeout=%d",
		sqliteDBFile,
		cfg.SQLiteJournalMode(),
		cfg.SQLiteBusyTimeout(),
	)
	sqliteDB, err := sql.Open("sqlite3", sqliteDSN)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer func(db string) {
		if err := sqliteDB.Close(); err != nil {
			log.Errorf("error closing SQLite database %s: %v", db, err)
		} else {
			log.Infof("Successfully closed connection to SQLite database %s", db)
		}
	}(sqliteDBFile)

	// Help prevent locked database errors. May also require a shared cache.
	// https://github.com/mattn/go-sqlite3#faq
	//
	// FIXME: This seems to halt the application mid-stride; something is
	// causing this app to lock up when using this setting. Based on brief
	// testing, this appears to occur when the prepared statement in the
	// dbqs.RowsCount() function is used. The workaround (if this is desired),
	// is to use a sql.Conn to force queries to run across a single
	// connection.
	//
	// sqliteDB.SetMaxOpenConns(1)

	logConnStats := func() {
		if cfg.LogDBStats() {
			log.Info(caller.GetParentFuncFileLineInfo())
			log.Infof("SQLite connection stats: %+v", sqliteDB.Stats())
			log.Infof("MySQL connection stats: %+v", mysqlDB.Stats())
		}
	}

	// test SQLite database connection before proceeding further
	if err = dbqs.VerifyDBConn(
		ctx,
		sqliteDB,
		cfg.ConnectionRetries(),
		cfg.ConnectionRetryDelay(),
		cfg.ConnectionTimeout(),
	); err != nil {
		log.Errorf(
			"verifying SQLite database connection failed: %v",
			err,
		)
		return
	}
	log.Infof("Connection established to SQLite database at %s", sqliteDBFile)
	logConnStats()

	// Start a transaction so we have SQLite database rollback potential
	// if the process is interrupted.
	// TODO: Do we need any specific transaction settings, or continue to
	// pass nil here?
	sqliteTX, sqliteTXErr := sqliteDB.BeginTx(ctx, nil)
	if sqliteTXErr != nil {
		log.Errorf(
			"failed to start SQL transaction for SQLite database: %v",
			sqliteTXErr,
		)
		return
	}

	// TODO: Move these elsewhere?
	rollbackTX := func(tx *sql.Tx) {
		if txRollbackErr := tx.Rollback(); txRollbackErr != nil {
			log.Errorf(
				"failed to roll back transaction: %v",
				txRollbackErr,
			)
			return
		}
		log.Warn("Transaction rollback complete")
	}

	for table, querySet := range cfg.DBQueries() {

		// log.Warn("Fake delay in order to allow for destruction")
		// time.Sleep(time.Second * 10)

		log.Debugf(
			"reading table %s from source database %s",
			table,
			cfg.MySQLDatabase(),
		)

		rowsCount, rowsCountErr := dbqs.RowsCount(mysqlDB, table)
		if rowsCountErr != nil {
			rollbackTX(sqliteTX)
			log.Error(rowsCountErr.Error())
			return
		}
		logConnStats()

		// Prepare SQLite database, regardless of whether we have any rows
		// to copy from source database table.
		if err := dbqs.PrepareSQLiteDB(
			sqliteTX,
			querySet,
			table,
			cfg.SQLiteCreateIndexes(),
		); err != nil {
			rollbackTX(sqliteTX)
			log.Error(err.Error())
		}

		if rowsCount == 0 {
			log.Infof(
				"no entries in table %s for database %s, skipping rows retrieval and SQLite database insert",
				table,
				cfg.MySQLDatabase(),
			)
			continue
		}

		log.Debugf("About to execute query to retrieve %d existing rows", rowsCount)
		mysqlRows, err := mysqlDB.Query(querySet[dbqs.SQLQueriesRead])
		if err != nil {
			rollbackTX(sqliteTX)
			log.Error(err.Error())
			return
		}
		log.Debug("Rows retrieved from source MySQL database")
		logConnStats()

		columnNames, err := mysqlRows.Columns()
		if err != nil {
			rollbackTX(sqliteTX)
			log.Errorf(
				"unable to fetch column names for table %s: %v",
				table,
				err,
			)
			return
		}

		colCount := len(columnNames)

		// The `valuePrts` slice is used to satisfy the `(*sql.Rows).Scan`
		// method requiring that arguments be specified as pointers to a
		// value.
		//
		// We create pointers in valuePtrs which directly point to the value
		// in the "sibling" `values` slice. For this to work, we intentionally
		// *do not* specify a length of zero; we need to prefill the slice
		// entries with applicable zero values that we can index into and
		// update later.
		values := make([]interface{}, colCount)
		valuePtrs := make([]interface{}, colCount)

		log.Debug("Before mysqlRows.Next()")

		log.Debugf(
			"%d rows to be copied from table %s to SQLite database file %s",
			rowsCount,
			table,
			sqliteDBFile,
		)

		for mysqlRows.Next() {
			logConnStats()

			// Initialize slice of pointers in order to satisfy the
			// `(*sql.Rows).Scan` method.
			for i := range columnNames {
				valuePtrs[i] = &values[i]
			}

			if err = mysqlRows.Scan(valuePtrs...); err != nil {
				rollbackTX(sqliteTX)
				log.Error(err.Error())
				return
			}
			log.Debug("Completed scanning from MySQL row, before SQLite insert")
			logConnStats()

			// The db.Exec method requires a slice of interface{}, but in
			// order to force SQLite to *not* store the MySQL data as
			// "blob" type we force a conversion to string before storing
			// in a slice of interface{} for db.Exec to use.
			//
			// If the user requested it, we also trim whitespace from the
			// source MySQL data.
			sqliteOutput := make([]interface{}, len(values))
			for idx := range values {
				byteArray, ok := values[idx].([]byte)
				if !ok {
					log.Debugf("Field %q does NOT contain byte slice value", columnNames[idx])
					log.Debugf("Saving field %q value to sqliteOutput slice as-is", columnNames[idx])
					sqliteOutput[idx] = values[idx]
				} else {
					log.Debugf("Field %q does contain byte slice value", columnNames[idx])
					log.Debugf("Converting field %q value to sqliteOutput as string value", columnNames[idx])
					sqliteOutput[idx] = string(byteArray)

					if cfg.TrimWhitespace() {
						log.Debugf("Trimming whitespace from field %q", columnNames[idx])

						// WARNING: logging field values could potentially
						// expose sensitive data (e.g., passwords)
						//
						// log.Debugf("value is %q before trim", values[idx])
						sqliteOutput[idx] = strings.TrimSpace(string(byteArray))
						// log.Debugf("value is %q after trim", values[idx])
					}
				}

				// Rely on fmt.Sprintf to handle formatting, potential
				// conversion of data to string type before storing back
				// in our []interface{} for use by db.Exec
				sqliteOutput[idx] = fmt.Sprintf("%v", sqliteOutput[idx])

			}

			// Populate SQLite database with row retrieved from MySQL
			if _, err = sqliteTX.Exec(querySet[dbqs.SQLQueriesWrite], sqliteOutput...); err != nil {
				rollbackTX(sqliteTX)
				log.Error(err.Error())
				return
			}
			log.Debug("No errors encountered from SQLite database insert")
			logConnStats()
		}

		// check for potential errors after processing rows
		if err := mysqlRows.Err(); err != nil {
			// if there is an error reading from MySQL, rollback SQLite db
			// changes as there may be some critical data that we're
			// missing
			rollbackTX(sqliteTX)
			log.Error(err.Error())
			return
		}
		logConnStats()

		log.Infof(
			"Successfully copied %d rows from source database table %s",
			rowsCount,
			table,
		)

		// explicit close here vs defer since we're in a loop and do not want
		// to persist the database connection for longer than necessary
		if err := mysqlRows.Close(); err != nil {
			rollbackTX(sqliteTX)
			log.Errorf(
				"error closing mysqlRows for table %s: %v",
				table,
				err,
			)
			return
		}
		log.Debug("Successfully closed mysqlRows")
		logConnStats()

	}

	if err := sqliteTX.Commit(); err != nil {
		// FIXME: This is a pretty serious failure; what else can we
		// do at this point?
		log.Errorf(
			"failed to commit transaction for SQLite database: %v",
			err,
		)
		return
	}
	log.Info("Successfully completed SQLite database transaction commit")

	log.Debug("No errors encountered reading data from MySQL")
	log.Infof(
		"Finished syncing database %s from host %s:%d to SQLite database %s",
		cfg.MySQLDatabase(),
		cfg.MySQLHost(),
		cfg.MySQLPort(),
		sqliteDBFile,
	)

}
