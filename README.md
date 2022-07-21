
<!-- omit in toc -->
# mysql2sqlite

Mirror MySQL database tables to local SQLite database.

[![Latest Release](https://img.shields.io/github/release/atc0005/mysql2sqlite.svg?style=flat-square)](https://github.com/atc0005/mysql2sqlite/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/mysql2sqlite.svg)](https://pkg.go.dev/github.com/atc0005/mysql2sqlite)
[![Validate Codebase](https://github.com/atc0005/mysql2sqlite/workflows/Validate%20Codebase/badge.svg)](https://github.com/atc0005/mysql2sqlite/actions?query=workflow%3A%22Validate+Codebase%22)
[![Validate Docs](https://github.com/atc0005/mysql2sqlite/workflows/Validate%20Docs/badge.svg)](https://github.com/atc0005/mysql2sqlite/actions?query=workflow%3A%22Validate+Docs%22)
[![Lint and Build using Makefile](https://github.com/atc0005/mysql2sqlite/workflows/Lint%20and%20Build%20using%20Makefile/badge.svg)](https://github.com/atc0005/mysql2sqlite/actions?query=workflow%3A%22Lint+and+Build+using+Makefile%22)
[![Quick Validation](https://github.com/atc0005/mysql2sqlite/workflows/Quick%20Validation/badge.svg)](https://github.com/atc0005/mysql2sqlite/actions?query=workflow%3A%22Quick+Validation%22)

<!-- omit in toc -->
## Table of Contents

- [Project home](#project-home)
- [Overview](#overview)
  - [`mysql2sqlite`](#mysql2sqlite)
  - [`check_mysql2sqlite`](#check_mysql2sqlite)
- [Features](#features)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [Documentation](#documentation)
  - [Assumptions](#assumptions)
  - [Build applications](#build-applications)
  - [Deploy applications](#deploy-applications)
  - [Configure applications](#configure-applications)
  - [Wrap-up](#wrap-up)
- [License](#license)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

This repo contains tools used to mirror a remote MySQL database to a local
SQLite database. Nearly all options are controlled via a YAML-formatted
configuration file.

| Tool Name            | Description                                                                                           |
| -------------------- | ----------------------------------------------------------------------------------------------------- |
| `mysql2sqlite`       | CLI app used to mirror a remote MySQL database to a local SQLite database.                            |
| `check_mysql2sqlite` | Nagios plugin used to validate synchronization status between remote MySQL and local SQLite database. |

### `mysql2sqlite`

CLI app used to mirror a remote MySQL database to a local SQLite database.
This application can be run as a one-off task or via a cron job or other
automated means.

### `check_mysql2sqlite`

Nagios plugin used to validate synchronization status between a remote MySQL
database and a local target SQLite database.

The output for this application is designed to provide the one-line summary
needed by Nagios for quick identification of a problem while providing longer,
more detailed information for use in email and Teams notifications
([atc0005/send2teams](https://github.com/atc0005/send2teams)).

## Features

- CLI tool for mirroring MySQL database tables to SQLite database
- Nagios plugin for validating mirrored SQLite database against the original
  source
- Configurable source database, destination database settings
- Configurable connection retry, retry delay behavior
- Configurable MySQL settings
  - max open connections
  - max idle connections
  - max connection lifetime
  - max idle connection time
- Configurable SQLite settings
  - busy timeout
  - journal mode
- Configurable logging settings
  - level
  - output "target" (`stdout`, `stderr`)
  - format

See the ([configuration](docs/configure.md)) documentation for all supported
settings.

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

These requirements are specific to Debian/Ubuntu-based distros. Packages will
likely be named differently for other distributions.

- Go
  - see this project's `go.mod` file for *preferred* version
  - this project tests against [officially supported Go
    releases][go-supported-releases]
    - the most recent stable release (aka, "stable")
    - the prior, but still supported release (aka, "oldstable")
- `CGO_ENABLED=1` environment variable (if not set by default)
  - requirement of SQLite database driver used
- `GCC`
- `GCC multilib`
- `GCC for Windows` (`mingw-w64`)
- `make`
  - if using the provided `Makefile`

See the [build](docs/build.md) instructions for more information.

### Running

- Windows 10
- Ubuntu Linux 18.04+

See official [Go install notes][go-docs-install] for specific operating
systems supported.

## Documentation

### Assumptions

Various assumptions are made in this documentation. These assumptions are made
in order to provide a more complete example that illustrates how binaries
provided by this project may be used. Modify as you feel appropriate.

### Build applications

See the [build](docs/build.md) instructions for more information.

As an alternative to building the binaries yourself, this project also
periodically provides binaries via new releases. If binaries for your platform
are not provided, please [file an
issue](https://github.com/atc0005/mysql2sqlite/issues/new) so that we may
evaluate the requirements for providing those binaries with future releases.

### Deploy applications

See the [deployment](docs/deploy.md) documentation for details.

### Configure applications

See the [configure](docs/configure.md) doc for details.

### Wrap-up

See the [wrap-up](docs/wrap-up.md) doc for remaining steps.

## License

From the [LICENSE](LICENSE) file:

```license
MIT License

Copyright (c) 2020 Adam Chalkley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## References

Various references used when developing this project can be found in our
[references](docs/references.md) doc.

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/mysql2sqlite>  "This project's GitHub repo"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

[go-supported-releases]: <https://go.dev/doc/devel/release#policy> "Go Release Policy"

<!-- []: PLACEHOLDER "DESCRIPTION_HERE" -->
