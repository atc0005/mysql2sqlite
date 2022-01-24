// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import "time"

// NOTE: See also logging.go for other constants

const (

	// MyAppName is the branded name of this application/project. This value will
	// be used in user-facing output.
	MyAppName string = "mysql2sqlite"

	// MyAppURL is the branded homepage or project repo location. This value will
	// be used in user-facing output.
	MyAppURL string = "https://github.com/atc0005/" + MyAppName

	// MyAppDescription is used when displaying help text.
	MyAppDescription string = "Generate SQLite database from specified MySQL database tables"
)

// MySQL connection encryption settings
// https://github.com/go-sql-driver/mysql#tls
const (
	MySQLEncryptionRequired   string = "true"
	MySQLEncryptionPreferred  string = "preferred"
	MySQLEncryptionSkipVerify string = "skip-verify"
	MySQLEncryptionDisabled   string = "false"
)

// Default flag settings if not overridden by user input.
const (
	// defaultTrimWhitespace controls whether leading and trailing whitespace
	// is automatically removed before writing out retrieved MySQL database
	// values to the SQLite database.
	defaultTrimWhitespace bool = false

	// defaultConnectionRetries is the number of times a connection to a
	// database will be retried by default before giving up and returning an
	// error.
	defaultConnectionRetries int = 5

	// defaultConnectionRetryDelay is the default number of seconds between
	// connection retry attempts.
	defaultConnectionRetryDelay int = 2

	// defaultConnectionTimeout is the default number of seconds before a
	// connection times out and is terminated.
	defaultConnectionTimeout int = 10

	defaultLogLevel  string = "info"
	defaultLogOutput string = "stdout"
	defaultLogFormat string = "text"

	// defaultLogDBStats controls whether database connection stats are logged
	// periodically throughout the application. This is intended to aid in
	// troubleshooting.
	defaultLogDBStats bool = false

	// defaultConfigFile is the fully-qualified path to the configuration file
	// provided by the user via CLI flag. The default is an empty string to
	// indicate that a value was not provided by the user.
	defaultConfigFile string = ""

	// defaultConfigFileName is the "bare" or "non-qualified" configuration
	// filename. If not specified explicitly via CLI flag, this file will be
	// looked for alongside the app's location in addition to the config path
	// for the user account running this app.
	defaultConfigFileName string = "config.yaml"

	// defaultMySQLUser is the default user account that should be used for
	// logins if not provided by the user. It is probably best to *not* set
	// this as it could potentially result in the IP Address of the system
	// running this application being blocked by the database server for
	// invalid logins.
	defaultMySQLUsername string = ""

	// defaultMySQLPassword is intentionally left as an empty string. It is
	// probably best to *not* set this as it could potentially result in the
	// IP Address of the system running this application being blocked by the
	// database server for invalid logins.
	defaultMySQLPassword string = ""

	// defaultMySQLHost is intentionally left as an empty string. While less
	// potentially problematic as setting a default user/password pair, this
	// could still potentially lead to issues.
	defaultMySQLHost string = ""

	// defaultMySQLPort reflects the IANA registered TCP port for MySQL.
	defaultMySQLPort int = 3306

	// defaultMySQLDatabase is intentionally left as an empty string. While
	// less potentially problematic as setting a default user/password pair,
	// this could still potentially lead to issues.
	defaultMySQLDatabase string = ""

	// defaultMySQLEncryption controls what encryption setting is used by
	// default to connect to a user-provided MySQL host. We default to
	// `preferred` in an effort to provide "security out of the box" for
	// servers that advertise TLS support, but allow non-TLS connections for
	// servers that do not.
	defaultMySQLEncryption string = MySQLEncryptionPreferred

	// defaultMySQLConnMaxLifetime is the maximum lifetime in minutes for
	// MySQL database connections
	defaultMySQLConnMaxLifetime time.Duration = time.Minute * 3

	// defaultMySQLConnMaxIdleTime is the maximum time in minutes that a
	// database connection can remain idle
	defaultMySQLConnMaxIdleTime time.Duration = time.Second * 10

	// defaultMySQLMaxOpenConns is the maximum number of database connections
	// for the MySQL server used by this application
	defaultMySQLMaxOpenConns int = 10

	// defaultMySQLMaxIdleConns is the maximum number of idle database
	// connections for the MySQL server used by this application
	defaultMySQLMaxIdleConns int = 10

	// defaultSQLiteDBFile is the bare filename of the SQLite database
	// created/managed by this application.
	defaultSQLiteDBFile string = "mailserver.db"

	// defaultSQLiteDBPath is the directory where the SQLite database
	// created/managed by this application is stored.
	defaultSQLiteDBPath string = "/var/cache/mysql2sqlite"

	// defaultSQLiteCreateIndexes indicates whether index creation queries are
	// used when generating the SQLite database file.
	defaultSQLiteCreateIndexes bool = false

	// defaultSQLiteBusyTimeout specifies the sleep time used when a table is
	// locked.
	defaultSQLiteBusyTimeout int = 1000

	// not possible to set this? Returning nil from our getter function,
	// relying on validation step to prevent actually returning nil
	// SQLQueries map/type in production run.
	// defaultSQLQueries = nil
)
