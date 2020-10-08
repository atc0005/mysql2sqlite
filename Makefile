# Copyright 2020 Adam Chalkley
#
# https://github.com/atc0005/mysql2sqlite
#
# Licensed under the MIT License. See LICENSE file in the project root for
# full license information.

# References:
#
# https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
# https://github.com/mapnik/sphinx-docs/blob/master/Makefile
# https://stackoverflow.com/questions/23843106/how-to-set-child-process-environment-variable-in-makefile
# https://stackoverflow.com/questions/3267145/makefile-execute-another-target
# https://unix.stackexchange.com/questions/124386/using-a-make-rule-to-call-another
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
# https://www.gnu.org/software/make/manual/html_node/Recipe-Syntax.html#Recipe-Syntax
# https://www.gnu.org/software/make/manual/html_node/Special-Variables.html#Special-Variables
# https://danishpraka.sh/2019/12/07/using-makefiles-for-go.html
# https://gist.github.com/subfuzion/0bd969d08fe0d8b5cc4b23c795854a13
# https://stackoverflow.com/questions/10858261/abort-makefile-if-variable-not-set
# https://stackoverflow.com/questions/38801796/makefile-set-if-variable-is-empty
# https://www.gnu.org/software/make/manual/make.html#Flavors
# https://stackoverflow.com/questions/6283320/vs-in-make-macros


#############################################################################
# "Variables defined with := in GNU make are expanded when they are defined
# rather than when they are used."
#
# https://stackoverflow.com/a/6283363/903870
#############################################################################

SHELL 					:= /bin/bash

# Space-separated list of cmd/BINARY_NAME directories to build
WHAT 					:= mysql2sqlite check_mysql2sqlite

# What package holds the "version" variable used in branding/version output?
# VERSION_VAR_PKG			:= $(shell go list .)
# VERSION_VAR_PKG			:= main
VERSION_VAR_PKG			:= $(shell go list .)/internal/config

OUTPUTDIR 				:= release_assets

# https://gist.github.com/TheHippo/7e4d9ec4b7ed4c0d7a39839e6800cc16
VERSION 				:= $(shell git describe --always --long --dirty)

# The default `go build` process embeds debugging information. Building
# without that debugging information reduces the binary size by around 28%.
#
# We also include additional flags in an effort to generate static binaries
# that do not have external dependencies. As of Go 1.15 this still appears to
# be a mixed bag, so YMMV.
#
# See https://github.com/golang/go/issues/26492 for more information.
#
# -s
#	Omit the symbol table and debug information.
#
# -w
#	Omit the DWARF symbol table.
#
# -tags 'osusergo,netgo'
#	Use pure Go implementation of user and group id/name resolution.
#	Use pure Go implementation of DNS resolver.
#
# -trimpath
#	https://golang.org/cmd/go/
#   removes all file system paths from the compiled executable, to improve
#   build reproducibility.
#
# -extldflags '-static'
#	Pass 'static' flag to external linker.
#
# -linkmode=external
#	https://golang.org/src/cmd/cgo/doc.go
#
#   NOTE: Using external linker requires installation of `gcc-multilib`
#   package when building 32-bit binaries on a Debian/Ubuntu system. It also
#   seems to result in an unstable build that crashes on startup. This *might*
#   be specific to the WSL environment used for builds. Further testing is
#   needed to confirm.
#
# CGO_ENABLED=1
#   CGO is disabled by default for cross-compilation. You need to enable it
#   explicitly to use CGO for multiple architectures.
BUILD_LDFLAGS_COMMON	:= -s -w -X $(VERSION_VAR_PKG).Version=$(VERSION)
BUILD_LDFLAGS_STATIC	:= -linkmode=external -extldflags '-static'
BUILDCMD_COMMON			:= CGO_ENABLED=1 go build -mod=vendor -a -trimpath
BUILDCMD_STATIC			:= $(BUILDCMD_COMMON) -tags 'osusergo,netgo,sqlite_omit_load_extension' -ldflags "$(BUILD_LDFLAGS_STATIC) $(BUILD_LDFLAGS_COMMON)"
BUILDCMD_DYNAMIC		:= $(BUILDCMD_COMMON) -ldflags "$(BUILD_LDFLAGS_COMMON)"

BUILD_TYPE_STATIC		:= static
BUILD_TYPE_DYNAMIC		:= dynamic

# Default build command and type if not overridden
BUILDCMD 				:= $(BUILDCMD_DYNAMIC)
BUILDTYPE				:= $(BUILD_TYPE_DYNAMIC)

# Use mingw as C compiler to build Windows cgo-enabled binaries.
WINCOMPILERX86 			:= 	CC=i686-w64-mingw32-gcc
WINCOMPILERX64 			:= 	CC=x86_64-w64-mingw32-gcc

