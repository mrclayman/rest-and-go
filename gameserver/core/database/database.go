package database

import (
	"gopkg.in/mgo.v2"
)

const (
	// errNotFound is issued by MongoDB in case a document
	// could not be found in the database collection
	errNotFound string = "not found"
)

// Database holds the session object to
// the database system and also a few
// auxiliary pieces of information (the
// name of the database, the prefix for the
// leaderboard collections,...)
type Database struct {
	session               *mgo.Session
	dbName                string
	leaderboardCollPrefix string
	playersCollName       string
}

// Close closes the internal session
func (db *Database) Close() {
	db.session.Close()
}
