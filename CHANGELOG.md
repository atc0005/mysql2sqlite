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

## [v0.3.15] - 2024-08-21

### Changed

#### Dependency Updates

- (GH-609) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.21.4 to go-ci-oldstable-alpine-buildx64-v0.21.5 in /dependabot/docker/builds/alpine/x64
- (GH-615) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.21.5 to go-ci-oldstable-alpine-buildx64-v0.21.6 in /dependabot/docker/builds/alpine/x64
- (GH-622) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.21.6 to go-ci-oldstable-alpine-buildx64-v0.21.7 in /dependabot/docker/builds/alpine/x64
- (GH-629) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.21.7 to go-ci-oldstable-alpine-buildx64-v0.21.8 in /dependabot/docker/builds/alpine/x64
- (GH-639) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.21.8 to go-ci-oldstable-alpine-buildx64-v0.21.9 in /dependabot/docker/builds/alpine/x64
- (GH-612) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.21.4 to go-ci-oldstable-alpine-buildx86-v0.21.5 in /dependabot/docker/builds/alpine/x86
- (GH-614) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.21.5 to go-ci-oldstable-alpine-buildx86-v0.21.6 in /dependabot/docker/builds/alpine/x86
- (GH-620) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.21.6 to go-ci-oldstable-alpine-buildx86-v0.21.7 in /dependabot/docker/builds/alpine/x86
- (GH-630) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.21.7 to go-ci-oldstable-alpine-buildx86-v0.21.8 in /dependabot/docker/builds/alpine/x86
- (GH-641) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.21.8 to go-ci-oldstable-alpine-buildx86-v0.21.9 in /dependabot/docker/builds/alpine/x86
- (GH-627) Go Dependency: Bump golang.org/x/sys from 0.22.0 to 0.23.0
- (GH-635) Go Dependency: Bump golang.org/x/sys from 0.23.0 to 0.24.0
- (GH-643) Go Runtime: Bump golang from 1.21.12 to 1.22.6 in /dependabot/docker/go
- (GH-642) Update project to Go 1.22 series

#### Other

- (GH-624) Push `REPO_VERSION` var into containers for builds

### Fixed

- (GH-644) Fix govet linting errors raised by updated linter

## [v0.3.14] - 2024-07-10

### Changed

#### Dependency Updates

- (GH-579) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.20.7 to go-ci-oldstable-alpine-buildx64-v0.20.8 in /dependabot/docker/builds/alpine/x64
- (GH-586) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.20.8 to go-ci-oldstable-alpine-buildx64-0.21.2 in /dependabot/docker/builds/alpine/x64
- (GH-593) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.21.2 to go-ci-oldstable-alpine-buildx64-v0.21.3 in /dependabot/docker/builds/alpine/x64
- (GH-599) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.21.3 to go-ci-oldstable-alpine-buildx64-v0.21.4 in /dependabot/docker/builds/alpine/x64
- (GH-580) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.20.7 to go-ci-oldstable-alpine-buildx86-v0.20.8 in /dependabot/docker/builds/alpine/x86
- (GH-588) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.20.8 to go-ci-oldstable-alpine-buildx86-v0.21.2 in /dependabot/docker/builds/alpine/x86
- (GH-594) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.21.2 to go-ci-oldstable-alpine-buildx86-v0.21.3 in /dependabot/docker/builds/alpine/x86
- (GH-602) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.21.3 to go-ci-oldstable-alpine-buildx86-v0.21.4 in /dependabot/docker/builds/alpine/x86
- (GH-591) Go Dependency: Bump github.com/alexflint/go-arg from 1.5.0 to 1.5.1
- (GH-603) Go Dependency: Bump golang.org/x/sys from 0.21.0 to 0.22.0
- (GH-595) Go Runtime: Bump golang from 1.21.11 to 1.21.12 in /dependabot/docker/go

## [v0.3.13] - 2024-06-07

### Changed

#### Dependency Updates

