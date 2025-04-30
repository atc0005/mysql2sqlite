// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/atc0005/mysql2sqlite/internal/dbqs"
)

// ConfigFile returns the user-provided path to the config file for this
// application or the default value if not provided. CLI flag or environment
// variables are the only way to specify a value for this setting.
func (c Config) ConfigFile() string {
	switch {
	case c.flagSettings.ConfigFile != nil:
		return *c.flagSettings.ConfigFile
	default:
		return defaultConfigFile
	}
}

// LogLevel returns the user-provided logging level or the default value if
// not provided. CLI flag values take precedence if provided.
func (c Config) LogLevel() string {
	switch {
	case c.flagSettings.LogLevel != nil:
		return *c.flagSettings.LogLevel
	case c.configFileSettings.Logging.Level != nil:
		return *c.configFileSettings.Logging.Level
	default:
		return defaultLogLevel
	}
}

// LogOutput returns the user-provided logging output or the default value if
// not provided. CLI flag values take precedence if provided.
func (c Config) LogOutput() string {
	switch {
	case c.flagSettings.LogOutput != nil:
		return *c.flagSettings.LogOutput
	case c.configFileSettings.Logging.Output != nil:
		return *c.configFileSettings.Logging.Output
	default:
		return defaultLogOutput
	}
}

// LogFormat returns the user-provided logging format or the default value if
// not provided. CLI flag values take precedence if provided.
func (c Config) LogFormat() string {
	switch {
	case c.flagSettings.LogFormat != nil:
		return *c.flagSettings.LogFormat
	case c.configFileSettings.Logging.Format != nil:
		return *c.configFileSettings.Logging.Format
	default:
		return defaultLogFormat
	}
}

// LogDBStats returns the user-provided choice of whether database connection
// stats are logged periodically or the default value if not provided.
func (c Config) LogDBStats() bool {
	switch {
	case c.configFileSettings.Logging.Stats != nil:
		return *c.configFileSettings.Logging.Stats
	default:
		return defaultLogDBStats
	}
}

// ConfigFileUsed returns the configuration file that was located and loaded
// for application use. This may match the user-specified file or it may
// instead be an alternate file automatically located if the user-specified
// file could not be located. This method relies upon the configuration
// validation checks applied at startup to ensure that a valid configuration
// file is returned.
func (c Config) ConfigFileUsed() string {
	switch {
	case c.configFileUsed != "":
		return c.configFileUsed
	default:
		return ""
	}
}

// DBServerReadTimeout returns the user-provided choice of what timeout value to use for
// attempts to query the remote database server. If not set, returns the
// default value for our application.
// func (c Config) DBServerReadTimeout() time.Duration {

// 	switch {
// 	case c.dbServerReadTimeout != 0:
// 		return time.Duration(c.dbServerReadTimeout) * time.Second
// 	default:
// 		log.Debugf(
// 			"Requested config read timeout value not specified, using default: %v",
// 			defaultConfigReadTimeout,
// 		)
// 		return time.Duration(defaultConfigReadTimeout) * time.Second
// 	}
// }

// MySQLUsername returns the user-provided MySQL username for this application
// or the default value if not provided.
func (c Config) MySQLUsername() string {
	switch {
	case c.configFileSettings.MySQLConfig.Username != nil:
		return *c.configFileSettings.MySQLConfig.Username
	default:
		return defaultMySQLUsername
	}
}

// MySQLPassword returns the password for the user-provided MySQL username &
// password pair for this application or the default value if not provided.
func (c Config) MySQLPassword() string {
	switch {
	case c.configFileSettings.MySQLConfig.Password != nil:
		return *c.configFileSettings.MySQLConfig.Password
	default:
		return defaultMySQLPassword
	}
}

// MySQLHost returns the user-provided IP Address, FQDN or UNIX socket for the
// MySQL server used by this application or the default value if not provided.
func (c Config) MySQLHost() string {
	switch {
	case c.configFileSettings.MySQLConfig.Host != nil:
		return *c.configFileSettings.MySQLConfig.Host
	default:
		return defaultMySQLHost
	}
}

