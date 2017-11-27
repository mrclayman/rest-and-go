package main

import (
	"errors"
	"fmt"

	"github.com/mrclayman/rest-and-go/database/data"
	"gopkg.in/mgo.v2"
)

const (
	playerCollection string = "players"

	dmLbCollection string = "leaderboard_dm"

	ctfLbCollection string = "leaderboard_ctf"

	lmsLbCollection string = "leaderboard_lms"

	duelLbCollection string = "leaderboard_duel"
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

func createDuelLeaderboard(b *mgo.Bulk, lbData ...data.DuelLeaderboardRecord) {
	for _, item := range lbData {
		b.Insert(item)
	}
}

// createCollection creates a collection
// with the given name in the database
func createCollection(session *mgo.Session, dbName, colName string, cData interface{}) error {
	c := session.DB(dbName).C(colName)
	b := c.Bulk()

	switch colName {
	case playerCollection:
		createPlayerCollection(b, cData.(data.PlayerRecords)...)
	case dmLbCollection:
		createDMLeaderboard(b, cData.(data.DMLeaderboardRecords)...)
	case ctfLbCollection:
		createCTFLeaderboard(b, cData.(data.CTFLeaderboardRecords)...)
	case lmsLbCollection:
		createLMSLeaderboard(b, cData.(data.LMSLeaderboardRecords)...)
	case duelLbCollection:
		createDuelLeaderboard(b, cData.(data.DuelLeaderboardRecords)...)
	default:
		return errors.New("Unhandled collection name " + colName)
	}

	if _, err := b.Run(); err != nil {
		return errors.New("An error occurred while creating collection '" + colName + "': " + err.Error())
	}
	fmt.Println("All OK")
	return nil
}

// createPlayerIndex creates an index for the "nick" and "password"
// items in records of the "players" collection
// NOTE: Azure indexes all elements of documents automatically
// so the index is not really necessary when importing data to Azure.
func createPlayerIndex(s *mgo.Session, dbName string) error {
	c := s.DB(dbName).C(playerCollection)
	i := mgo.Index{
		Key:    []string{"nick", "password"},
		Unique: true, // Unique index is not supported by Azure
	}

	return c.EnsureIndex(i)
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

	fmt.Println("Adding Duel leaderboard data")
	if err := createCollection(session, ci.Database, "leaderboard_duel", data.DuelLeaderboard); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Creating index on the player collection")
	if err := createPlayerIndex(session, ci.Database); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Data inserted, exiting")
}
