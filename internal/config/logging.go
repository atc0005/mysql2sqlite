// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/discard"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/logfmt"
	"github.com/apex/log/handlers/text"
)

// Log levels
const (
	// https://godoc.org/github.com/apex/log#Level

	// LogLevelFatal is used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	LogLevelFatal string = "fatal"

	// LogLevelError is for errors that should definitely be noted.
	LogLevelError string = "error"

	// LogLevelWarn is for non-critical entries that deserve eyes.
	LogLevelWarn string = "warn"

	// LogLevelInfo is for general application operational entries.
	LogLevelInfo string = "info"

	// LogLevelDebug is for debug-level messages and is usually enabled
	// when debugging. Very verbose logging.
	LogLevelDebug string = "debug"
)

// apex/log Handlers
//
// ---------------------------------------------------------
// cli - human-friendly CLI output
// discard - discards all logs
// es - Elasticsearch handler
// graylog - Graylog handler
// json - JSON output handler
// kinesis - AWS Kinesis handler
// level - level filter handler
// logfmt - logfmt plain-text formatter
// memory - in-memory handler for tests
// multi - fan-out to multiple handlers
// papertrail - Papertrail handler
// text - human-friendly colored output
// delta - outputs the delta between log calls and spinner
const (
	// LogFormatCLI provides human-friendly CLI output
	LogFormatCLI string = "cli"

	// LogFormatJSON provides JSON output
	LogFormatJSON string = "json"

	// LogFormatLogFmt provides logfmt plain-text output
	LogFormatLogFmt string = "logfmt"

	// LogFormatText provides human-friendly colored output
	LogFormatText string = "text"

	// LogFormatDiscard discards all logs
	LogFormatDiscard string = "discard"
)

// Log output targets. The chosen output target is user-configurable.
const (
	LogOutputStdout string = "stdout"
	LogOutputStderr string = "stderr"
)

// configureLogging is a wrapper function to enable setting requested logging
// settings.
func (c Config) configureLogging() {

	switch c.LogLevel() {
	case LogLevelFatal:
		log.SetLevel(log.FatalLevel)
	case LogLevelError:
		log.SetLevel(log.ErrorLevel)
	case LogLevelWarn:
		log.SetLevel(log.WarnLevel)
	case LogLevelInfo:
		log.SetLevel(log.InfoLevel)
	case LogLevelDebug:
		log.SetLevel(log.DebugLevel)
	}

	// Apply user-specified logging output target
	var outputTarget *os.File
	switch c.LogOutput() {
	case LogOutputStdout:
		outputTarget = os.Stdout
	case LogOutputStderr:
		outputTarget = os.Stderr
	default:
		outputTarget = os.Stdout
	}

	switch c.LogFormat() {
	case LogFormatText:
		log.SetHandler(text.New(outputTarget))
	case LogFormatCLI:
		log.SetHandler(cli.New(outputTarget))
	case LogFormatLogFmt:
		log.SetHandler(logfmt.New(outputTarget))
	case LogFormatJSON:
		log.SetHandler(json.New(outputTarget))
	case LogFormatDiscard:
		log.SetHandler(discard.New())
	}

}
