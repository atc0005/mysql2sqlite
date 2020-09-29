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

// RowsCount is a helper function that is used to return the number of rows
// for a specified table.
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
