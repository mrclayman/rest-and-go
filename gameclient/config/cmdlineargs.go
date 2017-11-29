package config

import (
	"errors"
	"flag"
)

// CmdLineArgs aggregates all recognized
// command line arguments
type CmdLineArgs struct {
	CfgFilePath string
}

// ParseCmdLineArgs parses the command line arguments
// provided by the user
func ParseCmdLineArgs() (*CmdLineArgs, error) {
	r := CmdLineArgs{}

	flag.StringVar(&r.CfgFilePath, "config", "client.ini", "Path to client configuration file")
	flag.StringVar(&r.CfgFilePath, "c", "client.ini", "Path to client configuration file")

	flag.Parse()

	if len(r.CfgFilePath) == 0 {
		return nil, errors.New("Missing configuration file specification")
	}

	return &r, nil
}
