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
func GetJSONFromRequest(req *http.Request, out interface{}) error {
	if req == nil {
		return RequestError{Message: "Nil request"}
	} else if contType := req.Header.Get("Content-Type"); contType != "application/json" {
		return RequestError{Message: "Wrong content type, expected 'application/json'"}
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return RequestError{"Failed to read request body"}
	}

	if err = json.Unmarshal(body, &out); err != nil {
		return RequestError{"Invalid JSON structure in request body"}
	}

	return nil
}

// WriteJSONToResponse writes a generic structure
// into a JSON string and writes that string into
// the HTTP response object, also defining the
// response's Content-Type header
func WriteJSONToResponse(resp http.ResponseWriter, in interface{}) error {
	buf, err := json.Marshal(in)
	if err != nil {
		return err
	}

	resp.Header().Set("Content-Type", "application/json")
	_, err = resp.Write(buf)
	return err
}
