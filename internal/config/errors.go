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

var ErrCfgFileNotFound = errors.New(
	"failed to load config file from known locations. Please " +
		"specify a path to a config file or ensure that one is present " +
		"in a known location. See the README file for additional details.",
)
