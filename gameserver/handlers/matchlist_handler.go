package handlers

import (
	"net/http"

	"github.com/mrclayman/rest-and-go/gameserver/core"
	"github.com/mrclayman/rest-and-go/gameserver/serverlog"
)

// MatchlistHandler handles requests regarding listing
// of ongoing matches for the player to choose from
type MatchlistHandler struct {
	core *core.Core
}

// NewMatchlistHandler returns a pointer to a new
// matchlist handler instance
func NewMatchlistHandler(c *core.Core) *MatchlistHandler {
	return &MatchlistHandler{core: c}
}

// ProcessRequest processes the incoming request and creates
// an appropriate response
func (h *MatchlistHandler) ProcessRequest(resp http.ResponseWriter, req *http.Request) {
	serverlog.Logger.Println("Received match list request")

	if req.Method != "GET" {
		serverlog.Logger.Println("Wrong HTTP method used in matchlist request")
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// GET parameters need to be parsed on-request
	req.ParseForm()

	// Authenticate user
	playerID, token, err := GetPlayerDataFromGET(req)
	if err != nil {
		serverlog.Logger.Printf("Could not obtain player's credentials from GET: %v", err.Error())
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	} else if !h.core.IsLoggedIn(playerID, token) {
		serverlog.Logger.Printf("Failed to authenticate token of player %v", playerID)
		http.Error(resp, "Could not authenticate player's token", http.StatusUnauthorized)
		return
	}

	serverlog.Logger.Println("Retrieving match list")
	matchlist := h.core.GetMatchlistForJSON()
	if err != nil {
		serverlog.Logger.Printf("Could not obtain match list for player %v: %v", playerID, err.Error())
		http.Error(resp, "Failed to retrieve match list: "+err.Error(), http.StatusInternalServerError)
	}

	WriteJSONToResponse(resp, matchlist)
	serverlog.Logger.Printf("Response to matchlist request of player %v dispatched", playerID)
}