- (GH-559) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.20.4 to go-ci-oldstable-alpine-buildx64-v0.20.5 in /dependabot/docker/builds/alpine/x64
- (GH-563) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.20.5 to go-ci-oldstable-alpine-buildx64-v0.20.6 in /dependabot/docker/builds/alpine/x64
- (GH-575) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.20.6 to go-ci-oldstable-alpine-buildx64-v0.20.7 in /dependabot/docker/builds/alpine/x64
- (GH-562) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.20.4 to go-ci-oldstable-alpine-buildx86-v0.20.5 in /dependabot/docker/builds/alpine/x86
- (GH-564) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.20.5 to go-ci-oldstable-alpine-buildx86-v0.20.6 in /dependabot/docker/builds/alpine/x86
- (GH-573) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.20.6 to go-ci-oldstable-alpine-buildx86-v0.20.7 in /dependabot/docker/builds/alpine/x86
- (GH-555) Go Dependency: Bump github.com/alexflint/go-arg from 1.4.3 to 1.5.0
- (GH-557) Go Dependency: Bump github.com/fatih/color from 1.16.0 to 1.17.0
- (GH-570) Go Dependency: Bump golang.org/x/sys from 0.20.0 to 0.21.0
- (GH-569) Go Runtime: Bump golang from 1.21.10 to 1.21.11 in /dependabot/docker/go

### Fixed

- (GH-567) Remove inactive maligned linter

## [v0.3.12] - 2024-05-11

### Changed

#### Dependency Updates

- (GH-538) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.20.1 to go-ci-oldstable-alpine-buildx64-v0.20.2 in /dependabot/docker/builds/alpine/x64
- (GH-546) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.20.2 to go-ci-oldstable-alpine-buildx64-v0.20.3 in /dependabot/docker/builds/alpine/x64
- (GH-549) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.20.3 to go-ci-oldstable-alpine-buildx64-v0.20.4 in /dependabot/docker/builds/alpine/x64
- (GH-536) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.20.1 to go-ci-oldstable-alpine-buildx86-v0.20.2 in /dependabot/docker/builds/alpine/x86
- (GH-544) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.20.2 to go-ci-oldstable-alpine-buildx86-v0.20.3 in /dependabot/docker/builds/alpine/x86
- (GH-551) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.20.3 to go-ci-oldstable-alpine-buildx86-v0.20.4 in /dependabot/docker/builds/alpine/x86
- (GH-539) Go Dependency: Bump golang.org/x/sys from 0.19.0 to 0.20.0
- (GH-542) Go Runtime: Bump golang from 1.21.9 to 1.21.10 in /dependabot/docker/go

## [v0.3.11] - 2024-04-11

### Changed

#### Dependency Updates

- (GH-528) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.20.0 to go-ci-oldstable-alpine-buildx64-v0.20.1 in /dependabot/docker/builds/alpine/x64
- (GH-530) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.20.0 to go-ci-oldstable-alpine-buildx86-v0.20.1 in /dependabot/docker/builds/alpine/x86
- (GH-531) Go Dependency: Bump golang.org/x/sys from 0.18.0 to 0.19.0
- (GH-524) Go Runtime: Bump golang from 1.21.8 to 1.21.9 in /dependabot/docker/go

## [v0.3.10] - 2024-04-01

### Changed

#### Dependency Updates

- (GH-517) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.19.0 to go-ci-oldstable-alpine-buildx86-v0.20.0 in /dependabot/docker/builds/alpine/x86
- (GH-514) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.19.0 to go-ci-oldstable-alpine-buildx64-v0.20.0 in /dependabot/docker/builds/alpine/x64

### Fixed

- (GH-520) Generate missing asset links

## [v0.3.9] - 2024-03-27

### Changed

#### Dependency Updates

