package handlers

import (
	"net/http"

	"github.com/mrclayman/rest-and-go/server/core"
)

// login provides the storage
// and format for JSON data sent in the
// login request
type loginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// LoginHandler handles login POST requests by verifying
// the client's credentials. Upon successful verification,
// the client is issued an authorization token used for
// further communication. In the case of an unsuccessful
// authorization, the HTTP code 403/Forbidden is returned
type LoginHandler struct {
	core *core.Core
}

// NewLoginHandler creates a pointer to a new
// instance of the login request handler structure
func NewLoginHandler(c *core.Core) *LoginHandler {
	return &LoginHandler{core: c}
}

// ProcessRequest handles the login POST request
func (h *LoginHandler) ProcessRequest(resp http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var credent loginRequest
	err := GetJSONFromRequest(req, &credent)
	if err != nil {
		http.Error(resp, err.(RequestError).Message, http.StatusBadRequest)
		return
	}

	// Authenticate
	id, ok := h.core.AuthenticatePlayer(credent.Name, credent.Password)
	if !ok {
		http.Error(resp, "Failed to authenticate user", http.StatusForbidden)
		return
	}

	// Success, generate a token for the player and add them
	token := core.GenerateAuthenticationToken()
	h.core.AddConnected(id, credent.Name, token)

	// Dispatch the response to the player
	respData := map[string]interface{}{"id": id, "token": token}
	// TODO Should I check that the response has been serialized correctly?
	WriteJSONToResponse(resp, respData)
}
