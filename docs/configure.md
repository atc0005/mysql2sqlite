<!-- omit in toc -->
# mysql2sqlite: Configuration

- [Project README](../README.md)

<!-- omit in toc -->
## Table of contents

- [Precedence](#precedence)
- [Command-line Arguments](#command-line-arguments)
- [Environment Variables](#environment-variables)
- [Configuration File](#configuration-file)
  - [Settings](#settings)
  - [Starter/Example file](#starterexample-file)
- [Worth noting](#worth-noting)
  - [Automatic config file detection](#automatic-config-file-detection)
  - [Log formats](#log-formats)

## Precedence

The priority order is:

1. Command line flags (highest priority)
1. Environment variables
1. Configuration file
1. Default settings (lowest priority)

A limited number of options are supported via command-line flags or
environment variables in order to "bootstrap" the applications provided by
this project. A configuration file is required to set all other required
values.

To be clear, it is not possible to use the applications provided by this
project without supplying a configuration file. A [starter config
file](../contrib/README.md) is provided, but it will need to be modified
before it will be useful.

## Command-line Arguments

- Flags marked as **`required`** must be set via CLI flag *or* within the
  YAML-formatted configuration file.
- Flags *not* marked as required are for settings where a useful default is
  already defined.

| Option        | Required                 | Default        | Repeat | Possible                                   | Description                                                                                                                      |
| ------------- | ------------------------ | -------------- | ------ | ------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------- |
| `h`, `help`   | No                       | `false`        | No     | `h`, `help`                                | Show Help text along with the list of supported flags.                                                                           |
| `config-file` | [*Maybe*](#worth-noting) | *empty string* | No     | *valid path to a file*                     | Fully-qualified path to a configuration file consulted for settings not already provided via CLI flags or environment variables. |
| `log-level`   | No                       | `info`         | No     | `fatal`, `error`, `warn`, `info`, `debug`  | Log message priority filter. Log messages with a lower level are ignored.                                                        |
| `log-output`  | No                       | `stdout`       | No     | `stdout`, `stderr`                         | Log messages are written to this output target.                                                                                  |
| `log-format`  | No                       | `text`         | No     | `cli`, `json`, `logfmt`, `text`, `discard` | Use the specified `apex/log` package "handler" to output log messages in that handler's format.                                  |

## Environment Variables

If set, environment variables override settings provided by a configuration
file. If used, command-line arguments override the equivalent environment
variables listed below. See the [Command-line
Arguments](#command-line-arguments) table for more information.

| Flag Name     | Environment Variable Name  | Notes | Example (mostly using default values)                                |
| ------------- | -------------------------- | ----- | -------------------------------------------------------------------- |
| `config-file` | `MYSQL2SQLITE_CONFIG_FILE` |       | `MYSQL2SQLITE_CONFIG_FILE="/usr/local/etc/mysql2sqlite/config.yaml"` |
| `log-level`   | `MYSQL2SQLITE_LOG_LEVEL`   |       | `MYSQL2SQLITE_LOG_LEVEL="info"`                                      |
| `log-output`  | `MYSQL2SQLITE_LOG_OUTPUT`  |       | `MYSQL2SQLITE_LOG_OUTPUT="stdout"`                                   |
| `log-format`  | `MYSQL2SQLITE_LOG_FORMAT`  |       | `MYSQL2SQLITE_LOG_FORMAT="text"`                                     |

## Configuration File

### Settings

Configuration file settings have the lowest priority and are overridden by
settings specified in other configuration sources, except for default values.
See the [Command-line Arguments](#command-line-arguments) table for more
information, including the available values for the listed configuration
settings.

| Setting Name (subkey)              | Section Name (parent key) | Default                   | Possible                                                                                 | Description                                                                                                                                                                                                                                   | Notes                                                            |
| ---------------------------------- | ------------------------- | ------------------------- | ---------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| `level`                            | `logging`                 | `info`                    | `fatal`, `error`, `warn`, `info`, `debug`                                                | Log message priority filter. Log messages with a lower level are ignored.                                                                                                                                                                     |                                                                  |
| `format`                           | `logging`                 | `text`                    | `cli`, `json`, `logfmt`, `text`, `discard`                                               | Log messages are written to this output target.                                                                                                                                                                                               |                                                                  |
| `output`                           | `logging`                 | `stdout`                  | `stdout`, `stderr`                                                                       | Use the specified `apex/log` package "handler" to output log messages in that handler's format.                                                                                                                                               |                                                                  |
| `connection_stats`                 | `logging`                 | `false`                   | `true`, `false`                                                                          | Whether connection statistics are logged during application execution. This can be useful for troubleshooting, but is probably too much information for normal operation.                                                                     |                                                                  |
| `trim_leading_trailing_whitespace` | `general`                 | `false`                   | `true`, `false`                                                                          | If enabled, this setting trims all leading and trailing whitespace from values retrieved from the MySQL database before writing to the SQLite database.                                                                                       |                                                                  |
| `connection_retries`               | `general`                 | `5`                       | *number of retries as a valid whole number*                                              | The number of times a connection should be retried before giving up and returning an error.                                                                                                                                                   |                                                                  |
| `connection_retry_delay`           | `general`                 | `2`                       | *number of seconds as a whole number*                                                    | The number of seconds between connection retry attempts.                                                                                                                                                                                      |                                                                  |
| `connection_timeout`               | `general`                 | `15`                      | *number of seconds as a whole number*                                                    | The number of seconds a connection should be tried before giving up and returning an error.                                                                                                                                                   |                                                                  |
| `username`                         | `mysql_config`            | *empty string*            | *valid MySQL user account name with sufficient access to database*                       | MySQL user account name with sufficient access to database.                                                                                                                                                                                   |                                                                  |
| `password`                         | `mysql_config`            | *empty string*            | *valid MySQL user account password for account used to read database*                    | Password for the MySQL database account. Read-only access is sufficient.                                                                                                                                                                      |                                                                  |
| `host`                             | `mysql_config`            | *empty string*            | *valid fqdn, IP Address or UNIX socket*                                                  | MySQL database host specified as IP Address, FQDN or UNIX socket. If you use a stunnel or other forwarded connection you will likely wish to set this to 127.0.0.1 and override the default port setting.                                     |                                                                  |
| `port`                             | `mysql_config`            | `3306`                    | *valid TCP port number*                                                                  | If the applications from this project run on the same box as the MySQL/MariaDB server then you probably want to enter 3306 here, otherwise if using stunnel or another port forwarding setup you will likely need to enter an alternate port. | This value is ignored if you specify a UNIX socket for the host. |
| `database`                         | `mysql_config`            | *empty string*            | *valid MySQL database name*                                                              | The database whose contents you wish to mirror to a local SQLite database.                                                                                                                                                                    |                                                                  |
| `encryption`                       | `mysql_config`            | `preferred`               | `true`, `false`, `skip-verify`, `preferred`                                              | Driver-specific TLS connection settings.                                                                                                                                                                                                      | [Docs](https://github.com/go-sql-driver/mysql#tls)               |
| `max_connection_lifetime`          | `mysql_config`            | `3`                       | *number of minutes as a whole number*                                                    | Maximum connection lifetime.                                                                                                                                                                                                                  |                                                                  |
| `max_open_connections`             | `mysql_config`            | `10`                      | *valid whole number less than maximum allowed by remote db host*                         | Maximum number of connections allowed to the source MySQL database for this application.                                                                                                                                                      |                                                                  |
| `max_idle_connections`             | `mysql_config`            | `10`                      | *valid whole number less than maximum allowed by remote db host*                         | Maximum number of idle connections allowed to the source MySQL database. Recommended to be set the same (or greater than) max_open_connections.                                                                                               |                                                                  |
| `max_idle_connection_time`         | `mysql_config`            | `10`                      | *valid whole number less than maximum allowed by remote db host*                         | Maximum amount of time a connection may be idle. Expired connections may be closed lazily before reuse.                                                                                                                                       |                                                                  |
| `busy_timeout`                     | `sqlite_config`           | `1000`                    | *number of milliseconds as a whole number*                                               | Sleep time used when a table is locked. If not zero, multiple sleep attempts are made until at least the specified value of milliseconds of sleeping have accumulated. After the specified value of milliseconds, `SQLITE_BUSY` is returned.  |                                                                  |
| `journal_mode`                     | `sqlite_config`           | `DELETE`                  | *valid [SQLite journaling mode](https://www.sqlite.org/pragma.html#pragma_journal_mode)* | Sets the journal mode for the SQLite database connection.                                                                                                                                                                                     |                                                                  |
| `db_filename`                      | `sqlite_config`           | `mailserver.db`           | *valid, non-qualified filename*                                                          | The "short" name of the SQLite database file. Do not specify the directory which will hold the file here.                                                                                                                                     |                                                                  |
| `base_dir`                         | `sqlite_config`           | `/var/cache/mysql2sqlite` | *valid, qualified directory name*                                                        | The directory that will hold the SQLite database file. The user running this script should have write access to this location. Make sure that applicable daemons (e.g., Postfix, Dovecot, etc) have read access to the file.                  |                                                                  |
| `create_indexes`                   | `sqlite_config`           | `false`                   | `true`, `false`                                                                          | Whether an "index" query will be required for each table query set defined within this file. If enabled, this application will use the "index" query to generate indexes when creating tables in the SQLite database file.                    |                                                                  |

The `queries` map format is described in detail below.

| Key (setting) name                 | Parent key          | Possible                                                                    | Description                                                                                                                                                                                                                    | Notes                                                                                                                      |
| ---------------------------------- | ------------------- | --------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------- |
| `queries`                          | N/A (top-level key) | a subkey named after a source MySQL table name                              | The `queries` key holds subkeys based on source MySQL table names.                                                                                                                                                             | This parent key name is static and should not be changed.                                                                  |
| `virtual_users` (see notes column) | `queries`           | subkeys of `read`, `new`, `write`; `index` also if `create_indexes` enabled | The `virtual_users` subkey represents a table named `virtual_users` in the source MySQL database. This key holds subkeys which provide the SQL queries used to synchronize this source MySQL table to a local SQLite database. | The `virtual_users` table name is an example. It should be substituted with a table name that matches your MySQL database. |
| `read`                             | `virtual_users`     | *valid SQL query*                                                           | The `read` subkey holds a SQL query used to read all rows from the MySQL database table.                                                                                                                                       | This subkey name is required and should not be changed.                                                                    |
| `new`                              | `virtual_users`     | *valid SQL query*                                                           | The `new` subkey holds a SQL query used to create the local SQLite database table.                                                                                                                                             | This subkey name is required and should not be changed.                                                                    |
| `write`                            | `virtual_users`     | *valid SQL query*                                                           | The `write` subkey holds a SQL query used to write a row to the local SQLite database table.                                                                                                                                   | This subkey name is required and should not be changed.                                                                    |
| `index`                            | `virtual_users`     | *valid SQL query*                                                           | The `index` subkey holds a SQL query used to create a new index for the local SQLite database table.                                                                                                                           | This subkey name is required **if** the `create_indexes` setting is enabled. This subkey name should not be changed.       |

In short, this is the expected format:

```yaml
---
general:
# ...

mysql_config:
# ...

sqlite_config:
# ...

logging:
# ...

queries:
  table_name:
    read: "Query used to read all rows from the MySQL database table"
    new: "Query used to create the local SQLite database table"
    write: "Query used to write a row to the local SQLite database table"
    index: "Query used to create a new index for the local SQLite database table"
```

### Starter/Example file

The [`config.example.yaml`](../contrib/README.md) file is provided as a
starting point for your own configuration file. The default values provided by
this configuration file should lineup with the default application values if
not specified. Please [file an
issue](https://github.com/atc0005/mysql2sqlite/issues/new) if you find that is
not the case.

Once modified, your copy of the configuration file can be placed in a location
of your choosing and referenced using the `--config-file` flag. See the
[Command-line arguments](#command-line-arguments) sections for usage details.

If the file cannot be located, the application will note this and then abort.

## Worth noting

### Automatic config file detection

If not specified, applications provided by this project will attempt to
automatically locate the configuration file by looking in the following
locations:

1. user-specified path to config file
   - e.g., `/usr/local/etc/mysql2sqlite/config.yaml`
   - e.g., `C:\mysql2sqlite\config.yaml`
1. config file found alongside the executable
   - e.g., `/home/username/Desktop/mysql2sqlite/config.yaml`
1. config file found in the user's configuration path
   - e.g., `/home/username/.config/mysql2sqlite/config.yaml`
   - e.g., `C:\Users\username\AppData\Roaming\mysql2sqlite\config.yaml`

If the file cannot be located, the application will note this and then abort.

### Log formats

Log format names map directly to the Handlers provided by the `apex/log`
package. Their descriptions are copied from the [official
README](https://github.com/apex/log/blob/master/Readme.md) and provided below
for reference:

| Log Format ("Handler") | Description                        |
| ---------------------- | ---------------------------------- |
| `cli`                  | human-friendly CLI output          |
| `json`                 | provides log output in JSON format |
| `logfmt`               | plain-text logfmt output           |
| `text`                 | human-friendly colored output      |
| `discard`              | discards all logs                  |
