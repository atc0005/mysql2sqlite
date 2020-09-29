// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"errors"
)

// see `constants.go` for other related values

// ErrCfgFileNotFound is a fatal error; this error indicates that the
// user-specified config file could not be found, nor one of the paths
// automatically checked by this application.
var ErrCfgFileNotFound = errors.New(
	"failed to load config file from known locations; " +
		"see the README file for additional details",
)
