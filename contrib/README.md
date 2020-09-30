<!-- omit in toc -->
# mysql2sqlite: Contrib content

<!-- omit in toc -->
## Table of Contents

- [Overview](#overview)
- [Future repo](#future-repo)
- [Index](#index)
- [References](#references)

## Overview

Contributed and example content used in `mysql2sqlite` documentation.

Content in this path is provided as reference material. Along with project
documentation, the goal is to help illustrate how applications provided by
this project can be used to mirror and validate a remote MySQL database
against a generated local SQLite database file.

While the use case of the documentation and example files is to support
Postfix lookup tables, the goal is for the design of the `mysql2sqlite` and
`check_mysql2sqlite` applications to support a number of scenarios.

## Future repo

Some of this content will likely be moved to a separate, dedicated repo in the
future. If that occurs, this README will be updated to reflect that new
location. See <https://github.com/atc0005/mysql2sqlite/issues/4> for more
information.

## Index

You will need to modify *most* of these files before you can reliably (or
safely) use them; it is best to assume that you will need to modify them
before deployment.

| Contrib directory path                       | Deployment path                                               | Description                                                                                                                                                                                            |
| -------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `cron.d/mysql2sqlite`                        | `/etc/cron.d/mysql2sqlite`                                    | crontab fragment that schedules execution of the wrapper script for the `mysql2sqlite` binary.                                                                                                         |
| `logrotate.d/mysql2sqlite`                   | `/etc/logrotate.d/mysql2sqlite`                               | logrotate config fragment. Defines log rotation settings for logs composed of output captured from the `mysql2sqlite` and `check_mysql2sqlite` tools.                                                  |
| `mailserver/sql/mysql_create_db.sql`         | e.g., `mysql -u root -h 127.0.0.1 < mysql_create_db.sql`      | SQL input file for MySQL 5.5.x or MariaDB 10.x and newer. This creates the database schema associated with the `mailserver` database.                                                                  |
| `mailserver/sql/mysql_create_ro_user.sql`    | e.g., `mysql -u root -h 127.0.0.1 < mysql_create_ro_user.sql` | SQL input file for MySQL 5.5.x or MariaDB 10.x and newer. This creates the read-only user account for accessing the `mailserver` database.                                                             |
| `mailserver/sql/mysql_test_data.sql`         | e.g., `mysql -u root -h 127.0.0.1 < mysql_test_data.sql`      | SQL input file for MySQL 5.5.x or MariaDB 10.x and newer. This imports test data associated with the `mailserver` database.                                                                            |
| `postfix/main.cf`                            | Don't. This is an incomplete template.                        | Template primary Postfix configuration file to illustrate use of lookup table config files that reference generated SQLite database file.                                                              |
| `postfix/master.cf`                          | Don't. This is an incomplete template.                        | Template Postfix daemon configuration file. Contains a few entries referenced by primary Postfix config file.                                                                                          |
| `postfix/lookup_tables/sqlite/*.cf`          | `/etc/postfix/lookup_tables/sqlite/*.cf`                      | Template Postfix lookup table config files that reference generated SQLite database file.                                                                                                              |
| `mysql2sqlite/config.example.yaml`           | `/usr/local/etc/mysql2sqlite/config.yaml`                     | Starter configuration file. Used as a starting point for new installations. Should be consulted periodically for updated (or changed) settings.                                                        |
| `mysql2sqlite/check_mysql2sqlite_wrapper.sh` | `/usr/local/sbin/check_mysql2sqlite_wrapper.sh`               | Wrap check_mysql2sqlite Nagios plugin to provide more control over output redirection to specific log file.                                                                                            |
| `mysql2sqlite/mysql2sqlite_wrapper.sh`       | `/usr/local/sbin/mysql2sqlite_wrapper.sh`                     | Wrap mysql2sqlite application executed via cron job to provide more control over output redirection to specific log file.                                                                              |
| `nagios/nrpe.d/mysql2sqlite.cfg`             | `/etc/nagios/nrpe.d/mysql2sqlite.cfg`                         | Nagios Remote Plugin Executor (`NRPE`) config fragment. This file defines the `NRPE` command used to perform service checks as requested by a Nagios console.                                          |
| `sudoers.d/mysql2sqlite`                     | `/etc/sudoers.d/mysql2sqlite`                                 | Allow cron wrapper script called by the mysql2sqlite service account permission to stop, start and (if needed) restart the Postfix daemon as part of regenerating the SQLite database used by Postfix. |

## References

- Inspiration
  - <https://github.com/WhyAskWhy/mysql2sqlite-dev>
  - <https://github.com/WhyAskWhy/mysql2sqlite>
  - <https://workaround.org/ispmail>

- Cron
  - <http://manpages.ubuntu.com/manpages/man5/crontab.5.html>

- Postfix
  - <http://www.postfix.org/DATABASE_README.html>
  - <http://www.postfix.org/sqlite_table.5.html>
  - <https://workaround.org/ispmail>

- logrotate
  - <https://github.com/logrotate/logrotate>
  - <http://manpages.ubuntu.com/manpages/man8/logrotate.8.html>