DOCKER_BUILD_IMG_X86	:= atc0005/go-ci:go-ci-stable-alpine-buildx86
DOCKER_BUILD_IMG_X64	:= atc0005/go-ci:go-ci-stable-alpine-buildx64

GOCLEANCMD				:=	go clean -mod=vendor ./...
GITCLEANCMD				:= 	git clean -xfd
CHECKSUMCMD				:=	sha256sum -b

.DEFAULT_GOAL 			:= help

  ##########################################################################
  # Targets will not work properly if a file with the same name is ever
  # created in this directory. We explicitly declare our targets to be phony
  # by making them a prerequisite of the special target .PHONY
  ##########################################################################

.PHONY: help
## help: prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: lintinstall
## lintinstall: install common linting tools
# https://github.com/golang/go/issues/30515#issuecomment-582044819
lintinstall:
	@echo "Installing linting tools"

	@export PATH="${PATH}:$(go env GOPATH)/bin"

	@echo "Explicitly enabling Go modules mode per command"
	(cd; GO111MODULE="on" go get honnef.co/go/tools/cmd/staticcheck)

	@echo Installing latest stable golangci-lint version per official installation script ...
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin
	golangci-lint --version

	@echo "Finished updating linting tools"

.PHONY: linting
## linting: runs common linting checks
linting:
	@echo "Running linting tools ..."

	@echo "Running go vet ..."
	@go vet -mod=vendor $(shell go list -mod=vendor ./... | grep -v /vendor/)

	@echo "Running golangci-lint ..."
	@golangci-lint run

	@echo "Running staticcheck ..."
	@staticcheck $(shell go list -mod=vendor ./... | grep -v /vendor/)

	@echo "Finished running linting checks"

.PHONY: gotests
## gotests: runs go test recursively, verbosely
gotests:
	@echo "Running go tests ..."
	@go test -mod=vendor ./...
	@echo "Finished running go tests"

.PHONY: goclean
## goclean: removes local build artifacts, temporary files, etc
goclean:
	@echo "Removing object files and cached files ..."
	@$(GOCLEANCMD)
	@echo "Removing any existing release assets"
	@mkdir -p "$(OUTPUTDIR)"
	@rm -vf $(wildcard ${OUTPUTDIR}/*/*-linux-*)
	@rm -vf $(wildcard ${OUTPUTDIR}/*/*-windows-*)

.PHONY: clean
## clean: alias for goclean
clean: goclean

.PHONY: gitclean
## gitclean: WARNING - recursively cleans working tree by removing non-versioned files
gitclean:
	@echo "Removing non-versioned files ..."
	@$(GITCLEANCMD)

.PHONY: pristine
## pristine: run goclean and gitclean to remove local changes
pristine: goclean gitclean

.PHONY: all
# https://stackoverflow.com/questions/3267145/makefile-execute-another-target
## all: generates dynamically linked assets for Linux and Windows systems
all: clean windows linux
	@echo "Completed all cross-platform builds ..."

.PHONE: all-static
## all-static: generates statically linked x86 and x64 assets for Linux and Windows systems
all-static: clean windows-static linux-static
	@echo "Completed all cross-platform builds ..."

.PHONY: windows
## windows: generates dynamically linked x86 and x64 Windows assets
windows: windows-x86 windows-x64
	@echo "Completed build tasks for windows"

.PHONY: windows-static
## windows-static: generates dynamically linked x86 and x64 Windows assets
windows-static: windows-x86-static windows-x64-static
	@echo "Completed build tasks for windows"

.PHONY: windows-x86
## windows-x86: generates dynamically linked Windows x86 assets
windows-x86:
	@echo "Building ($(BUILDTYPE)) release assets for windows x86 ..."
	@for target in $(WHAT); do \
		mkdir -p $(OUTPUTDIR)/$$target && \
		echo "Building $$target 386 binaries" && \
		env GOOS=windows GOARCH=386 $(WINCOMPILERX86) $(BUILDCMD) -o $(OUTPUTDIR)/$$target/$$target-$(VERSION)-windows-386.exe ./cmd/$$target && \
		echo "Generating $$target x86 checksum files" && \
		cd $(OUTPUTDIR)/$$target && \
		$(CHECKSUMCMD) $$target-$(VERSION)-windows-386.exe > $$target-$(VERSION)-windows-386.exe.sha256 && \
		cd $$OLDPWD; \
	done
	@echo "Completed ($(BUILDTYPE)) release assets build tasks for windows x86"

.PHONY: windows-x86-static
## windows-x86-static: generates assets statically, specifically for Windows x86 systems
windows-x86-static: BUILDCMD = $(BUILDCMD_STATIC)
windows-x86-static: BUILDTYPE = $(BUILD_TYPE_STATIC)
windows-x86-static: windows-x86

