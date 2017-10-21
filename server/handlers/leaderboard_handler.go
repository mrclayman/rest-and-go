package handlers

import (
	"log"
	"net/http"

	"github.com/mrclayman/rest-and-go/server/core"
)

// LeaderboardHandler handles requests for
// leaderboard listings
type LeaderboardHandler struct {
	core *core.Core
}

// NewLeaderboardHandler returns a pointer
// to a new leaderboard handler instance
func NewLeaderboardHandler(c *core.Core) *LeaderboardHandler {
	return &LeaderboardHandler{core: c}
}

// ProcessRequest process the client's request and prepares
// an appropriate response
func (h *LeaderboardHandler) ProcessRequest(resp http.ResponseWriter, req *http.Request) {
	log.Println("Received leaderboard list request")

	if req.Method != "GET" {
		log.Println("Wrong HTTP method used in leaderboard list request")
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// GET parameters need to be parsed on-request
	req.ParseForm()

	// Authenticate user
	playerID, token, err := GetPlayerDataFromGET(req)
	if err != nil {
		log.Printf("Could not obtain player's credentials from GET: %v", err.Error())
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	} else if !h.core.IsLoggedIn(playerID, token) {
		log.Printf("Failed to authenticate token of player %v", playerID)
		http.Error(resp, "Could not authenticate player's token", http.StatusUnauthorized)
		return
	}

	// The game type is specified as a subpath, so separate it
	// from the rest of the path and analyze it
	var strGameType string
	strGameType, req.URL.Path = SplitPath(req.URL.Path)
	gameType, ok := core.IsValidGameType(strGameType)
	if !ok {
		log.Printf("Invalid game type specified in leaderboard request of player %v", playerID)
		http.Error(resp, "Invalid game type specified", http.StatusBadRequest)
		return
	}

	// Obtain the desired leaderboard information and
	// serialize it into a JSON structure for dispatch
	leaderboard, err := h.core.GetLeaderboardForJSON(gameType)
	if err != nil {
		log.Printf("Could not obtain leaderboard for mode '%v' for player %v: %v", strGameType, playerID, err.Error())
		http.Error(resp, "Failed to obtain leaderboard: "+err.Error(), http.StatusInternalServerError)
	}

	WriteJSONToResponse(resp, leaderboard)
	log.Printf("Response to leaderboard request of player %v dispatched", playerID)
}