// MySQLPort returns the user-provided TCP port for the MySQL server used by
// this application or the default value if not provided.
func (c Config) MySQLPort() int {
	switch {
	case c.configFileSettings.MySQLConfig.Port != nil:
		return *c.configFileSettings.MySQLConfig.Port
	default:
		return defaultMySQLPort
	}
}

// MySQLAddress returns the address used by the MySQL driver to connect to the
// specified database. This is a formatted string containing the user-provided
// IP Address or FQDN and TCP port or the UNIX socket for the MySQL server
// used by this application. This value is empty if no user-provided MySQL
// host or TCP port is provided.
func (c Config) MySQLAddress() string {
	host := c.MySQLHost()

	switch {
	// If a forward or back slash is present we assume a UNIX socket was
	// specified for connectivity.
	case strings.ContainsAny(host, `\/`):
		return fmt.Sprintf(
			"unix(%s)",
			host,
		)

	// If a host value was specified we use a user-specified or default TCP
	// port for the address string.
	case host != "":
		port := strconv.Itoa(c.MySQLPort())

		return fmt.Sprintf(
			"@tcp(%s)",
			net.JoinHostPort(host, port),
		)

	default:
		return host
	}
}

// MySQLDatabase returns the user-provided MySQL database used by this
// application or the default value if not provided.
func (c Config) MySQLDatabase() string {
	switch {
	case c.configFileSettings.MySQLConfig.Database != nil:
		return *c.configFileSettings.MySQLConfig.Database
	default:
		return defaultMySQLDatabase
	}
}

// MySQLEncryption returns the user-provided choice regarding what encryption
// settings to use when connecting to the user-provided MySQL host used by
// this application or the default value if not provided.
func (c Config) MySQLEncryption() string {
	switch {
	case c.configFileSettings.MySQLConfig.Encryption != nil:
		return *c.configFileSettings.MySQLConfig.Encryption
	default:
		return defaultMySQLEncryption
	}
}

// MySQLConnMaxLifetime returns the user-provided maximum lifetime in minutes
// for database connections to the MySQL server ued by this application or the
// default value if not provided. See also
// https://github.com/go-sql-driver/mysql#important-settings
func (c Config) MySQLConnMaxLifetime() time.Duration {
	switch {
	case c.configFileSettings.MySQLConfig.ConnMaxLifetime != nil:
		return time.Minute * time.Duration(
			*c.configFileSettings.MySQLConfig.ConnMaxLifetime,
		)
	default:
		return defaultMySQLConnMaxLifetime
	}
}

// MySQLMaxOpenConns returns the user-provided maximum number of database
// connections for the MySQL server used by this application or the default
// value if not provided. See also
// https://github.com/go-sql-driver/mysql#important-settings
func (c Config) MySQLMaxOpenConns() int {
	switch {
	case c.configFileSettings.MySQLConfig.MaxOpenConns != nil:
		return *c.configFileSettings.MySQLConfig.MaxOpenConns
	default:
		return defaultMySQLMaxOpenConns
	}
}

// MySQLMaxIdleConns returns the user-provided maximum number of idle database
// connections for the MySQL server used by this application or the default
// value if not provided. See also
// https://github.com/go-sql-driver/mysql#important-settings
func (c Config) MySQLMaxIdleConns() int {
	switch {
	case c.configFileSettings.MySQLConfig.MaxIdleConns != nil:
		return *c.configFileSettings.MySQLConfig.MaxIdleConns
	default:
		return defaultMySQLMaxIdleConns
	}
}

// MySQLConnMaxIdleTime returns the user-provided maximum time in seconds that
// a database connection can remain idle or the default value if not provided.
// See also https://github.com/go-sql-driver/mysql#important-settings
func (c Config) MySQLConnMaxIdleTime() time.Duration {
	switch {
	case c.configFileSettings.MySQLConfig.ConnMaxIdleTime != nil:
		return time.Second * time.Duration(
			*c.configFileSettings.MySQLConfig.ConnMaxIdleTime,
		)
	default:
		return defaultMySQLConnMaxIdleTime
	}
}