- (GH-495) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.15.4 to go-ci-oldstable-alpine-buildx64-v0.16.0 in /dependabot/docker/builds/alpine/x64
- (GH-500) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.16.0 to go-ci-oldstable-alpine-buildx64-v0.16.1 in /dependabot/docker/builds/alpine/x64
- (GH-503) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.16.1 to go-ci-oldstable-alpine-buildx64-v0.19.0 in /dependabot/docker/builds/alpine/x64
- (GH-496) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.15.4 to go-ci-oldstable-alpine-buildx86-v0.16.0 in /dependabot/docker/builds/alpine/x86
- (GH-499) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.16.0 to go-ci-oldstable-alpine-buildx86-v0.16.1 in /dependabot/docker/builds/alpine/x86
- (GH-504) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.16.1 to go-ci-oldstable-alpine-buildx86-v0.19.0 in /dependabot/docker/builds/alpine/x86
- (GH-492) Go Dependency: Bump github.com/go-sql-driver/mysql from 1.7.1 to 1.8.0
- (GH-507) Go Dependency: Bump github.com/go-sql-driver/mysql from 1.8.0 to 1.8.1

### Fixed

- (GH-510) Fix build version insertion

## [v0.3.8] - 2024-03-08

### Changed

#### Dependency Updates

- (GH-488) Add todo/release label to "Go Runtime" PRs
- (GH-479) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.15.2 to go-ci-oldstable-alpine-buildx64-v0.15.3 in /dependabot/docker/builds/alpine/x64
- (GH-487) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.15.3 to go-ci-oldstable-alpine-buildx64-v0.15.4 in /dependabot/docker/builds/alpine/x64
- (GH-477) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.15.2 to go-ci-oldstable-alpine-buildx86-v0.15.3 in /dependabot/docker/builds/alpine/x86
- (GH-486) Build Image: Bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.15.3 to go-ci-oldstable-alpine-buildx86-v0.15.4 in /dependabot/docker/builds/alpine/x86
- (GH-473) canary: bump golang from 1.21.6 to 1.21.7 in /dependabot/docker/go
- (GH-468) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.15.0 to go-ci-oldstable-alpine-buildx64-v0.15.2 in /dependabot/docker/builds/alpine/x64
- (GH-467) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.15.0 to go-ci-oldstable-alpine-buildx86-v0.15.2 in /dependabot/docker/builds/alpine/x86
- (GH-481) Go Dependency: Bump golang.org/x/sys from 0.17.0 to 0.18.0
- (GH-482) Go Runtime: Bump golang from 1.21.7 to 1.21.8 in /dependabot/docker/go
- (GH-475) Update Dependabot PR prefixes (redux)
- (GH-474) Update Dependabot PR prefixes
- (GH-472) Update project to Go 1.21 series

## [v0.3.7] - 2024-02-19

### Changed

#### Dependency Updates

- (GH-454) canary: bump golang from 1.20.13 to 1.20.14 in /dependabot/docker/go
- (GH-421) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.14.3 to go-ci-oldstable-alpine-buildx64-v0.14.4 in /dependabot/docker/builds/alpine/x64
- (GH-429) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.14.4 to go-ci-oldstable-alpine-buildx64-v0.14.5 in /dependabot/docker/builds/alpine/x64
- (GH-440) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.14.5 to go-ci-oldstable-alpine-buildx64-v0.14.6 in /dependabot/docker/builds/alpine/x64
- (GH-459) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.14.6 to go-ci-oldstable-alpine-buildx64-v0.14.9 in /dependabot/docker/builds/alpine/x64
- (GH-464) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.14.9 to go-ci-oldstable-alpine-buildx64-v0.15.0 in /dependabot/docker/builds/alpine/x64
- (GH-420) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.14.3 to go-ci-oldstable-alpine-buildx86-v0.14.4 in /dependabot/docker/builds/alpine/x86
- (GH-430) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.14.4 to go-ci-oldstable-alpine-buildx86-v0.14.5 in /dependabot/docker/builds/alpine/x86
- (GH-439) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.14.5 to go-ci-oldstable-alpine-buildx86-v0.14.6 in /dependabot/docker/builds/alpine/x86
- (GH-456) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.14.6 to go-ci-oldstable-alpine-buildx86-v0.14.9 in /dependabot/docker/builds/alpine/x86
- (GH-461) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.14.9 to go-ci-oldstable-alpine-buildx86-v0.15.0 in /dependabot/docker/builds/alpine/x86
- (GH-432) go.mod: bump github.com/atc0005/go-nagios from 0.16.0 to 0.16.1
- (GH-431) go.mod: bump github.com/mattn/go-sqlite3 from 1.14.19 to 1.14.20
- (GH-433) go.mod: bump github.com/mattn/go-sqlite3 from 1.14.20 to 1.14.21
- (GH-442) go.mod: bump github.com/mattn/go-sqlite3 from 1.14.21 to 1.14.22
- (GH-452) go.mod: bump golang.org/x/sys from 0.16.0 to 0.17.0

