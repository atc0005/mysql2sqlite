<!-- omit in toc -->
# mysql2sqlite: References

- [Project README](../README.md)

<!-- omit in toc -->
## Table of contents

- [Overview](#overview)
- [References](#references)
  - [External](#external)
  - [Dependencies](#dependencies)
  - [Instruction / Examples](#instruction--examples)
  - [Related projects](#related-projects)

## Overview

The links below are for resources that were found to be useful (if not
absolutely essential) while developing this application.

## References

### External

- Nagios
  - <https://nagios-plugins.org/doc/guidelines.html>

- Cron
  - <http://manpages.ubuntu.com/manpages/man5/crontab.5.html>

- Postfix
  - <http://www.postfix.org/DATABASE_README.html>
  - <http://www.postfix.org/sqlite_table.5.html>
  - <https://workaround.org/ispmail>

- logrotate
  - <https://github.com/logrotate/logrotate>
  - <http://manpages.ubuntu.com/manpages/man8/logrotate.8.html>

### Dependencies

- `make` on Windows
  - <https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows>
- `gcc` on Windows
  - <https://en.wikipedia.org/wiki/MinGW>
  - <http://mingw-w64.org/>
  - <https://www.msys2.org/>

- Docker
  - <https://hub.docker.com/_/golang>
    - provides base images that the `atc0005/go-ci` images build upon

- Libraries/packages
  - Configuration
    - <https://github.com/alexflint/go-arg>
    - <https://github.com/go-yaml/yaml>
  - Logging
    - <https://github.com/apex/log>
  - Go database drivers (see also `go.mod`)
    - <https://github.com/go-sql-driver/mysql>
    - <https://github.com/mattn/go-sqlite3>

- cgo
  - <https://golang.org/cmd/cgo/>
  - <https://blog.golang.org/cgo>
  - <https://github.com/golang/go/wiki/cgo>
  - <https://dave.cheney.net/2016/01/18/cgo-is-not-go>

- Go / Static compilation
  - <https://github.com/golang/go/issues/26492>
  - <https://github.com/golang/go/issues/18773#issuecomment-274975226>

### Instruction / Examples

- <https://github.com/WhyAskWhy/mysql2sqlite-dev>
- <https://github.com/WhyAskWhy/mysql2sqlite>
- <https://github.com/joho/sqltocsv>
  - <https://github.com/joho/sqltocsv/blob/5650f27fd5b64a3806f6251308cba1a4d5ec598d/sqltocsv.go#L123-L159>

- SQLite
  - <https://www.sqlite.org/datatype3.html>
  - <https://www.sqlite.org/lang_createindex.html>
  - <https://dba.stackexchange.com/questions/45368/how-do-i-prevent-sqlite-database-locks>
  - <https://manski.net/2012/10/sqlite-performance/>
  - <https://www.sqlite.org/wal.html>
  - <https://www.sqlite.org/pragma.html#pragma_journal_mode>
  - <https://www.sqlite.org/pragma.html#pragma_busy_timeout>
  - <https://github.com/mattn/go-sqlite3/pull/827>
  - <https://github.com/mattn/go-sqlite3/issues/209>

- Get row count for table
  - <https://gist.github.com/trkrameshkumar/f4f1c00ef5d578561c96>
  - <https://gist.github.com/agis/7e8cd4c7a20d037c01e663402fcad9d0>

- Nagios
  - <https://github.com/atc0005/go-nagios>
  - <https://nagios-plugins.org/doc/guidelines.html>

### Related projects

- <https://github.com/atc0005/go-nagios>
- <https://github.com/atc0005/go-ci>
  - provides Docker images used for musl libc static builds via `Makefile`
