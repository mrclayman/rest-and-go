package handlers

import (
	"net/http"

	"github.com/mrclayman/rest-and-go/gameserver/core"
	"github.com/mrclayman/rest-and-go/gameserver/core/auth"
	"github.com/mrclayman/rest-and-go/gameserver/serverlog"
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
	serverlog.Logger.Println("Received login request")

	if req.Method != "POST" {
		serverlog.Logger.Println("Wrong HTTP method used in login request")
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var credent loginRequest
	err := GetJSONFromRequest(req, &credent)
	if err != nil {
		serverlog.Logger.Printf("Failed to parse JSON from login request body: %v", err.Error())
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate
	id, err := h.core.AuthenticatePlayer(credent.Name, credent.Password)
	if err != nil {
		serverlog.Logger.Println(err.Error())
		http.Error(resp, err.Error(), http.StatusForbidden)
		return
	}

	// Success, generate a token for the player and add them
	serverlog.Logger.Printf("Registering player '%v' (id %v) in the system", credent.Name, id)
	token := auth.GenerateAuthenticationToken()
	h.core.AddConnected(id, credent.Name, token)

	// Dispatch the response to the player
	respData := map[string]interface{}{"player_id": id, "token": token}
	// TODO Should I check that the response has been serialized correctly?
	WriteJSONToResponse(resp, respData)
	serverlog.Logger.Printf("Response to login request of player '%v' (id %v) dispatched", credent.Name, id)
}
