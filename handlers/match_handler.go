package handlers

import (
	"net/http"

	"github.com/mrclayman/rest-and-go/core"
)

// MatchHandler structure represents a handler
// processing REST/WS requests regarding match
// joining and match event processing
type MatchHandler struct {
	core *core.Core
	join *MatchJoinHandler
	room *MatchRoomHandler
}

// NewMatchHandler returns a pointer to a new
// instance of the match request handler
func NewMatchHandler(c *core.Core) *MatchHandler {
	return &MatchHandler{
		core: c,
		join: NewMatchJoinHandler(c),
		room: NewMatchRoomHandler(c),
	}
}

// ProcessRequest inspects the incoming request and dispatches it
// further to individual handlers based on the type of the request
func (h *MatchHandler) ProcessRequest(resp http.ResponseWriter, req *http.Request) {
	var action string
	switch action, req.URL.Path = SplitPath(req.URL.Path); action {
	case "join":
		h.join.ProcessRequest(resp, req)
	case "room":
		h.room.ProcessRequest(resp, req)
	default:
		http.Error(resp, "Resource '"+action+"' not found", http.StatusNotFound)
	}
}