### Fixed

- (GH-435) Fix configuration settings table syntax

## [v0.3.6] - 2024-01-19

### Changed

#### Dependency Updates

- (GH-412) canary: bump golang from 1.20.12 to 1.20.13 in /dependabot/docker/go
- (GH-414) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.14.2 to go-ci-oldstable-alpine-buildx64-v0.14.3 in /dependabot/docker/builds/alpine/x64
- (GH-417) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.14.2 to go-ci-oldstable-alpine-buildx86-v0.14.3 in /dependabot/docker/builds/alpine/x86
- (GH-407) ghaw: bump github/codeql-action from 2 to 3
- (GH-408) go.mod: bump github.com/mattn/go-sqlite3 from 1.14.18 to 1.14.19
- (GH-411) go.mod: bump golang.org/x/sys from 0.15.0 to 0.16.0

## [v0.3.5] - 2023-12-09

### Changed

#### Dependency Updates

- (GH-398) canary: bump golang from 1.20.11 to 1.20.12 in /dependabot/docker/go
- (GH-402) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.14.1 to go-ci-oldstable-alpine-buildx64-v0.14.2 in /dependabot/docker/builds/alpine/x64
- (GH-400) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.14.1 to go-ci-oldstable-alpine-buildx86-v0.14.2 in /dependabot/docker/builds/alpine/x86
- (GH-396) go.mod: bump golang.org/x/sys from 0.14.0 to 0.15.0

## [v0.3.4] - 2023-11-16

### Changed

#### Dependency Updates

- (GH-385) canary: bump golang from 1.20.10 to 1.20.11 in /dependabot/docker/go
- (GH-339) canary: bump golang from 1.20.7 to 1.20.8 in /dependabot/docker/go
- (GH-364) canary: bump golang from 1.20.8 to 1.20.10 in /dependabot/docker/go
- (GH-389) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.13.12 to go-ci-oldstable-alpine-buildx64-v0.14.1 in /dependabot/docker/builds/alpine/x64
- (GH-323) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.13.4 to go-ci-oldstable-alpine-buildx64-v0.13.5 in /dependabot/docker/builds/alpine/x64
- (GH-327) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.13.5 to go-ci-oldstable-alpine-buildx64-v0.13.6 in /dependabot/docker/builds/alpine/x64
- (GH-331) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.13.6 to go-ci-oldstable-alpine-buildx64-v0.13.7 in /dependabot/docker/builds/alpine/x64
- (GH-342) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.13.7 to go-ci-oldstable-alpine-buildx64-v0.13.8 in /dependabot/docker/builds/alpine/x64
- (GH-352) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.13.8 to go-ci-oldstable-alpine-buildx64-v0.13.9 in /dependabot/docker/builds/alpine/x64
- (GH-371) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx64-v0.13.9 to go-ci-oldstable-alpine-buildx64-v0.13.12 in /dependabot/docker/builds/alpine/x64
- (GH-378) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.13.12 to go-ci-oldstable-alpine-buildx86-v0.14.0 in /dependabot/docker/builds/alpine/x86
- (GH-325) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.13.4 to go-ci-oldstable-alpine-buildx86-v0.13.5 in /dependabot/docker/builds/alpine/x86
- (GH-328) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.13.5 to go-ci-oldstable-alpine-buildx86-v0.13.6 in /dependabot/docker/builds/alpine/x86
- (GH-333) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.13.6 to go-ci-oldstable-alpine-buildx86-v0.13.7 in /dependabot/docker/builds/alpine/x86
- (GH-340) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.13.7 to go-ci-oldstable-alpine-buildx86-v0.13.8 in /dependabot/docker/builds/alpine/x86
- (GH-350) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.13.8 to go-ci-oldstable-alpine-buildx86-v0.13.9 in /dependabot/docker/builds/alpine/x86
- (GH-369) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.13.9 to go-ci-oldstable-alpine-buildx86-v0.13.12 in /dependabot/docker/builds/alpine/x86
- (GH-390) docker: bump atc0005/go-ci from go-ci-oldstable-alpine-buildx86-v0.14.0 to go-ci-oldstable-alpine-buildx86-v0.14.1 in /dependabot/docker/builds/alpine/x86
- (GH-337) ghaw: bump actions/checkout from 3 to 4
- (GH-382) go.mod: bump github.com/fatih/color from 1.15.0 to 1.16.0
- (GH-373) go.mod: bump github.com/mattn/go-isatty from 0.0.19 to 0.0.20
- (GH-381) go.mod: bump github.com/mattn/go-sqlite3 from 1.14.17 to 1.14.18
- (GH-335) go.mod: bump golang.org/x/sys from 0.11.0 to 0.12.0
- (GH-359) go.mod: bump golang.org/x/sys from 0.12.0 to 0.13.0
- (GH-383) go.mod: bump golang.org/x/sys from 0.13.0 to 0.14.0

