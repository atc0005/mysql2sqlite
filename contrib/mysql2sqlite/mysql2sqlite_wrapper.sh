#!/bin/bash

# Copyright 2020 Adam Chalkley
#
# https://github.com/atc0005/mysql2sqlite
#
# Licensed under the MIT License. See LICENSE file in the project root for
# full license information.

# See README.md for deployment path, etc.

# ----------------------------------------------------------------

# Purpose: Wrap mysql2sqlite application executed via cron job
#
# The intent is to better handle redirecting output and setting specific user
# options without needed to modify the crontab config fragment. Run this
# script as the `postfix` user account.
#
# Make sure that the /var/log/mysql2sqlite directory is owned by `postfix`,
# group of mysql2sqlite.
#
# The database directory, `/var/cache/mysql2sqlite`, has owner `mysql2sqlite`,
# group `mysql2sqlite` and permissions `2750` in order to allow any services
# with membership in the `mysql2sqlite` group to read the current database and
# any others created in that path in the future (via setgid bit).
#
# The log directory is owned by user `mysql2sqlite` and group `mysql2sqlite`,
# permissions 0750. The files within are /var/log/mysql2sqlite/nagios.log and
# /var/log/mysql2sqlite/cron.log. Each are owned by user `mysql2sqlite` and
# group `mysql2sqlite`. Because the Nagios plugin runs as the `nagios` user
# account and the `nagios` user account is in the group `mysql2sqlite`, group
# `mysql2sqlite` has write permission to the log file.

our_config_file="/usr/local/etc/mysql2sqlite/config.yaml"
our_log_file="/var/log/mysql2sqlite/cron.log"
our_log_file_format="logfmt"
our_log_output_target="stderr"

# Add a slight delay to work around any potential systemd sensitivity to a
# "too quick" stop, start of Postfix.
our_delay_in_seconds="1"

# We don't want the output intended for the Nagios "console", so we toss it by
# redirecting it here.
our_nagios_plugin_output_target="/dev/null"

# Check to see if local SQLite database is out of sync with remote MySQL
# database. Only regenerate the local SQLite database if it is out of sync,
# otherwise, leave it as is.
#
# We toss the Nagios-specific stdout content as it duplicates details already
# being saved to the log file (captured stderr output).
echo "Calling check_mysql2sqlite to check if SQLite db regen is needed." >> "${our_log_file}"
/usr/lib/nagios/plugins/check_mysql2sqlite \
    --config-file "${our_config_file}" \
    --log-format "${our_log_file_format}" \
    --log-output "${our_log_output_target}" 1> "${our_nagios_plugin_output_target}" 2>> "${our_log_file}"

# Capture exit code from Nagios plugin. This code will be used to determine if
# a SQLite db regen is needed.
dbRegenCheckResult=$?

case ${dbRegenCheckResult} in
    0)
        echo "Skipping SQLite db regen; regen not needed per check_mysql2sqlite validation results (return code ${dbRegenCheckResult})." >> "${our_log_file}"
        ;;

    1)
        # Rely on exit code 1 (WARNING) to determine need for regen.
        echo "Regenerating SQLite db per check_mysql2sqlite validation results (return code ${dbRegenCheckResult})." >> "${our_log_file}"

        # Stop Postfix (per upstream advice) before regenerating SQLite
        # database that it depends on.
        echo "Stopping Postfix." >> "${our_log_file}"
        sudo /bin/systemctl stop postfix
        stopResult=$?
        if [ ${stopResult} -eq 0 ]
        then
            echo "OK: Postfix stopped successfully; exit code ${stopResult} returned." >> "${our_log_file}"
        else
            echo "FAIL: Postfix failed to stop; exit code ${stopResult} returned." >> "${our_log_file}"
            echo "FAIL: Exiting $0 in an effort to prevent making the situation any worse."
            exit ${stopResult}
        fi

        echo "Delaying ${our_delay_in_seconds} seconds before regenerating SQLite db." >> "${our_log_file}"
        sleep ${our_delay_in_seconds}

        echo "Regenerating SQLite database file." >> "${our_log_file}"
        /usr/local/sbin/mysql2sqlite \
            --config-file "${our_config_file}" \
            --log-format "${our_log_file_format}" 2>&1 >> "${our_log_file}"

        sqliteDBRegenResult=$?
        if [ ${sqliteDBRegenResult} -eq 0 ]
        then
            echo "OK: SQLite database regen successful; exit code ${sqliteDBRegenResult} returned." >> "${our_log_file}"
        else
            echo "FAIL: SQLite database reged failed; exit code ${sqliteDBRegenResult} returned." >> "${our_log_file}"
        fi

        # Unconditionally start Postfix, regardless of the outcome of the
        # SQLite database regen work. Nagios should notice any ill effects
        # from a failed regen and let us know for further investigation; we'll
        # use the captured log messages to assist with that.
        echo "Starting Postfix." >> "${our_log_file}"
        sudo /bin/systemctl start postfix
        startResult=$?
        if [ ${startResult} -eq 0 ]
        then
            echo "OK: Postfix started successfully; exit code ${startResult} returned." >> "${our_log_file}"
        else
            echo "FAIL: Postfix failed to start; exit code ${startResult} returned." >> "${our_log_file}"
        fi
        ;;

    *)
        echo "Unhandled result from check_mysql2sqlite; exit code ${dbRegenCheckResult} returned." >> "${our_log_file}"
        ;;

esac