// SQLiteDBFile returns the user-provided SQLite database file created/managed
// by this application or the default value if not provided.
func (c Config) SQLiteDBFile() string {
	switch {
	case c.configFileSettings.SQLiteConfig.Filename != nil:
		return *c.configFileSettings.SQLiteConfig.Filename
	default:
		return defaultSQLiteDBFile
	}
}

// SQLiteDBPath returns the user-provided directory where the SQLite database
// file created/managed by this application is stored or the default value if
// not provided.
func (c Config) SQLiteDBPath() string {
	switch {
	case c.configFileSettings.SQLiteConfig.Path != nil:
		return *c.configFileSettings.SQLiteConfig.Path
	default:
		return defaultSQLiteDBPath
	}
}

// SQLiteCreateIndexes returns the user-provided choice regarding whether
// index creation queries will be used when generating the SQLite database
// file or the default value if not provided.
func (c Config) SQLiteCreateIndexes() bool {
	switch {
	case c.configFileSettings.SQLiteConfig.CreateIndexes != nil:
		return *c.configFileSettings.SQLiteConfig.CreateIndexes
	default:
		return defaultSQLiteCreateIndexes
	}
}

// SQLiteBusyTimeout returns the user-provided choice regarding busy timeout
// behavior for the connection to the SQLite database file or the default
// value if not provided.
func (c Config) SQLiteBusyTimeout() int {
	switch {
	case c.configFileSettings.SQLiteConfig.BusyTimeout != nil:
		return *c.configFileSettings.SQLiteConfig.BusyTimeout
	default:
		return defaultSQLiteBusyTimeout
	}
}

// SQLiteJournalMode returns the user-provided choice of journal mode for the
// database associated with the database connection to the SQLite database
// file or the default value if not provided.
func (c Config) SQLiteJournalMode() string {
	switch {
	case c.configFileSettings.SQLiteConfig.JournalMode != nil:
		return *c.configFileSettings.SQLiteConfig.JournalMode
	default:
		return defaultSQLiteJournalMode
	}
}

// TrimWhitespace returns the user-provided choice regarding whether data
// retrieved from MySQL should have leading and trailing whitespace removed
// before inserting into the SQLite database or the default value if not
// provided.
func (c Config) TrimWhitespace() bool {
	switch {
	case c.configFileSettings.General.TrimWhitespace != nil:
		return *c.configFileSettings.General.TrimWhitespace
	default:
		return defaultTrimWhitespace
	}
}

// ConnectionRetries returns the user-provided number of times a connection
// attempt should be retried before returning an error or the default value if
// not provided.
func (c Config) ConnectionRetries() int {
	switch {
	case c.configFileSettings.General.ConnectionRetries != nil:
		return *c.configFileSettings.General.ConnectionRetries
	default:
		return defaultConnectionRetries
	}
}

// ConnectionRetryDelay returns the user-provided number of seconds between
// retry connection attempts or the default value if not provided.
func (c Config) ConnectionRetryDelay() time.Duration {
	switch {
	case c.configFileSettings.General.ConnectionRetryDelay != nil:
		return time.Duration(*c.configFileSettings.General.ConnectionRetryDelay) * time.Second
	default:
		return time.Duration(defaultConnectionRetryDelay) * time.Second
	}
}

// ConnectionTimeout returns the user-provided number of seconds before a
// connection attempt should be aborted or the default value if not provided.
func (c Config) ConnectionTimeout() time.Duration {
	switch {
	case c.configFileSettings.General.ConnectionTimeout != nil:
		return time.Duration(*c.configFileSettings.General.ConnectionTimeout) * time.Second
	default:
		return time.Duration(defaultConnectionTimeout) * time.Second
	}
}

// DBQueries returns the user-provided collection of tables and the queries
// used to read from a source database and write to a SQLite database. If not
// provided, nil is returned in order to force validation to fail.
func (c Config) DBQueries() dbqs.SQLQueries {
	switch {
	case c.configFileSettings.Queries != nil:
		return c.configFileSettings.Queries
	default:
		return nil
	}
}