## [v0.3.3] - 2023-08-17

### Added

- (GH-282) Add initial automated release notes config
- (GH-285) Add initial automated release build workflow

### Changed

- Dependencies
  - `Go`
    - `1.20.6` to `1.20.7`
  - `atc0005/go-ci`
    - `go-ci-stable-alpine-buildx64-v0.11.4` to
      `go-ci-oldstable-alpine-buildx64-v0.13.4`
    - `go-ci-stable-alpine-buildx86-v0.11.4` to
      `go-ci-oldstable-alpine-buildx86-v0.13.4`
  - `golang.org/x/sys`
    - `v0.10.0` to `v0.11.0`
- (GH-286) Update Dependabot config to monitor both branches
- (GH-317) Update project to use oldstable build images

### Fixed

- (GH-288) Add missing key to Dependabot configuration

## [v0.3.2] - 2023-07-14

### Overview

- RPM package improvements
- Bug fixes
- Dependency updates
- built using Go 1.20.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.20.5` to `1.20.6`
  - `atc0005/go-ci`
    - `go-ci-stable-alpine-buildx64-v0.11.0` to
      `go-ci-stable-alpine-buildx64-v0.11.4`
    - `go-ci-stable-alpine-buildx86-v0.11.0` to
      `go-ci-stable-alpine-buildx86-v0.11.4`
  - `atc0005/go-nagios`
    - `v0.15.0` to `v0.16.0`
  - `golang.org/x/sys`
    - `v0.9.0` to `v0.10.0`
- (GH-276) Update RPM postinstall scripts to use restorecon

### Fixed

- (GH-270) Misc CHANGELOG fixes

## [v0.3.1] - 2023-06-16

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- built using Go 1.20.5
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.20.4` to `1.20.5`
  - `atc0005/go-ci`
    - `go-ci-stable-alpine-buildx64-v0.10.5` to
      `go-ci-stable-alpine-buildx64-v0.11.0`
    - `go-ci-stable-alpine-buildx86-v0.10.5` to
      `go-ci-stable-alpine-buildx86-v0.11.0`
  - `atc0005/go-nagios`
    - `v0.14.0` to `v0.15.0`
  - `mattn/go-sqlite3`
    - `v1.14.16` to `v1.14.17`
  - `mattn/go-isatty`
    - `v0.0.18` to `v0.0.19`
  - `golang.org/x/sys`
    - `v0.8.0` to `v0.9.0`
  - `actions/checkout`
    - `v3.5.2` to `v3.5.3`
- (GH-263) Update vuln analysis GHAW to remove on.push hook

### Fixed

- (GH-259) Disable depguard linter
- (GH-266) Restore local CodeQL workflow

## [v0.3.0] - 2023-05-26

### Overview

- Add support for generating DEB, RPM packages
- Build improvements
- Generated binary changes
  - filename patterns
  - compression (~ 66% smaller)
  - executable metadata
