package database

import (
	"gopkg.in/mgo.v2"
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
