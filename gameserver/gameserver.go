package main

import (
	"crypto/tls"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/mrclayman/rest-and-go/gameserver/core"
	"github.com/mrclayman/rest-and-go/gameserver/handlers"
	"github.com/mrclayman/rest-and-go/gameserver/serverlog"
	"gopkg.in/mgo.v2"
)

// port defines the server's listen port number
const port uint16 = 8000

const localDBURL string = "mongodb://localhost:27017/testgamedb"

func createDBDialInfo(URL string) (*mgo.DialInfo, error) {
	var di *mgo.DialInfo
	var err error

	serverlog.Logger.Printf("Parsing database URL")
	if di, err = mgo.ParseURL(URL); err != nil {
		return nil, err
	}

	return di, nil
}

// createDBDialInfo is used to create mgo.DialInfo object
// that allows communication over TLS/SSL-secured communication channel
func createDBDialInfoAzure() (*mgo.DialInfo, error) {
	//const dbURL string = "mongodb://claytestgameserverdb:Y3IyfX0vFBlqNjyRw4VK8qRge6JxK2x80468XJppzC22KWAhsBcCQ8eHtOb2g6WEpkOHsM52TM2Vf2ObZtZgqA==@/?ssl=true&replicaSet=globaldb"

	return &mgo.DialInfo{
		Addrs:          []string{"claytestgameserverdb.documents.azure.com:10255"},
		Timeout:        time.Duration(30) * time.Second,
		FailFast:       true,
		Database:       "testgamedb",
		ReplicaSetName: "globaldb",
		Username:       "claytestgameserverdb",
		Password:       "Y3IyfX0vFBlqNjyRw4VK8qRge6JxK2x80468XJppzC22KWAhsBcCQ8eHtOb2g6WEpkOHsM52TM2Vf2ObZtZgqA==",
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}, nil
}

// Application structure binds together all the important
// parts of the server
type application struct {
	appCore    *core.Core
	dispatcher *handlers.MainDispatcher
}

func (a *application) Cleanup() {
	a.appCore.Cleanup()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//mgo.SetLogger(serverlog.Logger)

	di, err := createDBDialInfoAzure()
	//di, err := createDBDialInfo(localDBURL)
	if err != nil {
		serverlog.Logger.Fatal("Failed to create dial info: " + err.Error())
	}

	var c *core.Core
	c, err = core.NewCore(di)
	if err != nil {
		serverlog.Logger.Fatal("Failed to create server core object: " + err.Error())
	}

	app := &application{
		appCore:    c,
		dispatcher: handlers.NewMainDispatcher(c),
	}
	defer app.Cleanup()

	serverlog.Logger.Printf("Starting server on port %v", port)
	serverlog.Logger.Fatal(http.ListenAndServe(":"+strconv.Itoa(int(port)), app.dispatcher))
}
