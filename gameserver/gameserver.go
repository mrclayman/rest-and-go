package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/mrclayman/rest-and-go/gameserver/core"
	"github.com/mrclayman/rest-and-go/gameserver/handlers"
)

// port defines the server's listen port number
const port uint16 = 8000

// Application structure binds together all the important
// parts of the server
type application struct {
	appCore    *core.Core
	dispatcher *handlers.MainDispatcher
}

func main() {
	rand.Seed(time.Now().UnixNano())
	core := core.NewCore()
	app := &application{
		appCore:    core,
		dispatcher: handlers.NewMainDispatcher(core),
	}

	log.Printf("Starting server on port %v", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(int(port)), app.dispatcher))
}
