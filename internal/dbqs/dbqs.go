// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package dbqs

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/atc0005/mysql2sqlite/internal/caller"
)

// SQLQueries is a map of maps representing a collection of tables and the
// queries used to read from a source database and write to a SQLite database.
// Example: queries["virtual_domains"]["read"] = "SELECT * FROM virtual_domains"
type SQLQueries map[string]map[string]string

// SQLQuerySet is a map of query type to query. Each entry is used to read
// from a source database table, create, write to or update indexes for a new
// database table in the generated SQLite database.
type SQLQuerySet map[string]string

// Queries used to retrieve data from MySQL database and manage local SQLite
// database
const (
	SQLQueriesRead  string = "read"
	SQLQueriesNew   string = "new"
	SQLQueriesWrite string = "write"
	SQLQueriesIndex string = "index"
)

// var ErrStmtClose error = errors.New("failed to close statement")
// var ErrTableCount error = errors.New("failed to retrieve row count for table")

// VerifyDBConn is used to verify connectivity to a specified database. If the
// initial verification fails, a specified number of further attempts are made
// before an error is returned to signal that a database connection is
// unavailable.
func VerifyDBConn(ctx context.Context, db *sql.DB, retries int, retryDelay time.Duration, timeout time.Duration) error {

	myFuncName := caller.GetFuncName()

	// TODO: Expose timeout as a config setting
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var result error

	// initial attempt + number of specified retries
	attemptsAllowed := 1 + retries

	for attempt := 1; attempt <= attemptsAllowed; attempt++ {
		// the result from the last attempt is returned to the caller
		result = db.PingContext(ctx)

		switch {
		case result == nil:

			log.Debugf(
				"Successfully confirmed database connection after %d of %d attempts",
				attempt,
				attemptsAllowed,
			)

			// No further retries needed
			return nil

		// Check context explicitly in order to return as soon as possible
		case ctx.Err() != nil && result != nil:

			errMsg := fmt.Errorf(
				"%s: context cancelled or expired: %v; "+
					"aborting database connection verification after %d of %d attempts: %w",
				myFuncName,
				ctx.Err().Error(),
				attempt,
				attemptsAllowed,
				result,
			)
			return errMsg

		case result != nil:

			log.Warnf(
				"%s: attempt %d of %d to verify database connection failed: %v",
				myFuncName,
				attempt,
				attemptsAllowed,
				result,
			)

			log.Infof(
				"%s: applying retry delay of %v",
				myFuncName,
				retryDelay,
			)
			time.Sleep(retryDelay)

		}

	}

	return result

}

// RowsCount returns the number of rows for a specified table.
func RowsCount(db *sql.DB, table string) (int, error) {
	var rowsCount int
	rcQuery := fmt.Sprintf("SELECT COUNT(*) as count FROM %s", table)
	log.Debugf("About to fetch row count for %s", table)
	if err := db.QueryRow(rcQuery).Scan(&rowsCount); err != nil {
		return -1, fmt.Errorf(
			"failed to retrieve row count for table %s: %w",
			table,
			err,
		)
	}

	return rowsCount, nil
}

// PrepareSQLiteDB prepares a SQLite database for incoming data from the
// source database. As of the current version, this includes dropping tables
// and recreating them along with optional indexes, if enabled.
func PrepareSQLiteDB(tx *sql.Tx, querySet SQLQuerySet, table string, createIndexes bool) error {
	// Recreate SQLite database tables.
	//
	// TODO: Research syncing tables instead of recreating each time

	// FIXME: What is a better way to drop the table programatically?
	dropTableStmt := fmt.Sprintf("DROP TABLE IF EXISTS %s", table)
	if _, err := tx.Exec(dropTableStmt); err != nil {
		log.Errorf(
			"failed to run query to drop (potentially) preexisting table %s: %v",
			table,
			err,
		)
		return err
	}
	log.Debugf("Successfully ran DROP TABLE query for table %s", table)

	dropIndexStmt := fmt.Sprintf("DROP INDEX IF EXISTS %s", table)
	if _, err := tx.Exec(dropIndexStmt); err != nil {
		log.Errorf(
			"failed to run query to drop (potentially) preexisting index for table %s: %v",
			table,
			err,
		)
		return err
	}
	log.Debugf("Successfully ran DROP INDEX query for table %s", table)

	if _, err := tx.Exec(querySet[SQLQueriesNew]); err != nil {
		log.Errorf(
			"failed to run query to create table %s: %v",
			table,
			err,
		)
		return err
	}
	log.Debugf("Created table %s in SQLite database", table)

	// Q: does this cause an error?
	// A: no, attempting to retrieve a key *not* in a map returns a zero
	// value, in this case an empty string value. Executing an empty
	// string/query appears to be "valid", which gives unexpected results.
	// Our validation method will have to catch problems like this one.
	//
	// if _, err := sqliteDB.Exec(""); err != nil {
	// 	log.Errorf(
	// 		"failed to run empty exec query for table %s",
	// 		table,
	// 		err,
	// 	)
	//	return
	// }

	if createIndexes {
		if _, err := tx.Exec(querySet[SQLQueriesIndex]); err != nil {

			log.Errorf(
				"failed to run query to create index for table %s: %v",
				table,
				err,
			)
			return err
		}
		log.Debugf("Created index for table %s in SQLite database", table)
	}

	log.Debug("SQLite database is ready")
	return nil
}
