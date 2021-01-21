<!-- omit in toc -->
# mysql2sqlite: Wrap-up

- [Project README](../README.md)

<!-- omit in toc -->
## Table of contents

### Generate SQLite database file

After [deploying](deploy.md) and [configuring](configure.md) the applications,
the next step is the first manual run of the `mysql2sqlite` binary in order to
generate the `mailserver.db` SQLite database file. From there, the wrapper
scripts described previously can be used to monitor the database and keep it
in sync.

**NOTE**: The database must be created manually *at least once*, otherwise a
loop of sorts will occur.

This loop occurs because the wrapper script used to generate the SQLite
database requires either an `OK` or `WARNING` state as a response from the
Nagios plugin. If a `CRITICAL` state is returned (e.g., the database file is
missing or empty) the `mysql2sqlite` binary is never called. This occurs most
often on a system where `mysql2sqlite` has not previously been used.

The workaround is to run the `mysql2sqlite` binary once manually to bootstrap
the process.

1. `sudo -u mysql2sqlite /usr/local/sbin/mysql2sqlite --config-file
   /usr/local/etc/mysql2sqlite/config.yaml --log-format logfmt`
