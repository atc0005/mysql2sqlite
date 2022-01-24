// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/mysql2sqlite
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/apex/log"
	"github.com/atc0005/mysql2sqlite/internal/caller"
)

// see `constants.go` for other related values

// Version reflects the application version. This is overridden via Makefile
// for release builds.
var Version string = "dev build"

func (c Config) String() string {
	return fmt.Sprintf(
		"UnifiedConfig: { LogLevel: %s, LogFormat: %s, LogOutput: %s, "+
			"ConfigFile: %q, LogDBStats: %v, "+
			"SQLiteDBPath: %q, SQLiteDBFile: %q, CreateIndexes: %v, "+
			"SQLiteBusyTimeout: %v, SQLiteJournalMode: %v, TrimWhitespace: %v, "+
			"MySQLEncryption: %v, MySQLDatabase: %v, MySQLPort: %v, "+
			"MySQLHost: %v, MySQLUsername: %v, MySQLPassword: %v, "+
			"MySQLConnMaxLifetime: %v, MySQLMaxOpenConns: %v, "+
			"MySQLConnMaxIdleTime: %v, MySQLMaxIdleConns: %v }, ",
		c.LogLevel(),
		c.LogFormat(),
		c.LogOutput(),
		c.configFileUsed,
		c.LogDBStats(),
		c.SQLiteDBPath(),
		c.SQLiteDBFile(),
		c.SQLiteCreateIndexes(),
		c.SQLiteBusyTimeout(),
		c.SQLiteJournalMode(),
		c.TrimWhitespace(),
		c.MySQLEncryption(),
		c.MySQLDatabase(),
		c.MySQLPort(),
		c.MySQLHost(),
		c.MySQLUsername(),
		"REDACTED",
		c.MySQLConnMaxLifetime(),
		c.MySQLMaxOpenConns(),
		c.MySQLConnMaxIdleTime(),
		c.MySQLMaxIdleConns(),
	)
}

// Branding provides application name, version and origin
func Branding() string {
	return fmt.Sprintf(
		"%s %s (%s)",
		MyAppName,
		Version,
		MyAppURL,
	)
}

// Version emits version information and associated branding details whenever
// the user specifies the `--version` flag. The application exits after
// displaying this information.
func (c flagSettings) Version() string {
	return fmt.Sprintf("\n%s %s\n%s\n",
		MyAppName, Version, MyAppURL)
}

// Description emits branding information whenever the user specifies the `-h`
// flag. The application uses this as a header prior to displaying available
// CLI flag options.
func (c flagSettings) Description() string {
	return fmt.Sprintf("\n%s", MyAppDescription)
}

// MyBinaryName returns the name of this binary
func MyBinaryName() string {
	return filepath.Base(os.Args[0])
}

// NewConfig is a factory function that produces a new Config object based
// on user provided flag and config file values.
func NewConfig() (*Config, error) {

	myFuncName := caller.GetFuncName()

	var config Config

	// Bundle the returned `*.arg.Parser` for later use from `main()` so that
	// we can explicitly display usage or help details should the
	// user-provided settings fail validation.
	log.Debugf("%s: Parsing flags", myFuncName)
	config.flagParser = arg.MustParse(&config.flagSettings)

	// Apply initial logging settings based on any provided CLI flags. If not
	// specified, default logging settings are used. This is done to permit
	// debug logging to be enabled prior to loading config file settings.
	config.configureLogging()

	//
	// Attempt to load requested config file, fallback to known alternates
	// if user did not specify a config file
	//

	if err := config.Load(); err != nil {
		return nil, fmt.Errorf(
			"%s: failed to load config file: %w",
			myFuncName,
			err,
		)
	}

	// Apply logging settings based on any provided config file settings. If
	// no logging settings were previously provided via CLI flags this allows
	// for config file specified settings to be applied.
	config.configureLogging()

	log.Debug("Validating configuration ...")
	if err := config.Validate(); err != nil {
		return nil, err
	}
	log.Debug("Configuration validated")

	// log.Debugf("Config object: %v", config.String())

	return &config, nil

}

// Load attempts to first load the user-specified config file, then falls back
// to checking for a config file in the directory alongside the executable,
// then finally a config file from the user's configuration path. An error is
// returned if the configuration file cannot be loaded.
func (c *Config) Load() error {

	configFiles := make([]string, 0, 3)

	if c.ConfigFile() == "" {
		log.Debug("User-specified config file not provided")
	} else {
		log.Debug("User-specified config file provided, will attempt to load it")
		configFiles = append(configFiles, c.ConfigFile())
	}

	localFile, err := c.localConfigFile()
	if err != nil {
		log.Error("Failed to determine path to local file")
	}
	configFiles = append(configFiles, localFile)

	userConfigFile, err := c.userConfigFile()
	if err != nil {
		log.Error("Failed to determine path to user config file")
	}
	configFiles = append(configFiles, userConfigFile)

	for _, file := range configFiles {
		log.Debugf("Trying to load config file %q", file)
		ok, err := c.loadConfigFile(file)
		if ok {
			c.configFileLoaded = true
			c.configFileUsed = file
			// if there were no errors, we're done loading config files
			log.WithFields(log.Fields{
				"config_file": file,
			}).Debug("Config file successfully loaded")
			log.Debug("Config file successfully parsed")
			log.Debugf("After loading config file: %v", c.String())
			break
		}

		log.Warnf("Config file %q not found or unable to load", file)
		log.WithFields(log.Fields{
			"error": err,
		}).Debug("")
	}

	if !c.configFileLoaded {
		return ErrCfgFileNotFound
	}

	return nil
}
