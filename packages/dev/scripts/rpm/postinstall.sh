#!/bin/bash

# Copyright 2023 Adam Chalkley
#
# https://github.com/atc0005/mysql2sqlite
#
# Licensed under the MIT License. See LICENSE file in the project root for
# full license information.

project_org="atc0005"
project_shortname="mysql2sqlite"

project_fq_name="${project_org}/${project_shortname}"
project_url_base="https://github.com/${project_org}"
project_repo="${project_url_base}/${project_shortname}"
project_releases="${project_repo}/releases"
project_issues="${project_repo}/issues"
project_discussions="${project_repo}/discussions"

plugin_name="check_mysql2sqlite_dev"
plugin_path="/usr/lib64/nagios/plugins"

# Set required SELinux context to allow plugin use when SELinux is enabled.
if [ -f "${plugin_path}/${plugin_name}" ]; then

    # Make sure we can locate the selinuxenabled binary.
    if [ -x "$(command -v selinuxenabled)" ]; then
        selinuxenabled

        if [ $? -ne 0 ]; then
            echo -e "\nSELinux is not enabled, skipping application of contexts."
        else
            # SELinux is enabled. Set context.
            echo -e "\nApplying SELinux contexts on ${plugin_path}/${plugin_name} ..."
            restorecon -v ${plugin_path}/${plugin_name}

            if [ $? -eq 0 ]; then
                echo "Successfully applied SELinux contexts on ${plugin_path}/${plugin_name}"
            else
                echo "Failed to set SELinux contexts on ${plugin_path}/${plugin_name}"
            fi
        fi

    else
        echo "Error: Failed to locate selinuxenabled command." >&2
    fi

else
    echo "${plugin_path}/${plugin_name} could not be found!"
fi

echo
echo "Thank you for installing packages provided by the ${project_fq_name} project!"
echo
echo "#######################################################################"
echo "NOTE:"
echo
echo "This is a dev build; binaries installed by this package have a _dev"
echo "suffix to allow installation alongside stable versions."
echo
echo "Feedback for all releases is welcome, but especially so for dev builds."
echo "Thank you in advance!"
echo "#######################################################################"
echo
echo "Project resources:"
echo
echo "- Obtain latest release: ${project_releases}"
echo "- View/Ask questions: ${project_discussions}"
echo "- View/Open issues: ${project_issues}"
echo
