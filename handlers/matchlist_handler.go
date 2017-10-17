package handlers

import (
	"net/http"

	"github.com/mrclayman/rest-and-go/core"
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
	if req.Method != "GET" {
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// GET parameters need to be parsed on-request
	req.ParseForm()

	// Authenticate user
	playerID, token, err := GetPlayerDataFromGET(req)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	} else if !h.core.IsLoggedIn(playerID, token) {
		http.Error(resp, "Could not authenticate player's token", http.StatusUnauthorized)
		return
	}

	matchlist, err := h.core.GetMatchlistForJSON()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
	}
	WriteJSONToResponse(resp, matchlist)
}