- Bug fixes
- built using Go 1.20.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-233) Generate RPM/DEB packages using nFPM
- (GH-236) Add version details to Windows executables

### Changed

- Dependencies
  - `mattn/go-isatty`
    - `v0.0.18` to `v0.0.19`
- (GH-237) Switch to semantic versioning (semver) compatible versioning
  pattern
- (GH-238) Makefile: Compress binaries & use fixed filenames
- (GH-235) Makefile: Refresh recipes to add "standard" set, new
  package-related options
- (GH-234) Build dev/stable releases using go-ci Docker image
- (GH-251) Move `Build using Makefile container recipes` jobs to monthly
  schedule instead of for every PR

### Fixed

- (GH-247) Update project Go version to use stable release
- (GH-248) Fix CHANGELOG: note correct Go version for v0.2.8

## [v0.2.8] - 2023-05-17

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- built using Go 1.20.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.20.4`
    - **Statically linked**
    - created via `docker` Makefile recipe
    - linked to musl libc
      - Alpine package version `1.2.3-r2`
        - bundled within `atc0005/go-ci` [project
          containers](https://github.com/atc0005/go-ci/pkgs/container/go-ci)
          - `atc0005/go-ci:go-ci-stable-alpine-buildx86-v0.10.5`
          - `atc0005/go-ci:go-ci-stable-alpine-buildx64-v0.10.5`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `go-sql-driver/mysql`
    - `v1.7.0` to `v1.7.1`
  - `atc0005/go-nagios`
    - `v0.10.2` to `v0.14.0`
  - `fatih/color`
    - `v1.13.0` to `v1.15.0`
  - `mattn/go-isatty`
    - `v0.0.16` to `v0.0.18`
  - `go-logfmt/logfmt`
    - `v0.5.1` to `v0.6.0`
  - `golang.org/x/sys`
    - `v0.3.0` to `v0.8.0`
- (GH-205) Update nagios library usage, add time perfdata
- (GH-210) Add Go Module Validation, Dependency Updates jobs
- (GH-217) Update `Build using Makefile Docker recipes` CI workflow jobs to
  use `go-ci-stable-debian-build` image
- (GH-220) Drop `Push Validation` workflow
- (GH-221) Rework workflow scheduling
- (GH-223) Remove `Push Validation` workflow status badge

### Fixed

- (GH-206) Fix doc comment typo
- (GH-218) Explicitly mark CWD as trusted by Git
- (GH-224) Add missing os deps list to monthly workflow call
- (GH-229) Update vuln analysis GHAW to use on.push hook
- (GH-243) Use UNKNOWN state for invalid command-line args

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
  - `go-sql-driver/mysql`
    - `v1.6.0` to `v1.7.0`
  - `mattn/go-sqlite3`
    - `v1.14.15` to `v1.14.16`
  - `atc0005/go-nagios`
    - `v0.10.0` to `v0.10.2`
  - `mattn/go-colorable`
    - `v0.1.2` to `v0.1.13`
  - `mattn/go-isatty`
    - `v0.0.8` to `v0.0.16`
  - `golang.org/x/sys`
    - `v0.0.0-20190412213103-97732733099d` to `v0.3.0`
  - `alexflint/go-scalar`
    - `v1.1.0` to `v1.2.0`
  - `fatih/color`
    - `v1.7.0` to `v1.13.0`
  - `go-logfmt/logfmt`
    - `v0.4.0` to `v0.5.1`
  - `pkg/errors`
    - `v0.8.1` to `v0.9.1`
  - `kr/logfmt`
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

[Unreleased]: https://github.com/atc0005/mysql2sqlite/compare/v0.3.15...HEAD
[v0.3.15]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.15
[v0.3.14]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.14
[v0.3.13]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.13
[v0.3.12]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.12
[v0.3.11]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.11
[v0.3.10]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.10
[v0.3.9]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.9
[v0.3.8]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.8
[v0.3.7]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.7
[v0.3.6]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.6
[v0.3.5]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.5
[v0.3.4]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.4
[v0.3.3]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.3
[v0.3.2]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.2
[v0.3.1]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.1
[v0.3.0]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.3.0
[v0.2.8]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.2.8
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
