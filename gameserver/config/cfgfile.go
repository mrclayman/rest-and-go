package config

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/alyu/configparser"
)

// Config structure aggregates configuration
// values for all aspects of the server
type Config struct {
	DB  *DatabaseConfig
	Log *LoggingConfig
}

// DatabaseConfig contains configuration
// values for database connection
type DatabaseConfig struct {
	Host              string
	User              string
	Password          string
	Database          string
	RSName            string
	ConnectionTimeout time.Duration
	SSL               bool
}

// LoggingConfig stores configuration options
// pertaining to logging of events within the
// server
type LoggingConfig struct {
	LogMgo bool
}

// Cfg is the global configuration
// struct instance
var Cfg *Config

// ParseCfgFile parses the configuration
// file stored at ifPath and returns
// an instance of the Config struct
// with the proper values set.
func ParseCfgFile(ifPath string) (*Config, error) {
	c, err := configparser.Read(ifPath)
	if err != nil {
		return nil, err
	}

	var dbCfg *DatabaseConfig
	dbCfg, err = parseDatabaseConfig(c)
	if err != nil {
		return nil, err
	}

	var logCfg *LoggingConfig
	logCfg, err = parseLoggingConfig(c)
	if err != nil {
		return nil, err
	}

	return &Config{
		DB:  dbCfg,
		Log: logCfg,
	}, nil
}

// parseDatabaseConfig parses the database configuration
// part of the configuration file
func parseDatabaseConfig(c *configparser.Configuration) (*DatabaseConfig, error) {
	s, err := c.Section("database")
	if err != nil {
		return nil, err
	}

	dbHost := s.ValueOf("Host")
	if len(dbHost) == 0 {
		return nil, errors.New("No server hostname specified")
	}

	dbUser := s.ValueOf("User")
	if len(dbUser) == 0 {
		log.Println("Warning: No database user specified")
	}

	dbPassword := s.ValueOf("Password")
	if len(dbPassword) == 0 {
		log.Println("Warning: No user password provided")
	}

	dbName := s.ValueOf("Database")
	if len(dbName) == 0 {
		return nil, errors.New("No database name provided")
	}

	dbRSName := s.ValueOf("ReplicaSetName")
	if len(dbRSName) == 0 {
		log.Println("Warning: No replica set name provided")
	}

	dbConnTimeoutStr := s.ValueOf("ConnectionTimeout")
	if len(dbConnTimeoutStr) == 0 {
		dbConnTimeoutStr = "15"
		log.Printf("Warning: No connection timeout defined, using default of %v seconds", dbConnTimeoutStr)
	}

	var dbConnTimeout uint64
	dbConnTimeout, err = strconv.ParseUint(dbConnTimeoutStr, 10, 32)
	if err != nil {
		return nil, errors.New("Failed to parse connection timeout value: " + err.Error())
	}

	dbUseSSLStr := s.ValueOf("UseSSL")
	if len(dbUseSSLStr) == 0 {
		dbUseSSLStr = "false"
		log.Println("Warning: SSL use not defined, assuming " + dbUseSSLStr)
	}
	var dbUseSSL bool
	dbUseSSL, err = strconv.ParseBool(dbUseSSLStr)
	if err != nil {
		return nil, errors.New("Failed to parse UseSSL value: " + err.Error())
	}

	return &DatabaseConfig{
		Host:              dbHost,
		User:              dbUser,
		Password:          dbPassword,
		Database:          dbName,
		RSName:            dbRSName,
		ConnectionTimeout: time.Duration(int64(dbConnTimeout)) * time.Second,
		SSL:               dbUseSSL,
	}, nil
}

func parseLoggingConfig(c *configparser.Configuration) (*LoggingConfig, error) {
	s, err := c.Section("logging")
	if err != nil {
		log.Println("Logging options not defined, using defaults")
		return &LoggingConfig{}, nil
	}

	logMgoStr := s.ValueOf("LogMgo")
	if err != nil {
		logMgoStr = "false"
		log.Printf("LogMgo not defined in config, using default = %v", logMgoStr)
	}
	var logMgo bool
	logMgo, err = strconv.ParseBool(logMgoStr)
	if err != nil {
		return nil, errors.New("Failed to parse LogMgo value: " + err.Error())
	}

	return &LoggingConfig{
		LogMgo: logMgo,
	}, nil
}
