package config

import (
	"errors"
	"strconv"
	"time"

	"github.com/alyu/configparser"
)

// ConnectionConfig aggregates configuration
// information on how to connect to the server
type ConnectionConfig struct {
	ServerURL string
	Timeout   time.Duration
}

// Config aggregates all configuration
// options parsed from the configuration file
type Config struct {
	Conn *ConnectionConfig
}

// ParseCfgFile parses the configuration
// file stored at the path provided in the
// argument, and stores the contents in
// an instance of the Config structure
func ParseCfgFile(ifPath string) (*Config, error) {
	c, err := configparser.Read(ifPath)
	if err != nil {
		return nil, err
	}

	var conn *ConnectionConfig
	conn, err = parseConnConfig(c)
	if err != nil {
		return nil, err
	}

	return &Config{
		Conn: conn,
	}, nil
}

// parseConnConfig parses server connection-related informatio
func parseConnConfig(c *configparser.Configuration) (*ConnectionConfig, error) {
	s, err := c.Section("connection")
	if err != nil {
		return nil, err
	}

	serverURL := s.ValueOf("ServerURL")
	if len(serverURL) == 0 {
		return nil, errors.New("No server URL defined")
	}

	var timeout uint64
	timeout, err = strconv.ParseUint(s.ValueOf("ConnectionTimeout"), 10, 32)
	if err != nil {
		return nil, err
	}

	return &ConnectionConfig{
		ServerURL: serverURL,
		Timeout:   time.Duration(int64(timeout)) * time.Second,
	}, nil
}
