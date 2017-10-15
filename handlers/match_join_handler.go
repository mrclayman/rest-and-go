package handlers

import (
	"net/http"

	"github.com/mrclayman/rest_api_test/core"
)

// joinRequest aggregates all the information
// required to fulfill a request to join an
// existing match or create a new one based
// on the desired game type
type joinRequest struct {
	PID   core.PlayerID  `json:"player_id"`
	Token core.AuthToken `json:"token"`
	MID   core.MatchID   `json:"match_id"`
	GType core.GameType  `json:"game_type"`
}

// MatchJoinHandler handles requests to join
// an existing match, or to create a new one
type MatchJoinHandler struct {
	core *core.Core
}

// ProcessRequest processes the user's request
// and generates an appropriate response
func (h *MatchJoinHandler) ProcessRequest(resp http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Joining a match requires a reference to a particular
	// resource, check the URL to see if it is really there
	var action string
	action, req.URL.Path = SplitPath(req.URL.Path)
	if action != "join" {
		http.Error(resp, "Resource not found", http.StatusNotFound)
		return
	}

	// Parse and process the request
	var join joinRequest
	err := GetJSONFromRequest(req, &join)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	} else if !h.core.IsLoggedIn(join.PID, join.Token) {
		http.Error(resp, "Could not authenticate player's token", http.StatusUnauthorized)
		return
	} else if join.MID, err = h.core.JoinMatch(join.MID, join.PID, join.GType); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate the necessary WebSock token
	// and dispatch the response to the player
	wsToken := core.GenerateWebSockToken()
	output := map[string]interface{}{
		"match_id": join.MID,
		"ws_token": wsToken,
	}
	WriteJSONToResponse(resp, output)
}
