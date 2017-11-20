package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mrclayman/rest-and-go/database/data"
	"gopkg.in/mgo.v2"
)

// getServerURL asks the user for the
// MongoDB server's URL and checks that
// no error has occurred while parsing
// the user's input
func getServerURL() (string, error) {
	var serverURL string

	fmt.Print("Enter the URL of the MongoDB server: ")
	if _, err := fmt.Scanf("%v", &serverURL); err != nil {
		return "", err
	}
	return serverURL, nil
}

// getConnectionInfo asks the user for
// an address of the MongoDB server and
// creates a DialInfo structure to that
// holds the name of the database to use,
// among other things
func getConnectionInfo() (*mgo.DialInfo, error) {
	var URL string
	var err error
	if URL, err = getServerURL(); err != nil {
		return nil, err
	}

	var ci *mgo.DialInfo
	ci, err = mgo.ParseURL(URL)
	if err != nil {
		return nil, err
	}
	return ci, nil
}

func createPlayerCollection(b *mgo.Bulk, players ...data.PlayerRecord) {
	for _, item := range players {
		b.Insert(item)
	}
}

func createDMLeaderboard(b *mgo.Bulk, lbData ...data.DMLeaderboardRecord) {
	for _, item := range lbData {
		b.Insert(item)
	}
}

func createCTFLeaderboard(b *mgo.Bulk, lbData ...data.CTFLeaderboardRecord) {
	for _, item := range lbData {
		b.Insert(item)
	}
}

func createLMSLeaderboard(b *mgo.Bulk, lbData ...data.LMSLeaderboardRecord) {
	for _, item := range lbData {
		b.Insert(item)
	}
}

// createCollection creates a collection
// with the given name in the database
func createCollection(session *mgo.Session, dbName, colName string, cData interface{}) error {
	c := session.DB(dbName).C(colName)
	b := c.Bulk()

	switch concreteData := cData.(type) {
	case data.PlayerRecords:
		createPlayerCollection(b, concreteData...)
	case data.DMLeaderboardRecords:
		createDMLeaderboard(b, concreteData...)
	case data.CTFLeaderboardRecords:
		createCTFLeaderboard(b, concreteData...)
	case data.LMSLeaderboardRecords:
		createLMSLeaderboard(b, concreteData...)
	default:
		return errors.New("Unhandled type " + reflect.TypeOf(concreteData).Name())
	}

	if _, err := b.Run(); err != nil {
		return errors.New("An error occurred while creating collection '" + colName + "': " + err.Error())
	}
	fmt.Println("All OK")
	return nil
}

func main() {
	var err error
	var ci *mgo.DialInfo

	if ci, err = getConnectionInfo(); err != nil {
		fmt.Println("Failed to parse database name from user input:", err.Error())
		return
	}

	var session *mgo.Session
	if session, err = mgo.DialWithInfo(ci); err != nil {
		fmt.Println("Failed to create database session:", err.Error())
		return
	}
	defer session.Close()

	fmt.Println("Adding player data")
	if err := createCollection(session, ci.Database, "players", data.Players); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Adding DM leaderboard data")
	if err := createCollection(session, ci.Database, "leaderboard_dm", data.DMLeaderboard); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Adding CTF leaderboard data")
	if err := createCollection(session, ci.Database, "leaderboard_ctf", data.CTFLeaderboard); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Adding LMS leaderboard data")
	if err := createCollection(session, ci.Database, "leaderboard_lms", data.LMSLeaderboard); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Data inserted, exiting")
}
