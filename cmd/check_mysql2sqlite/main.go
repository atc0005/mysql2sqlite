// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"database/sql"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/atc0005/mysql2sqlite/internal/caller"
	"github.com/atc0005/mysql2sqlite/internal/config"
	"github.com/atc0005/mysql2sqlite/internal/dbqs"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/apex/log"

	"github.com/atc0005/go-nagios"
)

func main() {

	//
	// FIXME: Connection close info isn't being logged to nagios.log file
	//

	// Set initial "state" as valid, adjust as we go.
	var nagiosExitState = nagios.ExitState{
		LastError:      nil,
		ExitStatusCode: nagios.StateOKExitCode,
	}

	//nagiosExitState.ServiceOutput = "Tacos"

	// defer this from the start so it is the last deferred function to run
	defer nagiosExitState.ReturnCheckResults()

	nagiosExitState.WarningThreshold =
		"Recoverable error, potentially resolved with the next database sync job"
	nagiosExitState.CriticalThreshold =
		"Non-recoverable error, requires sysadmin intervention"

	cfg, err := config.NewConfig()
	if err != nil {
		nagiosExitState.LastError = err
		nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
		log.Error(err.Error())

		nagiosExitState.ServiceOutput = fmt.Sprintf(
			"%s: Failed to load configuration: %v",
			nagios.StateCRITICALLabel,
			err,
		)

		// no need to go any further, we *want* to exit right away; we don't
		// have a working configuration and there isn't anything further to do
		return
	}

	// Flesh out nagiosExitState with some additional common details now that
	// the flags and config file settings have been parsed
	nagiosExitState.LongServiceOutput = fmt.Sprintf(
		"* SQLite database: %q%s"+
			"* MySQL database: %q%s"+
			"* MySQL host: %q%s"+
			"* MySQL host port: %d%s",
		path.Join(cfg.SQLiteDBPath(), cfg.SQLiteDBFile()),
		nagios.CheckOutputEOL,
		cfg.MySQLDatabase(),
		nagios.CheckOutputEOL,
		cfg.MySQLHost(),
		nagios.CheckOutputEOL,
		cfg.MySQLPort(),
		nagios.CheckOutputEOL,
	)

	log.Info(config.Branding())

	// Only issue warning if stderr was not chosen as output target
	if cfg.LogOutput() != config.LogOutputStderr {
		log.Warn("Reminder: Use `--log-output stderr` flag to prevent interleaving log output with Nagios status information")
	}

	mysqlDSN := fmt.Sprintf(
		// Based on testing, automatic date/time parsing appears to be required
		// TODO: Consider exposing this via config file setting.
		"%s:%s@tcp(%s:%d)/%s?tls=%s&parseTime=true",
		// "%s:%s@tcp(%s:%d)/%s?tls=%s",
		cfg.MySQLUsername(),
		cfg.MySQLPassword(),
		cfg.MySQLHost(),
		cfg.MySQLPort(),
		cfg.MySQLDatabase(),
		cfg.MySQLEncryption(),
	)

	mysqlDB, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		nagiosExitState.LastError = err
		nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
		log.Error(err.Error())

		nagiosExitState.ServiceOutput = fmt.Sprintf(
			"%s: Failed to open MySQL database: %v",
			nagios.StateCRITICALLabel,
			err,
		)

		// we *want* to exit and allow any deferred functions to run; we don't
		// have a working connection to the database server and we're opting
		// to skip further attempts
		//
		// TODO: Add retry support in a future release
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
	if err = mysqlDB.Ping(); err != nil {
		nagiosExitState.LastError = err
		nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
		log.Error(err.Error())

		nagiosExitState.ServiceOutput = fmt.Sprintf(
			"%s: Failed to establish connection to MySQL database: %v",
			nagios.StateCRITICALLabel,
			err,
		)

		return
	}
	log.Infof(
		"Connection established to database server at %s:%d",
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

	// sqliteDB, err := sql.Open(
	// 	"sqlite3",
	// 	fmt.Sprintf(
	// 		"file:%s?_loc=auto",
	// 		sqliteDBFile,
	// 	),
	//
	// NOTE: Automatic locale handling *might* be needed depending on the
	// timezone settings between source database and SQLite database.Disabled
	// for now until we can determine that there is a problem.
	//
	sqliteDB, err := sql.Open("sqlite3", sqliteDSN)
	if err != nil {
		nagiosExitState.LastError = err
		nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
		log.Error(err.Error())

		nagiosExitState.ServiceOutput = fmt.Sprintf(
			"%s: Failed to open SQLite database: %v",
			nagios.StateCRITICALLabel,
			err,
		)

		// we *want* to exit and allow any deferred functions to run; we don't
		// have a working connection to the database server and we're opting
		// to skip further attempts
		//
		// TODO: Add retry support in a future release
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
	// causing this app to lock up when using this setting.
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
	if err = sqliteDB.Ping(); err != nil {
		nagiosExitState.LastError = err
		nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
		log.Error(err.Error())

		nagiosExitState.ServiceOutput = fmt.Sprintf(
			"%s: Failed to establish connection to SQLite database: %v",
			nagios.StateCRITICALLabel,
			err,
		)

		return
	}
	log.Infof("Connection established to SQLite database %s", sqliteDBFile)
	logConnStats()

	for table, querySet := range cfg.DBQueries() {
		log.Debugf(
			"reading table %s from source database %s",
			table,
			cfg.MySQLDatabase(),
		)

		mysqlRowsCount, rowsCountErr := dbqs.RowsCount(mysqlDB, table)
		if rowsCountErr != nil {
			nagiosExitState.LastError = err
			nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
			log.Error(err.Error())

			nagiosExitState.ServiceOutput = fmt.Sprintf(
				"%s: Failed to retrieve rows count for table %s in MySQL database: %v",
				nagios.StateCRITICALLabel,
				table,
				err,
			)

			return
		}
		logConnStats()

		sqliteRowsCount, rowsCountErr := dbqs.RowsCount(sqliteDB, table)
		if rowsCountErr != nil {
			nagiosExitState.LastError = err
			nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
			log.Error(err.Error())

			nagiosExitState.ServiceOutput = fmt.Sprintf(
				"%s: Failed to retrieve rows count for table %s in SQLite database: %v",
				nagios.StateCRITICALLabel,
				table,
				err,
			)

			return
		}
		logConnStats()

		// validate that the same number of rows are present in each table
		if mysqlRowsCount != sqliteRowsCount {
			err := fmt.Errorf(
				"Mismatched number of rows for %s table; MySQL: %d, SQLite: %d",
				table,
				mysqlRowsCount,
				sqliteRowsCount,
			)

			nagiosExitState.LastError = err
			// NOTE: The assumption here is that a mismatch is temporary until
			// the next scheduled execution of the application, at which point
			// the local SQLite database should be in sync with the source
			// database.
			nagiosExitState.ExitStatusCode = nagios.StateWARNINGExitCode
			log.Error(err.Error())

			nagiosExitState.ServiceOutput = fmt.Sprintf(
				"%s: Validation failure: %v",
				nagios.StateWARNINGLabel,
				err,
			)

			return
		}

		mysqlRows, readQueryErr := mysqlDB.Query(querySet[config.SQLQueriesRead])
		if readQueryErr != nil {
			nagiosExitState.LastError = readQueryErr
			nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
			log.Error(readQueryErr.Error())

			nagiosExitState.ServiceOutput = fmt.Sprintf(
				"%s: %s query for table %s in MySQL database failed: %v",
				nagios.StateCRITICALLabel,
				config.SQLQueriesRead,
				table,
				readQueryErr,
			)

			return
		}

		// duplicated to test connections closure stats
		// defer func() {
		// 	log.Warn("Closing mysqlRows via deferred func")
		// 	mysqlRows.Close()
		// 	logConnStats()
		// }()
		log.Debugf(
			"Rows retrieved from table %s in source MySQL database",
			table,
		)
		logConnStats()

		sqliteRows, readQueryErr := sqliteDB.Query(querySet[config.SQLQueriesRead])
		if readQueryErr != nil {
			nagiosExitState.LastError = readQueryErr
			nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
			log.Error(readQueryErr.Error())

			nagiosExitState.ServiceOutput = fmt.Sprintf(
				"%s: %s query for table %s in SQLite database failed: %v",
				nagios.StateCRITICALLabel,
				config.SQLQueriesRead,
				table,
				readQueryErr,
			)

			return
		}
		// duplicated to test connections closure stats
		// defer func() {
		// 	log.Warn("Closing sqliteRows via deferred func")
		// 	sqliteRows.Close()
		// 	logConnStats()
		// }()
		log.Debugf(
			"Rows retrieved from table %s in mirror SQLite database",
			table,
		)
		logConnStats()

		mysqlColumnNames, err := mysqlRows.Columns()
		if err != nil {
			nagiosExitState.LastError = err
			nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
			log.Error(err.Error())

			nagiosExitState.ServiceOutput = fmt.Sprintf(
				"%s: Failed to retrieve column names for table %s in MySQL database: %v",
				nagios.StateCRITICALLabel,
				table,
				err,
			)

			return
		}

		sqliteColumnNames, err := sqliteRows.Columns()
		if err != nil {
			nagiosExitState.LastError = err
			nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
			log.Error(err.Error())

			nagiosExitState.ServiceOutput = fmt.Sprintf(
				"%s: Failed to retrieve column names for table %s in SQLite database: %v",
				nagios.StateCRITICALLabel,
				table,
				err,
			)

			return
		}

		mysqlColCount := len(mysqlColumnNames)
		sqliteColCount := len(sqliteColumnNames)

		// validate that the same number of columns are present in each table
		if mysqlColCount != sqliteColCount {
			err := fmt.Errorf(
				"Mismatched number of columns for %s table; MySQL: %d, SQLite: %d",
				table,
				mysqlColCount,
				sqliteColCount,
			)

			nagiosExitState.LastError = err
			nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
			log.Error(err.Error())

			nagiosExitState.ServiceOutput = fmt.Sprintf(
				"%s: Validation failure: %v",
				nagios.StateWARNINGLabel,
				err,
			)

			return
		}
		logConnStats()

		// The `valuePrts` slice is used to satisfy the `(*sql.Rows).Scan`
		// method requiring that arguments be specified as pointers to a
		// value.
		//
		// We create pointers in valuePtrs which directly point to the value
		// in the "sibling" `values` slice. For this to work, we intentionally
		// *do not* specify a length of zero; we need to prefill the slice
		// entries with applicable zero values that we can index into and
		// update later.
		mysqlRowValues := make([]interface{}, mysqlColCount)
		mysqlRowValuePtrs := make([]interface{}, mysqlColCount)
		sqliteRowValues := make([]interface{}, sqliteColCount)
		sqliteRowValuePtrs := make([]interface{}, sqliteColCount)

		log.Debug("Before *Rows.Next()")

		// NOTE: at this point we have confirmed that mysqlRowsCount == sqliteRowsCount
		if mysqlRowsCount > 0 && sqliteRowsCount > 0 {

			log.Debugf(
				"%d rows to be read from source table %s",
				mysqlRowsCount,
				table,
			)

			for mysqlRows.Next() && sqliteRows.Next() {
				logConnStats()

				// Initialize slice of pointers in order to satisfy the
				// `(*sql.Rows).Scan` method.
				for idx := range mysqlColumnNames {
					mysqlRowValuePtrs[idx] = &mysqlRowValues[idx]
				}

				for idx := range sqliteColumnNames {
					sqliteRowValuePtrs[idx] = &sqliteRowValues[idx]
				}

				if err = mysqlRows.Scan(mysqlRowValuePtrs...); err != nil {
					nagiosExitState.LastError = err
					nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
					log.Error(err.Error())

					nagiosExitState.ServiceOutput = fmt.Sprintf(
						"%s: Failed to retrieve data for table %s in MySQL database: %v",
						nagios.StateCRITICALLabel,
						table,
						err,
					)

					return
				}
				log.Debug("Completed scanning from MySQL row")
				logConnStats()

				if err = sqliteRows.Scan(sqliteRowValuePtrs...); err != nil {
					nagiosExitState.LastError = err
					nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
					log.Error(err.Error())

					nagiosExitState.ServiceOutput = fmt.Sprintf(
						"%s: Failed to retrieve data for table %s in SQLite database: %v",
						nagios.StateCRITICALLabel,
						table,
						err,
					)

					return
				}
				log.Debug("Completed scanning from SQLite row")
				logConnStats()

				// The db.Exec method requires a slice of interface{}, but in
				// order to force SQLite to *not* store the MySQL data as
				// "blob" type we force a conversion to string before storing
				// in a slice of interface{} for db.Exec to use.
				//
				// If the user requested it, we also trim whitespace from the
				// source MySQL data.
				mysqlData := make([]interface{}, len(mysqlRowValues))
				for idx := range mysqlRowValues {
					byteArray, ok := mysqlRowValues[idx].([]byte)
					if !ok {
						log.Debugf("Field %q does NOT contain byte slice value", mysqlColumnNames[idx])
						log.Debugf("Saving field %q value to mysqlData slice as-is", mysqlColumnNames[idx])
						mysqlData[idx] = mysqlRowValues[idx]
					} else {
						log.Debugf("Field %q does contain byte slice value", mysqlColumnNames[idx])
						log.Debugf("Converting field %q value to mysqlData as string value", mysqlColumnNames[idx])
						mysqlData[idx] = string(byteArray)

						if cfg.TrimWhitespace() {
							log.Debugf("Trimming whitespace from field %q", mysqlColumnNames[idx])
							mysqlData[idx] = strings.TrimSpace(string(string(byteArray)))
						}
					}
					timeValue, ok := mysqlRowValues[idx].(time.Time)
					if ok {
						log.Debugf("Field %q contains time value", mysqlColumnNames[idx])
						log.Debugf("Saving field %q value to sqliteData as time value", mysqlColumnNames[idx])
						mysqlData[idx] = timeValue
					}

					// Rely on fmt.Sprintf to handle formatting, potential
					// conversion of data to string type before storing back
					// in our []interface{} for use by db.Exec
					mysqlData[idx] = fmt.Sprintf("%v", mysqlData[idx])

				}

				sqliteData := make([]interface{}, len(sqliteRowValues))
				for idx := range sqliteRowValues {
					byteArray, ok := sqliteRowValues[idx].([]byte)
					if !ok {
						log.Debugf("Field %q does NOT contain byte slice value", sqliteColumnNames[idx])
						log.Debugf("Saving field %q value to sqliteData slice as-is", sqliteColumnNames[idx])
						sqliteData[idx] = sqliteRowValues[idx]
					} else {
						log.Debugf("Field %q does contain byte slice value", sqliteColumnNames[idx])
						log.Debugf("Converting field %q value to sqliteData as string value", sqliteColumnNames[idx])
						sqliteData[idx] = string(byteArray)

						if cfg.TrimWhitespace() {
							log.Debugf("Trimming whitespace from field %q", sqliteColumnNames[idx])
							sqliteData[idx] = strings.TrimSpace(string(string(byteArray)))
						}
					}
					timeValue, ok := sqliteRowValues[idx].(time.Time)
					if ok {
						log.Debugf("Field %q contains time value", sqliteColumnNames[idx])
						log.Debugf("Saving field %q value to sqliteData as time value", sqliteColumnNames[idx])
						sqliteData[idx] = timeValue
					}

					// Rely on fmt.Sprintf to handle formatting, potential
					// conversion of data to string type before storing back
					// in our []interface{} for use by db.Exec
					sqliteData[idx] = fmt.Sprintf("%v", sqliteData[idx])

				}

				// Compare slice values between database tables
				//
				// NOTE: We checked row/column length earlier, so we should be
				// able to index into either of mysqlData or sqliteData at the
				// same points without risk of triggering a panic.
				for idx := range mysqlData {
					what := fmt.Sprintf("field %s in table %s; MySQL: %v, SQLite: %v",
						mysqlColumnNames[idx],
						table,
						mysqlData[idx],
						sqliteData[idx],
					)
					if mysqlData[idx] != sqliteData[idx] {
						err := fmt.Errorf("FAILED data match for %s", what)

						nagiosExitState.LastError = err
						// NOTE: The assumption here is that a mismatch is
						// temporary until the next scheduled execution of the
						// application, at which point the local SQLite
						// database should be in sync with the source
						// database.
						nagiosExitState.ExitStatusCode = nagios.StateWARNINGExitCode
						log.Error(err.Error())

						// NOTE: Intentionally use an abbreviated version of
						// the formatted error message. This is due to length
						// restrictions; i.e., the ServiceOutput value is used
						// for Teams, email and other notification Subject
						// lines.
						nagiosExitState.ServiceOutput = fmt.Sprintf(
							"%s: Validation failure for field %s in table %s",
							nagios.StateCRITICALLabel,
							mysqlColumnNames[idx],
							table,
						)

						return
					} else {
						log.Debugf("Successful data match for %s", what)
					}
				}
			}

			// check for potential errors after processing rows
			if err := mysqlRows.Err(); err != nil {
				nagiosExitState.LastError = err
				nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
				log.Error(err.Error())

				nagiosExitState.ServiceOutput = fmt.Sprintf(
					"%s: Errors occurred performing MySQL database queries: %v",
					nagios.StateCRITICALLabel,
					err,
				)

				return
			}
			log.Infof(
				"Successfully read %d rows from MySQL database table %s",
				mysqlRowsCount,
				table,
			)
			logConnStats()

			// check for potential errors after processing rows
			if err := sqliteRows.Err(); err != nil {
				nagiosExitState.LastError = err
				nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
				log.Error(err.Error())

				nagiosExitState.ServiceOutput = fmt.Sprintf(
					"%s: Errors occurred performing SQLite database queries: %v",
					nagios.StateCRITICALLabel,
					err,
				)

				return
			}
			log.Infof(
				"Successfully read %d rows from SQLite database table %s",
				mysqlRowsCount,
				table,
			)
			logConnStats()

			// explicit close here vs defer since we're in a loop and do not want
			// to persist the database connection for longer than necessary
			if err := mysqlRows.Close(); err != nil {
				err := fmt.Errorf(
					"error closing mysqlRows for table %s: %w",
					table,
					err,
				)
				nagiosExitState.LastError = err
				nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
				log.Error(err.Error())

				nagiosExitState.ServiceOutput = fmt.Sprintf(
					"%s: Failed to close mysqlRows object: %v",
					nagios.StateCRITICALLabel,
					err,
				)

				return
			}
			log.Debug("Successfully closed MySQL rows")
			logConnStats()

			if err := sqliteRows.Close(); err != nil {
				err := fmt.Errorf(
					"error closing sqliteRows for table %s: %w",
					table,
					err,
				)
				nagiosExitState.LastError = err
				nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
				log.Error(err.Error())

				nagiosExitState.ServiceOutput = fmt.Sprintf(
					"%s: Failed to close sqliteRows object: %v",
					nagios.StateCRITICALLabel,
					err,
				)

				return
			}
			log.Debug("Successfully closed SQLite rows")
			logConnStats()

		} else {
			log.Debugf(
				"Skipping evaluation of table %s; there are no rows for this table in either database",
				table,
			)
		}
		logConnStats()

	}

	successMsg := fmt.Sprintf(
		"%s: SQLite database %q matches source database %q",
		nagios.StateOKLabel,
		sqliteDBFile,
		cfg.MySQLDatabase(),
	)

	log.Info(successMsg)

	nagiosExitState.LastError = nil
	nagiosExitState.ExitStatusCode = nagios.StateOKExitCode
	nagiosExitState.ServiceOutput = successMsg

	// the next step after this is the execution of the deferred db handle
	// closures, then ultimately the deferred Nagios check results.
	logConnStats()

}
