package config

import (
	"errors"
	"flag"
)

// CmdLineArgs contains all recognized
// command line argument values
type CmdLineArgs struct {
	CfgFilePath string
}

// ParseCmdLineArgs parses command line
// arguments present and stores their
// values in an instance of CmdLineArgs
func ParseCmdLineArgs() (*CmdLineArgs, error) {
	r := CmdLineArgs{}

	flag.StringVar(&r.CfgFilePath, "config", "server.ini", "Path to the configuration file")
	flag.StringVar(&r.CfgFilePath, "c", "server.ini", "Path to the configuration file (shorthand)")

	var help bool
	flag.BoolVar(&help, "help", false, "Display usage instructions")
	flag.BoolVar(&help, "h", false, "Display usage instructions (shorthand)")

	flag.Parse()

	if help {
		flag.Usage()
		return nil, nil
	}

	if len(r.CfgFilePath) == 0 {
		return nil, errors.New("Missing configuration file specification")
	}

	return &r, nil
}
