package client

import (
	"fmt"
)

// FlushStdin reads the content of stdin until
// a newline character is hit. This is mainly useful
// in case the user entered gibberish which we
// need to get away from the stream before
// reattempting to read from it
func FlushStdin() {
	var crap string
	fmt.Scanln(&crap)
}
