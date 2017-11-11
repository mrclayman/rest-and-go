package client

import (
	"bufio"
	"fmt"
	"os"
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

// ReadLine reads a complete line (including white spaces)
// until a newline character is seen
func ReadLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return ""
	}

	return scanner.Text()
}
