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

## [v0.2.7] - 2022-12-09

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.1` to `1.19.4`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.3-r2`
        - bundled within `atc0005/go-ci` [project
          containers](https://github.com/atc0005/go-ci/pkgs/container/go-ci)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.7.2-0-g433353d`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.7.2-0-g433353d`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `github.com/go-sql-driver/mysql`
    - `v1.6.0` to `v1.7.0`
  - `github.com/mattn/go-sqlite3`
    - `v1.14.15` to `v1.14.16`
  - `atc0005/go-nagios`
    - `v0.10.0` to `v0.10.2`
  - `github.com/mattn/go-colorable`
    - `v0.1.2` to `v0.1.13`
  - `github.com/mattn/go-isatty`
    - `v0.0.8` to `v0.0.16`
  - `golang.org/x/sys`
    - `v0.0.0-20190412213103-97732733099d` to `v0.3.0`
  - `github.com/alexflint/go-scalar`
    - `v1.1.0` to `v1.2.0`
  - `github.com/fatih/color`
    - `v1.7.0` to `v1.13.0`
  - `github.com/go-logfmt/logfmt`
    - `v0.4.0` to `v0.5.1`
  - `github.com/pkg/errors`
    - `v0.8.1` to `v0.9.1`
  - `github.com/kr/logfmt`
    - `v0.0.0-20140226030751-b84e30acd515` to
      `v0.0.0-20210122060352-19f9bcb100e6`
- (GH-189) Refactor GitHub Actions workflows to import logic

### Fixed

- (GH-192) Broken links in README
- (GH-198) Fix Makefile Go module base path detection

## [v0.2.6] - 2022-09-22

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.1
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.13` to `1.19.1`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.3-r0`
        - bundled within `atc0005/go-ci` [project
          containers](https://github.com/atc0005/go-ci/pkgs/container/go-ci)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.7.0-0-gf9532bf`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.7.0-0-gf9532bf`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `atc0005/go-nagios`
    - `v0.9.1` to `v0.10.0`
  - `github/codeql-action`
    - `v2.1.22` to `v2.1.25`
- (GH-180) Update project to Go 1.19
- (GH-181) Update Makefile and GitHub Actions Workflows

### Fixed

- (GH-179) Add missing cmd doc files

## [v0.2.5] - 2022-08-23

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.17.13
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.12` to `1.17.13`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.3-r0`
        - bundled within `atc0005/go-ci` [project
          containers](https://github.com/atc0005/go-ci/pkgs/container/go-ci)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.6.22-0-gd3932e8`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.6.22-0-gd3932e8`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `mattn/go-sqlite3`
    - `v1.14.10` to `v1.14.15`

### Fixed

- (GH-176) Apply Go 1.19 specific doc comments linting fixes
- (GH-177) Swap io/ioutil package for io package

## [v0.2.4] - 2022-07-21

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.17.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.10` to `1.17.12`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.3-r0`
        - bundled within `atc0005/go-ci` [project
          containers](https://github.com/atc0005/go-ci/pkgs/container/go-ci)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.6.19-0-g3ff32f5`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.6.19-0-g3ff32f5`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `atc0005/go-nagios`
    - `v0.8.2` to `v0.9.1`

### Fixed

- (GH-169) Fix various atc0005/go-nagios usage linting errors
- (GH-170) Update lintinstall Makefile recipe

## [v0.2.3] - 2022-05-13

### Overview

- Dependency updates
- built using Go 1.17.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.9` to `1.17.10`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.2-r3`
        - bundled within `atc0005/go-ci` [project
          containers](https://github.com/atc0005/go-ci/pkgs/container/go-ci)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.6.11-0-g59b07ad`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.6.11-0-g59b07ad`
    - Windows (x86, x64)
    - Linux (x86, x64)

## [v0.2.2] - 2022-05-06

### Overview

- Dependency updates
- built using Go 1.17.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.7` to `1.17.9`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.2-r3`
        - bundled within `atc0005/go-ci` [project
          containers](https://github.com/atc0005/go-ci/pkgs/container/go-ci)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.6.8-0-g607425a`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.6.8-0-g607425a`
    - Windows (x86, x64)
    - Linux (x86, x64)

### Fixed

- (GH-159) fatal: unsafe repository ('/path/to/repo' is owned by someone else)
- (GH-161) WARNING: The requested image's platform (linux/386) does not match
  the detected host platform (linux/amd64) and no specific platform was
  requested

## [v0.2.1] - 2022-03-02

### Overview

- Dependency updates
- CI / linting improvements
- built using Go 1.17.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.6` to `1.17.7`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.2-r3`
        - bundled within `atc0005/go-ci` [project
          containers](https://github.com/atc0005/go-ci/pkgs/container/go-ci)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.6.2-0-g90db923`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.6.2-0-g90db923`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `alexflint/go-arg`
    - `v1.4.2` to `v1.4.3`
  - `actions/checkout`
    - `v2.4.0` to `v3`
  - `actions/setup-node`
    - `v2.5.1` to `v3`

- (GH-145) Expand linting GitHub Actions Workflow to include `oldstable`,
  `unstable` container images
- (GH-146) Switch Docker image source from Docker Hub to GitHub Container
  Registry (GHCR)

### Fixed

- (GH-144) Fix year in changelog entry
- (GH-148) gosec, revive linting errors

## [v0.2.0] - 2022-01-25

### Overview

- Added additional MySQL, SQLite settings
- Dependency updates
- built using Go 1.17.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-10) Add config file option for setting SQLite journal mode
- (GH-35) Add config file setting for specifying SQLite busy_timeout
- (GH-37) Re-enable support for setting max idle connection time when Go 1.16
  releases

### Changed

- Dependencies
  - `Go`
    - `1.16.12` to `1.17.6`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.2-r3`
        - bundled within `atc0005/go-ci` [project
          containers](https://hub.docker.com/r/atc0005/go-ci/tags?page=1&ordering=last_updated&name=alpine)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.3.42-0-g44e0ca5`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.3.42-0-g44e0ca5`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `atc0005/go-nagios`
    - `v0.8.1` to `v0.8.2`
  - `mattn/go-sqlite3`
    - `v1.14.9` to `v1.14.10`

## [v0.1.9] - 2021-12-29

### Overview

- Dependency updates
- built using Go 1.16.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.10` to `1.16.12`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.2-r3`
        - bundled within `atc0005/go-ci` [project
          containers](https://hub.docker.com/r/atc0005/go-ci/tags?page=1&ordering=last_updated&name=alpine)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.3.40-0-g0ccb379`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.3.40-0-g0ccb379`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `actions/setup-node`
    - `v2.4.1` to `v2.5.1`

## [v0.1.8] - 2021-11-09

### Overview

- Dependency updates
- built using Go 1.16.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.8` to `1.16.10`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.2-r3`
        - bundled within `atc0005/go-ci` [project
          containers](https://hub.docker.com/r/atc0005/go-ci/tags?page=1&ordering=last_updated&name=alpine)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.3.37-0-gca2d5e9`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.3.37-0-gca2d5e9`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `mattn/go-sqlite3`
    - `v1.14.8` to `v1.14.9`
  - `atc0005/go-nagios`
    - `v0.7.0` to `v0.8.1`
  - `actions/checkout`
    - `v2.3.4` to `v2.4.0`
  - `actions/setup-node`
    - `v2.4.0` to `v2.4.1`

### Fixed

- (GH-119) README | Add missing footnote for Go install guide
- (GH-123) False positive `G307: Deferring unsafe method "Close" on type
  "*os.File" (gosec)` linting error

## [v0.1.7] - 2021-09-25

### Overview

- Dependency updates
- built using Go 1.16.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.7` to `1.16.8`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.2-r3`
        - bundled within `atc0005/go-ci` [project
          containers](https://hub.docker.com/r/atc0005/go-ci/tags?page=1&ordering=last_updated&name=alpine)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.3.33-0-gc057814`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.3.33-0-gc057814`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `atc0005/go-nagios`
    - `v0.6.1` to `v0.7.0`

## [v0.1.6] - 2021-08-08

### Overview

- Dependency updates
- built using Go 1.16.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.6` to `1.16.7`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.2-r3`
        - bundled within `atc0005/go-ci` project containers
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.3.30-0-g09f6ec7`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.3.30-0-g09f6ec7`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `actions/setup-node`
    - updated from `v2.2.0` to `v2.4.0`

## [v0.1.5] - 2021-07-18

### Overview

- Add tests
- Bug fixes
- Dependency updates
- built using Go 1.16.6

### Added

- Add "canary" Dockerfile to track stable Go releases, serve as a reminder to
  generate fresh binaries
- Add tests to assert presence of custom flags

### Changed

- Dependencies
  - built using Go 1.16.6
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.2-r3`
        - bundled within `atc0005/go-ci` project containers
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.3.28-0-g665e0b9`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.3.28-0-g665e0b9`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `mattn/go-sqlite3`
    - `1.14.7` to `1.14.8`
  - `alexflint/go-arg`
    - `v1.3.0` to `v1.4.2`
  - `atc0005/go-nagios`
    - `v0.6.0` to `v0.6.1`
  - `actions/setup-node`
    - `v2.1.5` to `v2.2.0`
    - update `node-version` value to always use latest LTS version instead of
      hard-coded version

### Fixed

- Build docs have broken reference-style links

## [v0.1.4] - 2021-04-15

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.16.3

### Changed

- Replace godoc.org badge with pkg.go.dev badge

- Dependencies
  - built using Go 1.16.3
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.1.24-r10`
        - bundled within `atc0005/go-ci` project containers
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.3.20-0-ge80ea25`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.3.20-0-ge80ea25`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `go-sql-driver/mysql`
    - `v1.5.0` to `v1.6.0`
  - `mattn/go-sqlite3`
    - `1.14.6` to `1.14.7`
  - `actions/setup-node`
    - `v2.1.4` to `v2.1.5`

### Fixed

- Add doc coverage for manual first run requirement
- Increase golangci-lint timeout from 1m default
- Add missing log.Errorf() parameter

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

[Unreleased]: https://github.com/atc0005/mysql2sqlite/compare/v0.2.7...HEAD
[v0.2.7]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.2.7
[v0.2.6]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.2.6
[v0.2.5]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.2.5
[v0.2.4]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.2.4
[v0.2.3]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.2.3
[v0.2.2]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.2.2
[v0.2.1]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.2.0
[v0.1.9]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.9
[v0.1.8]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.8
[v0.1.7]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.7
[v0.1.6]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.6
[v0.1.5]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.5
[v0.1.4]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.4
[v0.1.3]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.3
[v0.1.2]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.0
