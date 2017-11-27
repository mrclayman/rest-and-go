package database

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// New creates a pre-filled player database
func New(di *mgo.DialInfo) (*Database, error) {
	log.Printf("Connecting to database")
	var s *mgo.Session
	var err error
	if s, err = mgo.DialWithInfo(di); err != nil {
		return nil, err
	}

	log.Printf("Database connected")
	retval := &Database{
		session:               s,
		dbName:                di.Database,
		leaderboardCollPrefix: "leaderboard_",
		playersCollName:       "players",
	}

	return retval, nil
}

// isError checks the map in the argument
// for the presence of either "$err" or "errmsg",
// both of which are values added by MongoDB
// in case there was a problem while executing
// the query
func isError(r map[string]interface{}) string {

	if v, ok := r["$err"]; ok {
		return v.(string)
	} else if v, ok := r["errmsg"]; ok {
		return v.(string)
	}

	return ""
}
