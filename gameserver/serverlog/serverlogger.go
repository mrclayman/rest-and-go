package serverlog

import (
	"log"
	"os"
)

// Logger defines a server-wide logger object
var Logger = log.New(os.Stderr, "", log.LstdFlags)
