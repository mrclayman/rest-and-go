package main

import (
	"net/http"

	"github.com/mrclayman/rest_api_test/core"
	"github.com/mrclayman/rest_api_test/handlers"
)

// Application structure binds together all the important
// parts of the server
type Application struct {
	core       *core.Core
	dispatcher *handlers.MainDispatcher
}

func main() {
	app := &Application{
		core:       new(core.Core),
		dispatcher: new(handlers.MainDispatcher{core}),
	}
	http.ListenAndServe(":8000", app.dispatcher)
}
