package handlers

import (
	"net/http"

	"github.com/mrclayman/rest_api_test/core"
	"github.com/mrclayman/rest_api_test/util"
)

// MainDispatcher aggregates all request handlers
// and distributes to them the incoming requests
type MainDispatcher struct {
	core         *core.Core
	loginHandler LoginHandler
	// TODO Add handlers for other client requests
	/*lboardHandler *LeaderboardHandler
	matchListHandler  *MatchListHandler */

}

// ServeHTTP is the main hub of request processing and dispatching to subhandlers
func (dispatcher *MainDispatcher) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	var head string
	head, request.URL.Path = util.SplitPath(request.URL.Path)
	switch head {
	case "login":
		dispatcher.loginHandler.ProcessRequest(response, request)
		return
	default:
		http.Error(response, "Resource not found", http.StatusNotFound)
	}
}
