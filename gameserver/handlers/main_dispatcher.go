package handlers

import (
	"net/http"

	"github.com/mrclayman/rest-and-go/gameserver/core"
)

// MainDispatcher aggregates all request handlers
// and distributes to them the incoming requests
type MainDispatcher struct {
	core   *core.Core
	login  *LoginHandler
	mlist  *MatchlistHandler
	lboard *LeaderboardHandler
	match  *MatchHandler
	logout *LogoutHandler
	// TODO Add join handler
}

// NewMainDispatcher creates a new instance of the
// main HTTP request dispatcher structure
func NewMainDispatcher(c *core.Core) *MainDispatcher {
	return &MainDispatcher{
		core:   c,
		login:  NewLoginHandler(c),
		mlist:  NewMatchlistHandler(c),
		lboard: NewLeaderboardHandler(c),
		match:  NewMatchHandler(c),
		logout: NewLogoutHandler(c),
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
	case "leaderboards":
		dispatcher.lboard.ProcessRequest(resp, req)
	case "match":
		dispatcher.match.ProcessRequest(resp, req)
	case "logout":
		dispatcher.logout.ProcessRequest(resp, req)
	default:
		http.Error(resp, "Resource not found", http.StatusNotFound)
	}
}
