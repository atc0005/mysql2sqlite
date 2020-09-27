# mysql2sqlite

Mirror MySQL database tables to SQLite

## Overview

## Instructions

### Setup development environment

1. Create a new LXD container or VM
   - as of this writing Ubuntu 16.04 LTS is the quickest setup path due to the
     assumptions made by the `mysql2sqlite-dev` project
1. Clone the `mysql2sqlite-dev` project within this new environment
1. Run bash scripts
   1. `bash bin/setup_dev_environment.sh`
   1. `bash bin/create_db.sh`
   1. `python3 /tmp/mysql2sqlite-dev/bin/validate_dbs.py`
1. Setup Go development environment
   1. [Download](https://golang.org/dl/) Go
   1. [Install](https://golang.org/doc/install) Go
      - NOTE: Pay special attention to the remarks about `$HOME/.profile`
1. Clone *this* repo

### Build

TODO: Flesh this out with Makefile example, etc.

## References

- <https://github.com/WhyAskWhy/mysql2sqlite-dev>
- <https://github.com/WhyAskWhy/mysql2sqlite>

- <https://github.com/go-sql-driver/mysql#usage>
- <https://github.com/go-sql-driver/mysql#timetime-support>
- <https://github.com/joho/sqltocsv>
  - <https://github.com/joho/sqltocsv/blob/5650f27fd5b64a3806f6251308cba1a4d5ec598d/sqltocsv.go#L123-L159>

- <https://www.sqlite.org/datatype3.html>
- <https://www.sqlite.org/lang_createindex.html>

- Get row count for table
  - <https://gist.github.com/trkrameshkumar/f4f1c00ef5d578561c96>
  - <https://gist.github.com/agis/7e8cd4c7a20d037c01e663402fcad9d0>
