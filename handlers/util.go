package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
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
func GetJSONFromRequest(req *http.Request, out *interface{}) error {
	if req == nil {
		return RequestError{Message: "Nil request"}
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return RequestError{"Failed to read request body"}
	}

	if err = json.Unmarshal(body, out); err != nil {
		return RequestError{"Invalid JSON structure in request body"}
	}

	return nil
}
