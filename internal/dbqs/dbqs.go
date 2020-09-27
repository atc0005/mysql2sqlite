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

//var ErrStmtClose error = errors.New("failed to close statement")
//var ErrTableCount error = errors.New("failed to retrieve row count for table")

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

		// While the context is passed to mstClient.SendWithContext and it
		// should ensure that it is respected, we check here explicitly in
		// order to return early in an effort to prevent undesired message
		// attempts
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

func RowsCount(db *sql.DB, table string) (int, error) {
	var rowsCount int
	rcQuery := fmt.Sprintf("SELECT COUNT(*) as count FROM %s", table)
	log.Debugf("About to prepare query to fetch row count for %s", table)
	rcStmt, err := db.Prepare(rcQuery)
	if err != nil {
		return -1, fmt.Errorf(
			"failed to prepare row count statement for table %s: %w",
			table,
			err,
		)
	}
	log.Debugf("About to fetch row count for %s", table)
	if err := rcStmt.QueryRow().Scan(&rowsCount); err != nil {
		return -1, fmt.Errorf(
			"failed to retrieve row count for table %s: %w",
			table,
			err,
		)
	}
	log.Debugf("About to close prepared statement for fetching row count for %s", table)
	if err := rcStmt.Close(); err != nil {
		return -1, fmt.Errorf(
			"failed to close rcStmt after retrieving row count for table %s: %w",
			table,
			err,
		)
	}

	return rowsCount, nil
}

// RowsCountMatch receives a table name and
// func RowsCountMatch(table string, dbs ...*sql.DB) (bool, error) {
// 	if len(dbs) < 2 {
// 		return false, fmt.Errorf("invalid number of databases to compare; at least two are required")
// 	}

// 	// all database tables should match, so using the first one for comparison
// 	// against the others should be valid
// 	firstDBRowsCount, rowsCountErr := RowsCount(dbs[0], table)
// 	if rowsCountErr != nil {
// 		return false, rowsCountErr
// 	}

// 	for idx := range dbs {
// 		dbsRowsCount, rowsCountErr := RowsCount(dbs[idx], table)
// 		if rowsCountErr != nil {
// 			return false, rowsCountErr
// 		}
// 		if firstDBRowsCount != dbsRowsCount {
// 			return false, fmt.Errorf(
// 				"mismatched number of rows for %s table; got %d, expected %d",
// 				table,
// 				dbsRowsCount,
// 				firstDBRowsCount,
// 			)
// 		}
// 	}

// 	return true, nil
// }
