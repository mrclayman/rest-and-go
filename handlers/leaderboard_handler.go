package handlers

import (
	"net/http"

	"github.com/mrclayman/rest_api_test/core"
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
	if req.Method != "GET" {
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// GET parameters need to be parsed on-request
	req.ParseForm()

	// Authenticate user
	playerID, token, err := GetPlayerDataFromGetArgs(req)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	} else if !h.core.IsLoggedIn(playerID, token) {
		http.Error(resp, "Could not authenticate player's token", http.StatusUnauthorized)
		return
	}

	// The game type is specified as a subpath, so separate it
	// from the rest of the path and analyze it
	var strGameType string
	strGameType, req.URL.Path = SplitPath(req.URL.Path)
	gameType, ok := core.IsValidGameType(strGameType)
	if !ok {
		http.Error(resp, "Invalid game type specified", http.StatusBadRequest)
		return
	}

	// Obtain the desired leaderboard information and
	// serialize it into a JSON structure for dispatch
	leaderboard, err := h.core.GetLeaderboardForJSON(gameType)
	if err != nil {
		http.Error(resp, "Failed to obtain leaderboard", http.StatusInternalServerError)
	}
	WriteJSONToResponse(resp, leaderboard)
}
