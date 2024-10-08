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

plugin_name="check_mysql2sqlite"
plugin_path="/usr/lib/nagios/plugins"

echo
echo "Thank you for installing packages provided by the ${project_fq_name} project!"
echo
echo "Project resources:"
echo
echo "- Obtain latest release: ${project_releases}"
echo "- View/Ask questions: ${project_discussions}"
echo "- View/Open issues: ${project_issues}"
echo
