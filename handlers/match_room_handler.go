package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mrclayman/rest_api_test/core"
)

// MatchRoomHandler handles WebSocket requests
// from a player that has previously joined
// a match
type MatchRoomHandler struct {
	core     *core.Core
	upgrader websocket.Upgrader
}

// NewMatchRoomHandler returns a pointer to a
// new instance of the match room request handler
func NewMatchRoomHandler(c *core.Core) *MatchRoomHandler {
	return &MatchRoomHandler{
		core:     c,
		upgrader: websocket.Upgrader{},
	}
}

// ProcessRequest processes the WebSocket request
// from the player and creates an appropriate response
func (h *MatchRoomHandler) ProcessRequest(resp http.ResponseWriter, req *http.Request) {
	// Obtain the connection object from the request
	conn, err := h.upgrader.Upgrade(resp, req, nil)
	if err != nil {
		http.Error(resp, "Failed to obtain connection from request: "+err.Error(), http.StatusInternalServerError)
	}

	defer conn.Close()

	// TODO Finish the implementation
}
