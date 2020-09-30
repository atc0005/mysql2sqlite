<!-- omit in toc -->
# mysql2sqlite: Building

- [Project README](../README.md)

<!-- omit in toc -->
## Table of contents

- [Build source code](#build-source-code)
- [Optional: Use existing binaries](#optional-use-existing-binaries)

## Build source code

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
   - NOTE: Pay special attention to the remarks about `$HOME/.profile`
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/mysql2sqlite`
   1. `cd mysql2sqlite`
1. Install dependencies
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     - `sudo yum install make gcc`
   - for Windows
     - Emulated environments (*easier*)
       - Skip all of this and build using the default `go build` command in
         Windows (see below for use of the `-mod=vendor` flag)
       - build using Windows Subsystem for Linux Ubuntu environment and just
         copy out the Windows binaries from that environment
       - If already running a Docker environment, use a container with the Go
         tool-chain already installed
       - If already familiar with LXD, create a container and follow the
         installation steps given previously to install required dependencies
     - Native tooling (*harder*)
       - see the StackOverflow Question `32127524` link in the
         [References](references.md) section for potential options for
         installing `make` on Windows
       - see the mingw-w64 project homepage link in the
         [References](references.md) section for options for installing `gcc`
         and related packages on Windows
1. Build binaries
   - for the current operating system, explicitly using bundled dependencies
         in top-level `vendor` folder
     - `go build -mod=vendor ./cmd/check_mysql2sqlite/`
     - `go build -mod=vendor ./cmd/mysql2sqlite/`
   - for all supported platforms (where `make` is installed)
      - `make all`
   - for use on Windows
      - `make windows`
   - for use on Linux
     - `make linux`
1. Locate the newly compiled binaries from the applicable `/tmp` subdirectory
   path.
   - if using `Makefile`
     - look in `/tmp/mysql2sqlite/release_assets/check_mysql2sqlite/`
     - look in `/tmp/mysql2sqlite/release_assets/mysql2sqlite/`
   - if using `go build`
     - look in `/tmp/mysql2sqlite/`

See the [deploy](deploy.md) doc for instructions for how to deploy the newly
generated binaries.

## Optional: Use existing binaries

As an alternative to building the binaries yourself, this project also
periodically provides binaries as part of new releases. If binaries for your
platform are not provided, please [file an
issue](https://github.com/atc0005/mysql2sqlite/issues/new) so that we may
evaluate the requirements for providing those binaries with future releases.
