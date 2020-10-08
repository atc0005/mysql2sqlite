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
     - `sudo apt-get install make bsdmainutils gcc gcc-multilib gcc-mingw-w64`
   - for Windows
     - Emulated environments (*easier*)
       - build using Windows Subsystem for Linux Ubuntu environment and just
         copy out the Windows binaries from that environment or build within a
         path accessible from both Windows and WSL (e.g.,
         `/mnt/c/Users/YOUR_USERNAME/Desktop/mysql2sqlite`)
       - If already running a Docker environment, use a container with the Go
         tool-chain already installed along with the necessary packages to
         allow cross-compilation. If using the official `golang` Docker image,
         you will need to install the same packages listed previously for
         Ubuntu. Alternatively, you may also use the Alpine-based Docker images created for statically linked, cgo-enabled cross-compilation.
         - `atc0005/go-ci:go-ci-stable-alpine-buildx86`
         - `atc0005/go-ci:go-ci-stable-alpine-buildx64`
       - If already familiar with LXD, create an Ubuntu  container and follow
         the installation steps given previously to install required
         dependencies
     - Native tooling (*harder*)
       - see the StackOverflow Question `32127524` link in the
         [References](references.md) section for potential options for
         installing `make` on Windows
       - see the mingw-w64 project homepage link in the
         [References](references.md) section for options for installing `gcc`
         and related packages on Windows
1. Build binaries
   - dynamically linked
     - for the current operating system, explicitly using
         bundled dependencies in top-level `vendor` folder
       - `CGO_ENABLED=1 go build -mod=vendor ./cmd/check_mysql2sqlite/`
       - `CGO_ENABLED=1 go build -mod=vendor ./cmd/mysql2sqlite/`
     - for all supported platforms (where `make` is installed)
       - `make all`
   - statically linked
     - for all supported platforms
       - `make all-static`
       - `make docker`
         - requires that you have a working Docker installation first
         - links against `musl libc` (smaller) instead of `glibc` (more
           common, default option for dynamic linkage)
     - for just windows
       - `make windows-static`
     - for just linux
       - `make linux-static`
   - run `make` without options to see the full list of supported "recipes"
1. Locate the newly compiled binaries from the applicable `/tmp` subdirectory
   path.
   - if using `Makefile`
     - look in `/tmp/mysql2sqlite/release_assets/check_mysql2sqlite/`
     - look in `/tmp/mysql2sqlite/release_assets/mysql2sqlite/`
   - if using `go build` (with options provided earlier)
     - look in `/tmp/mysql2sqlite/`

See the [deploy](deploy.md) doc for instructions for how to deploy the newly
generated binaries.

## Optional: Use existing binaries

As an alternative to building the binaries yourself, this project also
periodically provides binaries as part of new releases. If binaries for your
platform are not provided, please [file an
issue](https://github.com/atc0005/mysql2sqlite/issues/new) so that we may
evaluate the requirements for providing those binaries with future releases.
