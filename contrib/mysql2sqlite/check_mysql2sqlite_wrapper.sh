#!/bin/bash

# Copyright 2020 Adam Chalkley
#
# https://github.com/atc0005/mysql2sqlite
#
# Licensed under the MIT License. See LICENSE file in the project root for
# full license information.

# See README.md for deployment path, etc.

# ----------------------------------------------------------------

# Purpose: Wrap check_mysql2sqlite Nagios plugin
#
# The intent is to better handle redirecting output. stdout from
# `check_mysql2sqlite` should go to stdout and stderr should be sent to the
# specified log file, all as the same user.

# Since Nagios considers anything sent to stdout to be intended for its
# review, we set log output to stderr. Since Nagios ignores stderr output,
# sending log output there allows us to potentially redirect that output to a
# log file later if we so choose.
#
# The database directory, `/var/cache/mysql2sqlite`, has owner `mysql2sqlite`,
# group `mysql2sqlite` and permissions `2750` in order to allow any services
# with membership in the `mysql2sqlite` group to read the current database and
# any others created in that path in the future (via setgid bit).
#
# The log directory and files are owned by user `mysql2sqlite` and group
# `mysql2sqlite`, permissions 0750.

/usr/lib/nagios/plugins/check_mysql2sqlite \
    --config-file /usr/local/etc/mysql2sqlite/config.yaml \
    --log-format logfmt \
    --log-output stderr 2>> /var/log/mysql2sqlite/nagios.log
