package config

import (
	"errors"
	"strconv"
	"time"

	"github.com/alyu/configparser"
	"github.com/mrclayman/rest-and-go/gameserver/serverlog"
)

const (
	// defaultListenPort sets the default value
	// for the server's listen port setting
	defaultListenPort uint16 = 8000

	// defaultConnectionTimeout is the default database
	// connection timeout, in seconds
	defaultConnectionTimeout int = 15

	// defaultUseSSL sets the default value for
	// SSL-enabled database connection setting
	defaultUseSSL bool = false

	// defaultLogWS sets the default value for
	// WS interface (mgo library) logging setting
	defaultLogWS bool = false

	// defaultRESTMaxWorkerCount defines the maximum
	// number of workers processing REST interface requests
	defaultRESTMaxWorkerCount uint = 10

	// defaultRESTMaxBacklogLength defines the maximum
	// number of REST requests waiting in the processing
	// queue
	defaultRESTMaxBacklogLength uint = 50
)

// Config structure aggregates configuration
// values for all configurable aspects of the server
type Config struct {
	Net  *NetConfig
	DB   *DatabaseConfig
	Log  *LoggingConfig
	REST *RESTConfig
}

// NetConfig contains configuration data
// on the server's network settings
type NetConfig struct {
	ListenPort uint16
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
	LogWS bool
}

// RESTConfig defines configuration options
// of the REST interface
type RESTConfig struct {
	MaxWorkerCount   uint
	MaxBacklogLength uint
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

	var netCfg *NetConfig
	netCfg, err = parseNetConfig(c)
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

	var restCfg *RESTConfig
	restCfg, err = parseRESTConfig(c)
	if err != nil {
		return nil, err
	}

	return &Config{
		Net:  netCfg,
		DB:   dbCfg,
		Log:  logCfg,
		REST: restCfg,
	}, nil
}

// parseNetConfig parses server's network configuration
// and returns an instance with the parsed information
func parseNetConfig(c *configparser.Configuration) (*NetConfig, error) {
	s, err := c.Section("network")
	if err != nil {
		serverlog.Logger.Println("Network settings not defined, using defaults")
		return &NetConfig{ListenPort: defaultListenPort}, nil
	}

	port := uint64(defaultListenPort)
	if portStr := s.ValueOf("ListenPort"); len(portStr) == 0 {
		serverlog.Logger.Printf("Listen port not defined, using default = %v", port)
	} else if port, err = strconv.ParseUint(portStr, 10, 16); err != nil {
		return nil, errors.New("Invalid port specification, " + portStr + " is not a valid port number")
	}

	return &NetConfig{
		ListenPort: uint16(port),
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
		serverlog.Logger.Println("Warning: No database user specified, assuming default <none>")
	}

	dbPassword := s.ValueOf("Password")
	if len(dbPassword) == 0 {
		serverlog.Logger.Println("Warning: No user password provided, assuming default <none>")
	}

	dbName := s.ValueOf("Database")
	if len(dbName) == 0 {
		return nil, errors.New("No database name provided")
	}

	dbRSName := s.ValueOf("ReplicaSetName")
	if len(dbRSName) == 0 {
		serverlog.Logger.Println("Warning: No replica set name provided, assuming default <none>")
	}

	dbConnTimeout := uint64(defaultConnectionTimeout)
	if dbConnTimeoutStr := s.ValueOf("ConnectionTimeout"); len(dbConnTimeoutStr) == 0 {
		serverlog.Logger.Printf("Warning: No connection timeout defined, using default = %v seconds", dbConnTimeout)
	} else if dbConnTimeout, err = strconv.ParseUint(dbConnTimeoutStr, 10, 32); err != nil {
		return nil, errors.New("Failed to parse connection timeout value: " + err.Error())
	}

	dbUseSSL := defaultUseSSL
	if dbUseSSLStr := s.ValueOf("UseSSL"); len(dbUseSSLStr) == 0 {
		serverlog.Logger.Printf("Warning: SSL use not defined, assuming default = %v", dbUseSSL)
	} else if dbUseSSL, err = strconv.ParseBool(dbUseSSLStr); err != nil {
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

// parseLoggingConfig parses configuration settings pertaining
// to server logging features
func parseLoggingConfig(c *configparser.Configuration) (*LoggingConfig, error) {
	s, err := c.Section("logging")
	if err != nil {
		serverlog.Logger.Println("Logging options not defined, using defaults")
		return &LoggingConfig{LogWS: defaultLogWS}, nil
	}

	logWS := defaultLogWS
	if logWSStr := s.ValueOf("LogWS"); err != nil {
		serverlog.Logger.Printf("LogWS not defined in config, using default = %v", logWS)
	} else if logWS, err = strconv.ParseBool(logWSStr); err != nil {
		return nil, errors.New("Failed to parse LogWS value: " + err.Error())
	}

	return &LoggingConfig{
		LogWS: logWS,
	}, nil
}

// parseRESTConfig parses REST interface configuration options
func parseRESTConfig(c *configparser.Configuration) (*RESTConfig, error) {
	s, err := c.Section("rest")
	if err != nil {
		serverlog.Logger.Println("REST interface options not defined, using defaults")
		return &RESTConfig{
			MaxWorkerCount:   defaultRESTMaxWorkerCount,
			MaxBacklogLength: defaultRESTMaxBacklogLength,
		}, nil
	}

	maxWorkerCount := uint64(defaultRESTMaxWorkerCount)
	if strMaxWorkerCount := s.ValueOf("MaxWorkerCount"); len(strMaxWorkerCount) == 0 {
		serverlog.Logger.Println("Max REST worker count not defined, using default =", maxWorkerCount)
	} else if maxWorkerCount, err = strconv.ParseUint(strMaxWorkerCount, 10, 32); err != nil {
		return nil, errors.New("Invalid max REST worker count value: " + err.Error())
	}

	maxBacklogLength := uint64(defaultRESTMaxBacklogLength)
	if strMaxBacklogLength := s.ValueOf("MaxBacklogLength"); len(strMaxBacklogLength) == 0 {
		serverlog.Logger.Println("Max REST backlog length not defined, using default =", maxBacklogLength)
	} else if maxBacklogLength, err = strconv.ParseUint(strMaxBacklogLength, 10, 32); err != nil {
		return nil, errors.New("Invalid max REST backlog length value: " + err.Error())
	}

	return &RESTConfig{
		MaxBacklogLength: uint(maxBacklogLength),
		MaxWorkerCount:   uint(maxWorkerCount),
	}, nil
}