.PHONY: windows-x64
## windows-x64: generates assets specifically for x64 Windows systems
windows-x64:
	@echo "Building ($(BUILDTYPE)) release assets for windows x64 ..."
	@for target in $(WHAT); do \
		mkdir -p $(OUTPUTDIR)/$$target && \
		echo "Building $$target amd64 binaries" && \
		env GOOS=windows GOARCH=amd64 $(WINCOMPILERX64) $(BUILDCMD) -o $(OUTPUTDIR)/$$target/$$target-$(VERSION)-windows-amd64.exe ./cmd/$$target && \
		echo "Generating $$target checksum files" && \
		cd $(OUTPUTDIR)/$$target && \
		$(CHECKSUMCMD) $$target-$(VERSION)-windows-amd64.exe > $$target-$(VERSION)-windows-amd64.exe.sha256 && \
		cd $$OLDPWD; \
	done
	@echo "Completed ($(BUILDTYPE)) release assets build tasks for windows x64"

.PHONY: windows-x64-static
## windows-x64-static: generates assets statically, specifically for Windows x64 systems
windows-x64-static: BUILDCMD = $(BUILDCMD_STATIC)
windows-x64-static: BUILDTYPE = $(BUILD_TYPE_STATIC)
windows-x64-static: windows-x64

.PHONY: linux
## linux: generates dynamically linked x86 and x64 assets for Linux distros
linux: linux-x86 linux-x64
	@echo "Completed ($(BUILDTYPE)) release assets build tasks for linux"

.PHONE: linux-static
## linux-static: generates statically linked x86 and x64 assets for Linux distros
linux-static: linux-x86-static linux-x64-static
	@echo "Completed ($(BUILDTYPE)) release assets build tasks for linux"

.PHONY: linux-x86
## linux-x86: generates assets specifically for Linux x86 systems
linux-x86:
	@echo "Building ($(BUILDTYPE)) release assets for linux x86 ..."
	@for target in $(WHAT); do \
		mkdir -p $(OUTPUTDIR)/$$target && \
		echo "Building $$target 386 binaries" && \
		env GOOS=linux GOARCH=386 $(BUILDCMD) -o $(OUTPUTDIR)/$$target/$$target-$(VERSION)-linux-386 ./cmd/$$target && \
		echo "Generating $$target checksum files" && \
		cd $(OUTPUTDIR)/$$target && \
		$(CHECKSUMCMD) $$target-$(VERSION)-linux-386 > $$target-$(VERSION)-linux-386.sha256 && \
		cd $$OLDPWD; \
	done
	@echo "Completed ($(BUILDTYPE)) release assets build tasks for linux x86"

.PHONY: linux-x86-static
## linux-x86-static: generates assets statically, specifically for Linux x86 systems
linux-x86-static: BUILDCMD = $(BUILDCMD_STATIC)
linux-x86-static: BUILDTYPE = $(BUILD_TYPE_STATIC)
linux-x86-static: linux-x86

.PHONY: linux-x64
## linux-x64: generates assets specifically for Linux x64 systems
linux-x64:
	@echo "Building ($(BUILDTYPE)) release assets for linux x64 ..."
	@for target in $(WHAT); do \
		mkdir -p $(OUTPUTDIR)/$$target && \
		echo "Building $$target amd64 binaries" && \
		env GOOS=linux GOARCH=amd64 $(BUILDCMD) -o $(OUTPUTDIR)/$$target/$$target-$(VERSION)-linux-amd64 ./cmd/$$target && \
		echo "Generating $$target checksum files" && \
		cd $(OUTPUTDIR)/$$target && \
		$(CHECKSUMCMD) $$target-$(VERSION)-linux-amd64 > $$target-$(VERSION)-linux-amd64.sha256 && \
		cd $$OLDPWD; \
	done
	@echo "Completed ($(BUILDTYPE)) release assets build tasks for linux x64"


.PHONY: linux-x64-static
## linux-x64-static: generates assets statically, specifically for Linux x64 systems
linux-x64-static: BUILDCMD = $(BUILDCMD_STATIC)
linux-x64-static: BUILDTYPE = $(BUILD_TYPE_STATIC)
linux-x64-static: linux-x64

.PHONY: docker
## docker: generates assets for Linux distros and Windows using Docker
docker: clean
	@docker run \
		--rm \
		-i \
		-v $$PWD:$$PWD \
		-w $$PWD \
		$(DOCKER_BUILD_IMG_X86) \
		make windows-x86-static linux-x86-static
	@docker run \
		--rm \
		-i \
		-v $$PWD:$$PWD \
		-w $$PWD \
		$(DOCKER_BUILD_IMG_X64) \
		make windows-x64-static linux-x64-static
	@echo "Completed all cross-platform builds via Docker containers ..."
