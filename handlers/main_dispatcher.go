package handlers

import (
	"net/http"

	"github.com/mrclayman/rest_api_test/core"
)

// MainDispatcher aggregates all request handlers
// and distributes to them the incoming requests
type MainDispatcher struct {
	core  *core.Core
	login LoginHandler
	mlist MatchlistHandler
	// TODO Add handlers for other client requests
	/*lboardHandler *LeaderboardHandler
	matchListHandler  *MatchListHandler */
}

// NewMainDispatcher creates a new instance of the
// main HTTP request dispatcher structure
func NewMainDispatcher(c *core.Core) *MainDispatcher {
	return &MainDispatcher{
		core:  c,
		login: LoginHandler{core: c},
		mlist: MatchlistHandler{core: c},
	}
}

// ServeHTTP is the main hub of request processing and dispatching to subhandlers
func (dispatcher *MainDispatcher) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = SplitPath(req.URL.Path)
	switch head {
	case "login":
		dispatcher.login.ProcessRequest(resp, req)
	case "matches":
		dispatcher.mlist.ProcessRequest(resp, req)
	default:
		http.Error(resp, "Resource not found", http.StatusNotFound)
	}
}
