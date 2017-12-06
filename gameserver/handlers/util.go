package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/mrclayman/rest-and-go/gameserver/core/auth"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
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
	return GetJSONFromBytes(body, out)
}

// GetJSONFromBytes unmarshals JSON data from a byte slice
func GetJSONFromBytes(bytes []byte, out interface{}) error {
	if err := json.Unmarshal(bytes, &out); err != nil {
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
	log.Printf("Sending JSON reply: %v", string(buf))

	resp.Header().Set("Content-Type", "application/json")
	_, err = resp.Write(buf)
	return err
}

// getValueFromGET obtains a single value from
// the request's GET arguments. It is expected that exactly
// one value is present in the arguments and that the GET
// arguments have been parsed beforehand using ParseForm()
func getValueFromGET(req *http.Request, name string) (string, bool) {
	value := req.Form.Get(name)
	var ok bool
	if len(value) > 0 {
		ok = true
	}

	return value, ok
}

// GetPlayerDataFromGET reads the player's identification
// data from a GET request. The method does not check that the
// request is indeed a GET request and it also assumes the GET
// arguments have already been parsed using a call to ParseForm()
func GetPlayerDataFromGET(req *http.Request) (player.ID, auth.AuthToken, error) {
	ID := player.InvalidID
	token := auth.InvalidAuthToken
	var err error

	if strPlayerID, ok := getValueFromGET(req, "id"); !ok {
		return player.InvalidID, auth.InvalidAuthToken,
			RequestError{"Failed to obtain player ID from request"}
	} else if ID, err = player.StringToID(strPlayerID); err != nil {
		return player.InvalidID, auth.InvalidAuthToken,
			RequestError{"Failed to convert argument to player ID"}
	}

	var strToken string
	var ok bool
	if strToken, ok = getValueFromGET(req, "token"); !ok {
		return player.InvalidID, auth.InvalidAuthToken,
			RequestError{"Failed to obtain player's authentication token from request"}
	}

	token = auth.StringToAuthToken(strToken)

	return ID, token, nil
}
