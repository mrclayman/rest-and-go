package handlers

import (
	"net/http"

	"github.com/mrclayman/rest_api_test/core"
	"github.com/mrclayman/rest_api_test/util"
)

type MainDispatcher struct {
	core         *core.Core
	loginHandler LoginHandler
	/*lboardHandler *LeaderboardHandler
	matchListHandler  *MatchListHandler */

}

// ProcessRequest is the main hub of request processing and dispatching to subhandlers
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
