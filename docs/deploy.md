<!-- omit in toc -->
# mysql2sqlite: Deployment

- [Project README](../README.md)

<!-- omit in toc -->
## Table of contents

- [Overview](#overview)
- [Assumptions](#assumptions)
- [Privileges required](#privileges-required)
  - [Service accounts](#service-accounts)
  - [Ownership / permissions](#ownership--permissions)
- [Directions](#directions)
  - [Create service account and group](#create-service-account-and-group)
  - [Add existing service accounts to new group](#add-existing-service-accounts-to-new-group)
  - [Prepare SQLite database cache](#prepare-sqlite-database-cache)
  - [Prepare log directory, files](#prepare-log-directory-files)
  - [Deploy binaries](#deploy-binaries)
  - [Deploy configuration file](#deploy-configuration-file)
  - [Deploy wrapper scripts](#deploy-wrapper-scripts)
  - [Deploy sudoers config fragment](#deploy-sudoers-config-fragment)
  - [Deploy logrotate config fragment](#deploy-logrotate-config-fragment)
  - [Deploy cron config fragment](#deploy-cron-config-fragment)
  - [Deploy Nagios "client" configuration file](#deploy-nagios-client-configuration-file)
  - [Postfix](#postfix)
    - [Install Support packages](#install-support-packages)
    - [Deploy lookup table config fragments](#deploy-lookup-table-config-fragments)
    - [Restart Postfix to apply changes](#restart-postfix-to-apply-changes)

## Overview

While you can run the `mysql2sqlite` binary directly as a one-off
synchronization job, the most common use case will be running the
`mysql2sqlite` binary from a cron job and the `check_mysql2sqlite` binary via
the Nagios Remote Plugin Executor (`NRPE`).

This deployment document walks you through deploying both applications
provided by this project in order to synchronize a remote MySQL database named
`mailserver` will be synchronized to a local SQLite database of the same
overall name: `mailserver.db`.

In our example, this database is used by Postfix for various lookup tables as
part of its normal operation. Substitute this database filename for whatever
database you wish to generate/maintain using `mysql2sqlite`.

## Assumptions

Various assumptions are made in this documentation. These assumptions are made
in order to provide a more complete example that illustrates how binaries
provided by this project may be used. Modify as you feel appropriate.

## Privileges required

These directions illustrate creating a "service account" to run the
`mysql2sqlite` application as a dedicated user. This application does not have
complex requirements and does not require elevated privileges; running as the
root user account is not needed nor advised.

### Service accounts

These directions also assume that the `postfix` service account will need to
read this database and uses its membership in the `mysql2sqlite` group to
accomplish this. Substitute this also by specifying the name of an applicable
service account that will need to access the SQLite database file that you
generate.

### Ownership / permissions

The `mysql2sqlite` service account is used to run the `mysql2sqlite` binary
from a cron job in order to regenerate a local SQLite database. This service
account requires read/write access to the SQLite database and the cron log
file.

A wrapper script is used to run the `check_mysql2sqlite` Nagios plugin in
order to validate that the local database matches the remote database. Only
when there is a discrepancy is the `mysql2sqlite` binary run to synchronize
both databases.

All output is captured and redirected with desired log output going to a
dedicated "cron" log file.

The `nagios` (or `nrpe`) service account is used to run the
`check_mysql2sqlite` binary from the Nagios Remote Plugin Executor (`NRPE`) in
order to validate the local SQLite database against the source MySQL database.

This plugin is also run via a wrapper script. When run, a specific flag is
passed that directs log output to `stderr` which the wrapper script redirects
to a "nagios" log file specific to the plugin. `stdout` output is not
redirected, passing on to the Nagios Remote Plugin Executor (`NRPE`) and back
to the Nagios console which initiated the check.

## Directions

### Create service account and group

1. `sudo useradd mysql2sqlite --system --shell /usr/sbin/nologin --user-group`
1. `echo "mysql2sqlite:$(python3 -c 'import uuid; print(uuid.uuid4())')" | sudo chpasswd`
   - this is optional and dependent upon a system installation of Python 3
   - replace `python3` with `python` if applicable

### Add existing service accounts to new group

For this example, we'll assume that these user accounts exist:

| Software                               | User Account | Notes        |
| -------------------------------------- | ------------ | ------------ |
| Postfix                                | `postfix`    |              |
| Nagios Remote Plugin Executor (`NRPE`) | `nagios`     | Debian-based |
| Nagios Remote Plugin Executor (`NRPE`) | `nrpe`       | RedHat-based |

1. `sudo usermod --groups mysql2sqlite postfix`
1. `sudo usermod --groups mysql2sqlite nagios`
   - substitute with `nrpe` if applicable

### Prepare SQLite database cache

This sets up the directory to hold the database and creates a placeholder file
with the desired permissions.

1. `sudo mkdir -vp /var/cache/mysql2sqlite`
1. `sudo touch /var/cache/mysql2sqlite/mailserver.db`
1. `sudo chown -Rv mysql2sqlite:mysql2sqlite /var/cache/mysql2sqlite`
1. `sudo chmod -v 2750 /var/cache/mysql2sqlite`
   - this enables `setgid` in order to allow the `mysql2sqlite` group to be
     inherited for any content created within this path
1. `sudo chmod -v 0640 /var/cache/mysql2sqlite/mailserver.db`

### Prepare log directory, files

This sets up the log directory and creates placeholder log files with the
desired permissions.

1. Create directory, placeholder files
   1. `sudo mkdir -vp /var/log/mysql2sqlite`
   1. `sudo touch /var/log/mysql2sqlite/nagios.log`
   1. `sudo touch /var/log/mysql2sqlite/cron.log`
1. Set ownership, permissions
   1. `sudo chown -Rv mysql2sqlite:mysql2sqlite /var/log/mysql2sqlite`
   1. `sudo chmod -v 0660 /var/log/mysql2sqlite/nagios.log`
      - the `nagios` or `nrpe` service account is a member of the
        `mysql2sqlite` group
   1. `sudo chmod -v 0640 /var/log/mysql2sqlite/cron.log`

### Deploy binaries

Continuing where the [build](build.md) instructions left off, follow the steps
below to deploy the generated `check_mysql2sqlite` and `mysql2sqlite` binaries
to locations on a system that will run them.

1. Deploy `check_mysql2sqlite`
   - as `/usr/lib/nagios/plugins/check_mysql2sqlite` on Debian-based systems
   - as `/usr/lib64/nagios/plugins/check_mysql2sqlite` on RedHat-based systems
     - note: don't forget about SELinux contexts!
1. Deploy `mysql2sqlite`
   - as `/usr/local/sbin/mysql2sqlite`
1. Set execute bit on `mysql2sqlite` binary
   1. `sudo chmod -v +x /usr/local/sbin/mysql2sqlite`
1. Set execute bit on `check_mysql2sqlite` binary
   - `sudo chmod -v +x /usr/lib/nagios/plugins/check_mysql2sqlite`
   - `sudo chmod -v +x /usr/lib64/nagios/plugins/check_mysql2sqlite`

### Deploy configuration file

This sets up a directory to hold the `mysql2sqlite` and `check_mysql2sqlite`
configuration file, copies the file and sets permissions to prevent casual
read access to the file.

1. Create directory
   1. `sudo mkdir -vp /usr/local/etc/mysql2sqlite`
1. Deploy configuration file
   - template: `contrib/mysql2sqlite/config.example.yaml`
   - destination: `/usr/local/etc/mysql2sqlite/config.yaml`
1. Modify configuration file
   - see [configure](configure.md) doc for details
1. Set ownership, permissions
   1. `sudo chown -v root:mysql2sqlite
      /usr/local/etc/mysql2sqlite/config.yaml`
   1. `sudo chmod -v 0640 /usr/local/etc/mysql2sqlite/config.yaml`
   - this file contains username/password pair, so it is important to restrict
     access to only user/groups that need to use it

### Deploy wrapper scripts

1. Deploy `mysql2sqlite_wrapper.sh`
   - as `/usr/local/sbin/mysql2sqlite_wrapper.sh`
1. Deploy `check_mysql2sqlite_wrapper.sh`
   - as `/usr/local/sbin/check_mysql2sqlite_wrapper.sh`
1. Set ownership, permissions
   1. `sudo chown -v root:root /usr/local/sbin/mysql2sqlite_wrapper.sh`
   1. `sudo chmod -v +x /usr/local/sbin/mysql2sqlite_wrapper.sh`
   1. `sudo chown -v root:root /usr/local/sbin/check_mysql2sqlite_wrapper.sh`
   1. `sudo chmod -v +x /usr/local/sbin/check_mysql2sqlite_wrapper.sh`

### Deploy sudoers config fragment

1. Deploy `sudoers.d/mysql2sqlite`
   - as `/etc/sudoers.d/mysql2sqlite`
1. Set ownership, permissions
   1. `sudo chown -v root:root /etc/sudoers.d/mysql2sqlite`
   1. `sudo chmod -v 0400 /etc/sudoers.d/mysql2sqlite`

### Deploy logrotate config fragment

1. Deploy `logrotate.d/mysql2sqlite`
   - as `/etc/logrotate.d/mysql2sqlite`
1. Set ownership, permissions
   1. `sudo chown -v root:root /etc/logrotate.d/mysql2sqlite`
   1. `sudo chmod -v 644 /etc/logrotate.d/mysql2sqlite`

### Deploy cron config fragment

1. Deploy `cron.d/mysql2sqlite`
   - as `/etc/cron.d/mysql2sqlite`
1. Set ownership, permissions
   1. `sudo chown -v root:root /etc/cron.d/mysql2sqlite`
   1. `sudo chmod -v 644 /etc/cron.d/mysql2sqlite`

### Deploy Nagios "client" configuration file

The Nagios Remote Plugin Executor (`NRPE`) configuration fragment is deployed
to the same system where the `mysql2sqlite` and `check_mysql2sqlite` binaries
will be executed.

1. Deploy `nagios/nrpe.d/mysql2sqlite.cfg`
   - as `/etc/nagios/nrpe.d/mysql2sqlite.cfg`
1. Set ownership, permissions
   1. `sudo chown -v root:root /etc/nagios/nrpe.d/mysql2sqlite.cfg`
   1. `sudo chmod -v 644 /etc/nagios/nrpe.d/mysql2sqlite.cfg`
1. Restart the `nrpe` daemon
   - Debian-based: `sudo service nagios-nrpe-server restart`
   - RedHat-based: `sudo service nrpe restart`

### Postfix

#### Install Support packages

These directions assume that an existing Postfix installation is present,
likely with MySQL support packages already installed. At this point the next
step is to install the support packages necessary for Postfix to query SQLite
databases for lookup table information.

1. `sudo apt-get update`
1. `sudo apt-get install postfix-mysql`
   - on the chance that the package is not already present
1. `sudo apt-get install postfix-sqlite`

#### Deploy lookup table config fragments

Once the Postfix SQLite support package is installed, we need to deploy the
lookup table config fragments which instruct Postfix how to find the database
and what queries to use to obtain lookup table information. As is the case for
the rest of these directions, the exact steps needed for your environment will
vary.

1. `sudo mkdir -vp /etc/postfix/lookup_tables/sqlite`
1. Deploy all `postfix/lookup_tables/sqlite/*.cf` files to
   `/etc/postfix/lookup_tables/sqlite/`

#### Restart Postfix to apply changes

1. `sudo systemctl restart postfix`
