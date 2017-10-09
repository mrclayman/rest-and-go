package util

import (
	"net/http"
	"path"
	"strings"

	"github.com/mrclayman/rest_api_test/errors"
)

// SplitPath splits off the first component of p, which will be cleaned of
// relative components before processing. On return, head will never contain
// a slash and tail will always be a rooted path without a trailing slash.
func SplitPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

// GetJSONFromRequest parses the body of the request
// and transforms it into a JSON structure for further
// processing
func GetJSONFromRequest(request *http.Request) ([]byte, error) {
	if request == nil {
		return nil, errors.RequestError{Message: "Nil request"}
	}

	const maxSize int = 256
	body := make([]byte, maxSize)
	var bytesRead int
	var err error
	// TODO Finish once you figure out how to effectively read the
	// entire body without knowing its size in advance
}
