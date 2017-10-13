package handlers

import (
	"net/http"

	"github.com/mrclayman/rest_api_test/core"
)

// loginCredentials provides the storage
// and format for JSON data sent in the
// login request
type login struct {
	name     string
	password string
}

// LoginHandler handles login POST requests by verifying
// the client's credentials. Upon successful verification,
// the client is issued an authorization token used for
// further communication. In the case of an unsuccessful
// authorization, the HTTP code 403/Forbidden is returned
type LoginHandler struct {
	db *core.Database
}

// ProcessRequest handles the login POST request
func (h *LoginHandler) ProcessRequest(resp http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var cred login
	err := GetJSONFromRequest(req, &cred)
	id, ok := h.db.AuthenticatePlayer(login.name, login.password)
}
