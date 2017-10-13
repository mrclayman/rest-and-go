package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/mrclayman/rest_api_test/core"
	"github.com/mrclayman/rest_api_test/handlers"
)

// Application structure binds together all the important
// parts of the server
type Application struct {
	appCore    *core.Core
	dispatcher *handlers.MainDispatcher
}

func main() {
	rand.Seed(time.Now().UnixNano())
	core := core.NewCore()
	app := &Application{
		appCore:    core,
		dispatcher: handlers.NewMainDispatcher(core),
	}
	http.ListenAndServe(":8000", app.dispatcher)
}
