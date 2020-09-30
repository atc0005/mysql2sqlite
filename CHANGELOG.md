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

[Unreleased]: https://github.com/atc0005/mysql2sqlite/compare/v0.1.0...HEAD
[v0.1.0]: https://github.com/atc0005/mysql2sqlite/releases/tag/v0.1.0
