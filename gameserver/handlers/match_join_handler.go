package handlers

import (
	"net/http"

	"github.com/mrclayman/rest-and-go/gameserver/core"
	"github.com/mrclayman/rest-and-go/gameserver/core/auth"
	"github.com/mrclayman/rest-and-go/gameserver/core/match"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
	"github.com/mrclayman/rest-and-go/gameserver/serverlog"
)

// joinRequest aggregates all the information
// required to fulfill a request to join an
// existing match or create a new one based
// on the desired game type
type joinRequest struct {
	PlayerID player.ID      `json:"player_id"`
	Token    auth.AuthToken `json:"token"`
	Match    match.ID       `json:"match"`
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
	serverlog.Logger.Println("Received match join request")

	if req.Method != "POST" {
		serverlog.Logger.Println("Wrong HTTP method used in match join request")
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse and process the request
	var join joinRequest
	err := GetJSONFromRequest(req, &join)
	if err != nil {
		serverlog.Logger.Println("Failed to parse matchlist request body: " + err.Error())
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	} else if !h.core.IsLoggedIn(join.PlayerID, join.Token) {
		serverlog.Logger.Println("Could not authenticate player's token")
		http.Error(resp, "Could not authenticate player's token", http.StatusUnauthorized)
		return
	}

	// Generate WebSocket token and add the player
	// to the match, or create a new match if necessary
	wsToken := auth.GenerateWebSocketToken()
	if join.Match.Number, err = h.core.JoinMatch(join.Match, join.PlayerID, wsToken); err != nil {
		serverlog.Logger.Println(err.Error())
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	serverlog.Logger.Printf("Registered player %v in match %v", join.PlayerID, join.Match.Number)

	output := map[string]interface{}{
		"match":    join.Match,
		"ws_token": wsToken,
	}

	WriteJSONToResponse(resp, output)
	serverlog.Logger.Printf("Response to match join request of player %v dispatched", join.PlayerID)
}
