// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

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
		log.Fatal(err.Error())
	}

	log.Info(config.Branding())

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
		log.Fatal(err.Error())
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

	// Disabled for now in order to support Go 1.14. See GH-28 for details.
	//mysqlDB.SetConnMaxIdleTime(cfg.MySQLConnMaxIdleTime()) // Go 1.15+

	// test MySQL database connection before proceeding further
	if err = dbqs.VerifyDBConn(
		ctx,
		mysqlDB,
		cfg.ConnectionRetries(),
		cfg.ConnectionRetryDelay(),
		cfg.ConnectionTimeout(),
	); err != nil {
		log.Fatalf(
			"verifying MySQL database connection failed: %v",
			err,
		)
	}
	log.Infof(
		"Connection established to source database server at %s:%d",
		cfg.MySQLHost(),
		cfg.MySQLPort(),
	)

	sqliteDBFile := path.Join(cfg.SQLiteDBPath(), cfg.SQLiteDBFile())

	sqliteDSN := fmt.Sprintf(
		// Enable Write-Ahead Logging (WAL) in an effort to avoid EXCLUSIVE
		// locks on the database which block Postfix from reading the database
		// while this application does its work. Based on light reading, this
		// *also* appears to boost the performance of read-only operations, so
		// we use it for that purpose as well.
		//
		// The DELETE journaling mode is the normal behavior. In the DELETE
		// mode, the rollback journal is deleted at the conclusion of each
		// transaction. Indeed, the delete operation is the action that causes
		// the transaction to commit.
		//
		// https://dba.stackexchange.com/questions/45368/how-do-i-prevent-sqlite-database-locks
		// https://manski.net/2012/10/sqlite-performance/
		// https://www.sqlite.org/wal.html
		// https://www.sqlite.org/pragma.html#pragma_journal_mode
		// https://www.sqlite.org/pragma.html#pragma_busy_timeout
		// https://github.com/mattn/go-sqlite3/pull/827
		// https://github.com/mattn/go-sqlite3/issues/209
		//
		// TODO: (GH-10) Add config file setting for specifying the journal mode.
		//
		"%s?_journal_mode=%s&_busy_timeout=1000",
		sqliteDBFile,
		"DELETE",
	)
	sqliteDB, err := sql.Open("sqlite3", sqliteDSN)
	if err != nil {
		log.Fatal(err.Error())
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
	//sqliteDB.SetMaxOpenConns(1)

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
		log.Fatalf(
			"verifying SQLite database connection failed: %v",
			err,
		)
	}
	log.Infof("Connection established to SQLite database at %s", sqliteDBFile)
	logConnStats()

	// Start a transaction so we have SQLite database rollback potential
	// if the process is interrupted.
	// TODO: Do we need any specific transaction settings, or continue to
	// pass nil here?
	sqliteTX, sqliteTXErr := sqliteDB.BeginTx(ctx, nil)
	if sqliteTXErr != nil {
		log.Fatalf(
			"failed to start SQL transaction for SQLite database: %v",
			sqliteTXErr,
		)
	}

	// TODO: Move these elsewhere?
	rollbackTX := func(tx *sql.Tx) {
		if txRollbackErr := tx.Rollback(); txRollbackErr != nil {
			log.Fatalf(
				"failed to roll back transaction: %v",
				txRollbackErr,
			)
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
			log.Fatal(rowsCountErr.Error())
		}
		logConnStats()

		log.Debugf("About to execute query to retrieve %d existing rows", rowsCount)
		mysqlRows, err := mysqlDB.Query(querySet[config.SQLQueriesRead])
		if err != nil {
			rollbackTX(sqliteTX)
			log.Fatal(err.Error())
		}
		log.Debug("Rows retrieved from source MySQL database")
		logConnStats()

		// Recreate SQLite database tables.
		//
		// TODO: Research syncing tables instead of recreating each time

		// FIXME: What is a better way to drop the table programatically?
		dropTableStmt := fmt.Sprintf("DROP TABLE IF EXISTS %s", table)
		if _, err := sqliteTX.Exec(dropTableStmt); err != nil {
			rollbackTX(sqliteTX)
			log.Fatalf(
				"failed to run query to drop (potentially) preexisting table %s: %v",
				table,
				err,
			)
		}
		log.Debugf("Successfully ran DROP TABLE query for table %s", table)
		logConnStats()

		dropIndexStmt := fmt.Sprintf("DROP INDEX IF EXISTS %s", table)
		if _, err := sqliteTX.Exec(dropIndexStmt); err != nil {
			rollbackTX(sqliteTX)
			log.Fatalf(
				"failed to run query to drop (potentially) preexisting index for table %s: %v",
				table,
				err,
			)
		}
		log.Debugf("Successfully ran DROP INDEX query for table %s", table)
		logConnStats()

		if _, err := sqliteTX.Exec(querySet[config.SQLQueriesNew]); err != nil {
			rollbackTX(sqliteTX)
			log.Fatalf(
				"failed to run query to create table %s: %v",
				table,
				err,
			)
		}
		log.Debugf("Created table %s in SQLite database", table)
		logConnStats()

		// Q: does this cause an error?
		// A: no, attempting to retrieve a key *not* in a map returns a zero
		// value, in this case an empty string value. Executing an empty
		// string/query appears to be "valid", which gives unexpected results.
		// Our validation method will have to catch problems like this one.
		//
		// if _, err := sqliteDB.Exec(""); err != nil {
		// 	log.Fatalf(
		// 		"failed to run empty exec query for table %s",
		// 		table,
		// 		err,
		// 	)
		// }

		if cfg.SQLiteCreateIndexes() {
			if _, err := sqliteTX.Exec(querySet[config.SQLQueriesIndex]); err != nil {
				rollbackTX(sqliteTX)
				log.Fatalf(
					"failed to run query to create index for table %s: %v",
					table,
					err,
				)
			}
			log.Debugf("Created index for table %s in SQLite database", table)
			logConnStats()
		}

		log.Debug("SQLite database is ready")

		columnNames, err := mysqlRows.Columns()
		if err != nil {
			rollbackTX(sqliteTX)
			log.Fatalf(
				"unable to fetch column names for table %s: %v",
				table,
				err,
			)
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

		if rowsCount > 0 {

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
					log.Fatal(err.Error())
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
							sqliteOutput[idx] = strings.TrimSpace(string(string(byteArray)))
							// log.Debugf("value is %q after trim", values[idx])
						}
					}

					// Rely on fmt.Sprintf to handle formatting, potential
					// conversion of data to string type before storing back
					// in our []interface{} for use by db.Exec
					sqliteOutput[idx] = fmt.Sprintf("%v", sqliteOutput[idx])

				}

				// Populate SQLite database with row retrieved from MySQL
				if _, err = sqliteTX.Exec(querySet[config.SQLQueriesWrite], sqliteOutput...); err != nil {
					rollbackTX(sqliteTX)
					log.Fatal(err.Error())
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
				log.Fatal(err.Error())
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
				log.Fatalf(
					"error closing mysqlRows for table %s: %v",
					table,
					err,
				)
			}
			log.Debug("Successfully closed mysqlRows")
			logConnStats()

		} else {

			log.Infof(
				"no entries in table %s for database %s, skipping rows retrieval and SQLite database insert",
				table,
				cfg.MySQLDatabase(),
			)
		}

	}

	if err := sqliteTX.Commit(); err != nil {
		// FIXME: This is a pretty serious failure; what else can we
		// do at this point?
		log.Fatalf(
			"failed to commit transaction for SQLite database: %v",
			err,
		)
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
