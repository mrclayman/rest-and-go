package handlers

import (
	"net/http"

	"github.com/mrclayman/rest-and-go/core"
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

// NewMatchJoinHandler returns a pointer to a new
// instance of the match join request handler
func NewMatchJoinHandler(c *core.Core) *MatchJoinHandler {
	return &MatchJoinHandler{core: c}
}

// ProcessRequest processes the user's request
// and generates an appropriate response
func (h *MatchJoinHandler) ProcessRequest(resp http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
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
	}

	// Generate WebSocket token and add the player
	// to the match, or create a new match if necessary
	wsToken := core.GenerateWebSocketToken()
	if join.MID, err = h.core.JoinMatch(join.MID, join.PID, wsToken, join.GType); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	output := map[string]interface{}{
		"match_id": join.MID,
		"ws_token": wsToken,
	}
	WriteJSONToResponse(resp, output)
}
