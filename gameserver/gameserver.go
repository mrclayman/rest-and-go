package main

import (
	"crypto/tls"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/mrclayman/rest-and-go/gameserver/config"
	"github.com/mrclayman/rest-and-go/gameserver/core"
	"github.com/mrclayman/rest-and-go/gameserver/handlers"
	"github.com/mrclayman/rest-and-go/gameserver/serverlog"
	"gopkg.in/mgo.v2"
)

// port defines the server's listen port number
const port uint16 = 8000

const localDBURL string = "mongodb://localhost:27017/testgamedb"

// createDBDialInfo is used to create mgo.DialInfo object
// that allows communication over TLS/SSL-secured communication channel
func createDBDialInfo(c *config.DatabaseConfig) (*mgo.DialInfo, error) {
	var dialFn func(*mgo.ServerAddr) (net.Conn, error)
	if c.SSL {
		dialFn = func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		}
	}

	return &mgo.DialInfo{
		Addrs:          []string{c.Host},
		Timeout:        c.ConnectionTimeout,
		FailFast:       true,
		Database:       c.Database,
		ReplicaSetName: c.RSName,
		Username:       c.User,
		Password:       c.Password,
		DialServer:     dialFn,
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

	cmdArgs, err := config.ParseCmdLineArgs()
	if err != nil {
		log.Fatal("Failed to parse command line arguments: " + err.Error())
	}

	var cfgFile *config.Config
	if cfgFile, err = config.ParseCfgFile(cmdArgs.CfgFilePath); err != nil {
		log.Fatal("Failed to obtain valid configuration from file: " + err.Error())
	}

	if cfgFile.Log.LogMgo {
		mgo.SetLogger(serverlog.Logger)
	}

	di, err := createDBDialInfo(cfgFile.DB)
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
