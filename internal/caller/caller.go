// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package caller

import (
	"fmt"
	"runtime"
)

// GetFuncFileLineInfo is a wrapper around the runtime.Caller() function. This
// function returns the calling function name, filename and line number to
// help with debugging efforts.
func GetFuncFileLineInfo() string {

	if pc, file, line, ok := runtime.Caller(1); ok {
		return fmt.Sprintf(
			"func %s called (from %q, line %d): ",
			runtime.FuncForPC(pc).Name(),
			file,
			line,
		)
	}

	return "error: unable to recover caller origin via runtime.Caller()"
}

// GetFuncName is a wrapper around the runtime.Caller() function. This
// function returns the calling function name and discards other return
// values.
func GetFuncName() string {

	if pc, _, _, ok := runtime.Caller(1); ok {
		return runtime.FuncForPC(pc).Name()
	}

	return "error: unable to recover caller origin via runtime.Caller()"
}

// GetParentFuncFileLineInfo is a wrapper around the runtime.Caller()
// function. This function returns the parent calling function name, filename
// and line number to help with debugging efforts.
func GetParentFuncFileLineInfo() string {

	if pc, file, line, ok := runtime.Caller(2); ok {
		return fmt.Sprintf(
			"func %s called (from %q, line %d): ",
			runtime.FuncForPC(pc).Name(),
			file,
			line,
		)
	}

	return "error: unable to recover caller parent origin via runtime.Caller()"
}

// GetParentFuncName is a wrapper around the runtime.Caller() function. This
// function returns the parent calling function name and discards other return
// values.
func GetParentFuncName() string {

	if pc, _, _, ok := runtime.Caller(2); ok {
		return runtime.FuncForPC(pc).Name()
	}

	return "error: unable to recover caller parent origin via runtime.Caller()"
}
