package handlers

import (
	"log"
	"net/http"

	"github.com/mrclayman/rest-and-go/gameserver/core"
)

// joinRequest aggregates all the information
// required to fulfill a request to join an
// existing match or create a new one based
// on the desired game type
type joinRequest struct {
	PlayerID core.PlayerID  `json:"player_id"`
	Token    core.AuthToken `json:"token"`
	MatchID  core.MatchID   `json:"match_id"`
	GType    core.GameType  `json:"game_type"`
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
	log.Println("Received match join request")

	if req.Method != "POST" {
		log.Println("Wrong HTTP method used in match join request")
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse and process the request
	var join joinRequest
	err := GetJSONFromRequest(req, &join)
	if err != nil {
		log.Println("Failed to parse matchlist request body: " + err.Error())
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	} else if !h.core.IsLoggedIn(join.PlayerID, join.Token) {
		log.Println("Could not authenticate player's token")
		http.Error(resp, "Could not authenticate player's token", http.StatusUnauthorized)
		return
	}

	// Generate WebSocket token and add the player
	// to the match, or create a new match if necessary
	wsToken := core.GenerateWebSocketToken()
	if join.MatchID, err = h.core.JoinMatch(join.MatchID, join.PlayerID, wsToken, join.GType); err != nil {
		log.Println(err.Error())
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Registered player %v in match %v", join.PlayerID, join.MatchID)

	output := map[string]interface{}{
		"match_id": join.MatchID,
		"ws_token": wsToken,
	}

	WriteJSONToResponse(resp, output)
	log.Printf("Response to match join request of player %v dispatched", join.PlayerID)
}