# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/mysql2sqlite/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.1.3] - 2021-01-21

### Changed

- Builds
  - Force latest container image to be used for builds

- Dependencies
  - built using Go 1.15.7
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.1.24-r10`
        - bundled within `atc0005/go-ci` project containers
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.3.7-0-g6d59440`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.3.7-0-g6d59440`
    - Windows
      - x86
      - x64
    - Linux
      - x86
      - x64
  - `mattn/go-sqlite3`
    - `1.14.5` to `1.14.6`
  - `atc0005/go-nagios`
    - `v0.5.2` to `v0.6.0`
  - `actions/setup-node`
    - `v2.1.2` to `v2.1.4`
  - `gopkg.in/yaml.v2`
    - `v2.3.0` to `v2.4.0`

### Fixed

- Invalid `err` variable reference
- Inconsistent use of `err` variable, custom named error variables

## [v0.1.2] - 2020-11-17

### Changed

- Dependencies
  - built using Go 1.15.5
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.1.24-r9`
        - bundled within `atc0005/go-ci` project containers
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.2.14-0-ge9c045f`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.2.14-0-ge9c045f`
    - Windows
      - x86
      - x64
    - Linux
      - x86
      - x64
  - `mattn/go-sqlite3`
    - `1.14.4` to `1.14.5`
  - `atc0005/go-nagios`
    - `v0.5.1` to `v0.5.2`
  - `actions/checkout`
    - `v2.3.3` to `v2.3.4`

### Fixed

- Typo in prior CHANGELOG entry

## [v0.1.1] - 2020-10-09

### Added

- New Makefile recipes
  - `all-static`
  - `linux`
  - `linux-static`
  - `linux-x86`
  - `linux-x86-static`
  - `linux-x64`
  - `linux-x64-static`
  - `windows`
  - `windows-x86`
  - `windows-x86-static`
  - `windows-x64`
  - `windows-x64-static`
  - `docker`
  - NOTE: If `static` suffix isn't present, the recipe produces dynamic
    executables

- Static builds
  - created via `docker` Makefile recipe
  - available via various other `-static` suffixed Makefile recipes
  - linked to musl libc
    - Alpine package version `1.1.24-r9`
      - bundled within `atc0005/go-ci` project containers
        - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.2.7-0-gcbb8139`
        - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.2.7-0-gcbb8139`
      - provided indirectly by `golang:1.15.2-alpine3.12` and
        `i386/golang:1.15.2-alpine3.12` base containers

### Changed

- GitHub Actions Workflows
  - `Lint and Build using Makefile`
    - `Build dynamically linked binaries using Makefile`
    - `Build statically linked binaries using Makefile`
      - glibc based, tossed after build
    - `Build static binaries using Docker images`
      - musl libc based, tossed after build
  - `Validate Codebase`
    - `Build codebase`
      - explicit `CGO_ENABLED=1` environment variable set for build commands
- Add `-trimpath` build flag
- Dependencies
  - `mattn/go-sqlite3`
    - `1.14.3` to `1.14.4`
  - `actions/setup-node`
    - `v2.1.1` to `v2.1.2`

### Fixed

- Windows builds are broken
  - cgo was not enabled
- Fix insert statement for MySQL test data
  - worked with MariaDB 10.0.x, failed with MariaDB 10.1.x
- Go version provided by GitHub Actions virtual environment printed
  - this inferred that the version of Go used for the `Build static binaries
    using Docker images` job was older than it actually is

## [v0.1.0] - 2020-09-30

Initial release!

This release provides an early release version of two applications used to
mirror and validate source MySQL database tables to a local SQLite database.

### Added

- CLI tool for mirroring MySQL database tables to SQLite database
- Nagios plugin for validating mirrored SQLite database against the original
  source
- Configurable source database, destination database settings
- Configurable connection retry, retry delay behavior
- Configurable logging settings
  - level
  - output "target" (`stdout`, `stderr`)
  - format

- Binary release
  - Built using Go 1.15.2
  - Windows
    - x86
    - x64
  - Linux
    - x86
    - x64

[Unreleased]: https://github.com/atc0005/mysql2sqlite/compare/v0.1.3...HEAD
[v0.1.3]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.3
[v0.1.2]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.0
