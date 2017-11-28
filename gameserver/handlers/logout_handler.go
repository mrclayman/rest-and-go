package handlers

import (
	"net/http"

	"github.com/mrclayman/rest-and-go/gameserver/core"
	"github.com/mrclayman/rest-and-go/gameserver/core/auth"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
	"github.com/mrclayman/rest-and-go/gameserver/serverlog"
)

type logoutRequest struct {
	ID    player.ID      `json:"id"`
	Token auth.AuthToken `json:"token"`
}

// LogoutHandler handles logout requests from players
// that are disconnecting from the server
type LogoutHandler struct {
	core *core.Core
}

// NewLogoutHandler returns a pointer to
// a newly created logout handler instance
func NewLogoutHandler(c *core.Core) *LogoutHandler {
	return &LogoutHandler{
		core: c,
	}
}

// ProcessRequest processes the incoming request and
// generates an appropriate response
func (h *LogoutHandler) ProcessRequest(resp http.ResponseWriter, req *http.Request) {
	serverlog.Logger.Println("Logout request received")

	if req.Method != "POST" {
		serverlog.Logger.Println("Wrong HTTP method used in logout request")
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Authenticate the player
	var logout logoutRequest
	if err := GetJSONFromRequest(req, &logout); err != nil {
		serverlog.Logger.Println("Failed to parse logout request body: " + err.Error())
		http.Error(resp, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	} else if !h.core.IsLoggedIn(logout.ID, logout.Token) {
		serverlog.Logger.Printf("Failed to authenticate token of player %v\n", logout.ID)
		http.Error(resp, "Failed to authenticate player's token", http.StatusUnauthorized)
		return
	}

	// Perform the logout procedure
	if err := h.core.QuitPlayer(logout.ID); err != nil {
		serverlog.Logger.Println("Could not log out the player: " + err.Error())
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	//WriteJSONToResponse(resp, nil)
	serverlog.Logger.Printf("Player %v logged out of the system\n", logout.ID)
}
