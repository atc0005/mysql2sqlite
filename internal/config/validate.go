// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// Package config provides types and functions to collect, validate and apply
// user-provided settings.
package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/atc0005/mysql2sqlite/internal/caller"
	"github.com/atc0005/mysql2sqlite/internal/dbqs"
)

// TCP port ranges
// http://www.iana.org/assignments/port-numbers
// Port numbers are assigned in various ways, based on three ranges: System
// Ports (0-1023), User Ports (1024-49151), and the Dynamic and/or Private
// Ports (49152-65535)
const (
	TCPReservedPort            int = 0
	TCPSystemPortStart         int = 1
	TCPSystemPortEnd           int = 1023
	TCPUserPortStart           int = 1024
	TCPUserPortEnd             int = 49151
	TCPDynamicPrivatePortStart int = 49152
	TCPDynamicPrivatePortEnd   int = 65535
)

// Validate verifies all struct fields have been provided acceptable values
func (c Config) Validate() error {

	myFuncName := caller.GetFuncName()

	// We can't directly test the results of c.ConfigFile() since we also
	// support auto-loading the config file from a set of known locations.
	// Instead, we look an internal flag to check whether the application was
	// able to load *one* of the available config files.
	//
	// if c.ConfigFile() == "" {
	// 	return fmt.Errorf(
	// 		"%s: missing fully-qualified path to config file to load",
	// 		myFuncName,
	// 	)
	// }
	// log.Debugf("c.ConfigFile() validates: %#v", c.ConfigFile())
	if !c.configFileLoaded {
		return ErrCfgFileNotFound
	}
	log.Debug("c.configFileLoaded indicates successful load of configuration file")

	switch c.LogLevel() {
	case LogLevelFatal:
	case LogLevelError:
	case LogLevelWarn:
	case LogLevelInfo:
	case LogLevelDebug:
	default:
		return fmt.Errorf(
			"%s: invalid option %q provided for log level",
			myFuncName,
			c.LogLevel(),
		)
	}
	log.Debugf("c.LogLevel() validates: %#v", c.LogLevel())

	switch c.LogOutput() {
	case LogOutputStderr:
	case LogOutputStdout:
	default:
		return fmt.Errorf("invalid option %q provided for log output",
			c.LogOutput())
	}
	log.Debugf("c.LogOutput() validates: %#v", c.LogOutput())

	switch c.LogFormat() {
	case LogFormatCLI:
	case LogFormatJSON:
	case LogFormatLogFmt:
	case LogFormatText:
	case LogFormatDiscard:
	default:
		return fmt.Errorf(
			"%s: invalid option %q provided for log format",
			myFuncName,
			c.LogFormat(),
		)
	}
	log.Debugf("c.LogFormat() validates: %#v", c.LogFormat())

	if c.MySQLUsername() == "" {
		return fmt.Errorf(
			"%s: missing database username",
			myFuncName,
		)
	}
	log.Debugf("c.MySQLUsername() validates: %#v", c.MySQLUsername())

	if c.MySQLPassword() == "" {
		return fmt.Errorf(
			"%s: missing database user password",
			myFuncName,
		)
	}
	// while it would be useful for troubleshooting purposes to log this
	// value, it is too risky to leave this enabled by default
	log.Debugf("c.MySQLPassword() validates: %#v", "REDACTED")

	if c.MySQLHost() == "" {
		return fmt.Errorf(
			"%s: missing database host",
			myFuncName,
		)
	}
	log.Debugf("c.MySQLHost() validates: %#v", c.MySQLHost())

	switch {
	case (c.MySQLPort() >= TCPSystemPortStart) && (c.MySQLPort() <= TCPDynamicPrivatePortEnd):
		log.Debugf(
			"Port %d is within the range of %d and %d",
			c.MySQLPort(),
			TCPSystemPortStart,
			TCPDynamicPrivatePortEnd,
		)
	default:
		errMsg := fmt.Sprintf(
			"invalid port %d specified; outside of the range of %d and %d",
			c.MySQLPort(),
			TCPSystemPortStart,
			TCPDynamicPrivatePortEnd,
		)
		log.Debugf("%s: %v", myFuncName, errMsg)

		return errors.New(errMsg)
	}
	log.Debugf("c.MySQLPort() validates: %+v", c.MySQLPort())

	switch c.MySQLEncryption() {
	case MySQLEncryptionRequired:
	case MySQLEncryptionPreferred:
	case MySQLEncryptionSkipVerify:
	case MySQLEncryptionDisabled:
	default:
		return fmt.Errorf(
			"%s: invalid option %q provided for MySQL encryption",
			myFuncName,
			c.MySQLEncryption(),
		)
	}
	log.Debugf("c.MySQLEncryption() validates: %#v", c.MySQLEncryption())

	// TODO: Tweak the value used here to reflect real world values
	if c.MySQLConnMaxLifetime() < 1*time.Second {
		return fmt.Errorf(
			"%s: invalid maximum lifetime for MySQL database connection: %v",
			myFuncName,
			c.MySQLConnMaxLifetime(),
		)
	}
	log.Debugf("c.MySQLConnMaxLifetime() validates: %#v", c.MySQLConnMaxLifetime())

	if c.MySQLMaxOpenConns() < 1 {
		return fmt.Errorf(
			"%s: invalid maximum open connections for MySQL database: %v",
			myFuncName,
			c.MySQLMaxOpenConns(),
		)
	}
	log.Debugf("c.MySQLMaxOpenConns() validates: %#v", c.MySQLMaxOpenConns())

	// TODO: Tweak the value used here to reflect real world values
	if c.MySQLConnMaxIdleTime() < 1*time.Second {
		return fmt.Errorf(
			"%s: invalid maximum idle time for MySQL database connection: %v",
			myFuncName,
			c.MySQLConnMaxIdleTime(),
		)
	}
	log.Debugf("c.MySQLConnMaxIdleTime() validates: %#v", c.MySQLConnMaxIdleTime())

	// FIXME: Preventing *any* idle connections (e.g., value of 0) might be
	// legitimate here?
	if c.MySQLMaxIdleConns() < 1 {
		return fmt.Errorf(
			"%s: invalid maximum idle connections for MySQL database: %v",
			myFuncName,
			c.MySQLMaxIdleConns(),
		)
	}
	log.Debugf("c.MySQLMaxIdleConns() validates: %#v", c.MySQLMaxIdleConns())

	if c.SQLiteDBFile() == "" {
		return fmt.Errorf(
			"%s: missing SQLite database filename",
			myFuncName,
		)
	}
	log.Debugf("c.SQLiteDBFile() validates: %#v", c.SQLiteDBFile())

	if c.SQLiteDBPath() == "" {
		return fmt.Errorf(
			"%s: missing SQLite database path",
			myFuncName,
		)
	}
	log.Debugf("c.SQLiteDBPath() validates: %#v", c.SQLiteDBPath())

	// we rely on the true/false nature of booleans to handle validation of
	// the "create indexes" setting.

	// validate presence of all required user-provided queries
	switch c.DBQueries() {
	case nil:
		return fmt.Errorf(
			"%s: database queries used to sync database to SQLite db file not provided",
			myFuncName,
		)
	default:
		for table, querySet := range c.configFileSettings.Queries {
			log.Debugf("table is: %v", table)

			// each of these query types should be defined within the
			// configuration file with `index` being optional
			expectedQuerySet := []string{
				dbqs.SQLQueriesRead,
				dbqs.SQLQueriesNew,
				dbqs.SQLQueriesWrite,
			}

			if c.SQLiteCreateIndexes() {
				expectedQuerySet = append(expectedQuerySet, dbqs.SQLQueriesIndex)
			}

			for _, expectedQuery := range expectedQuerySet {
				query, ok := querySet[expectedQuery]
				if !ok {
					return fmt.Errorf(
						"%s: %s query not provided for database table %v",
						myFuncName,
						expectedQuery,
						table,
					)
				}
				log.Debugf("%q query for %s: %q\n", expectedQuery, table, query)
			}
		}
	}

	if c.ConnectionRetries() < 0 {
		return fmt.Errorf(
			"%s: invalid connection retries count: %v",
			myFuncName,
			c.ConnectionRetries(),
		)
	}
	log.Debugf("c.ConnectionRetries() validates: %#v", c.ConnectionRetries())

	if c.ConnectionRetryDelay() < 1*time.Second {
		return fmt.Errorf(
			"%s: invalid connection retry delay: %v",
			myFuncName,
			c.ConnectionRetryDelay(),
		)
	}
	log.Debugf("c.ConnectionRetryDelay() validates: %#v", c.ConnectionRetryDelay())

	if c.ConnectionTimeout() < 1*time.Second {
		return fmt.Errorf(
			"%s: invalid connection timeout: %v",
			myFuncName,
			c.ConnectionTimeout(),
		)
	}
	log.Debugf("c.ConnectionTimeout() validates: %#v", c.ConnectionTimeout())

	// we rely on the true/false nature of booleans to handle validation of
	// the "trim whitespace" setting.

	// Optimist
	log.Debug("All validation checks pass")
	return nil

}
