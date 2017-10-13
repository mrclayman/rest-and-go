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
	// TODO Add handlers for other client requests
	/*lboardHandler *LeaderboardHandler
	matchListHandler  *MatchListHandler */
}

// NewMainDispatcher creates a new instance of the
// main HTTP request dispatcher structure
func NewMainDispatcher(c *core.Core) *MainDispatcher {
	return &MainDispatcher{
		core:  c,
		login: LoginHandler{db: c.db},
	}
}

// ServeHTTP is the main hub of request processing and dispatching to subhandlers
func (dispatcher *MainDispatcher) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = util.SplitPath(req.URL.Path)
	switch head {
	case "login":
		dispatcher.login.ProcessRequest(resp, req)
		return
	default:
		http.Error(resp, "Resource not found", http.StatusNotFound)
	}
}
