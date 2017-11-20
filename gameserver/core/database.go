package core

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// playerRecord is an internal representation
// of a player record in the database
type playerRecord struct {
	playerID PlayerID
	password string
}

// playerTable is a exactly that, a table of
// player records. Keys are the players' login
// names.
type playerTable map[string]*playerRecord

// playerNameIDTable maps player's id's to their nicknames
type playerNameIDTable map[PlayerID]string

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

// newDatabase creates a pre-filled player database
func newDatabase(URL string) (*Database, error) {

	var retval *Database

	if di, err := mgo.ParseURL(URL); err != nil {
		return nil, err
	} else if s, err := mgo.DialWithInfo(di); err != nil {
		return nil, err
	} else {
		retval = &Database{
			session:               s,
			dbName:                di.Database,
			leaderboardCollPrefix: "leaderboard_",
			playersCollName:       "players",
		}
	}

	return retval, nil
}

// AuthenticatePlayer looks up the player's login name
// in the table of players and compares the password
// provided in the argument with the one stored in the
// table. If there is a match, true and the player's
// internal ID is returned. In the case the player
// name provided is not in the database, or the passwords
// do not match, PlayerTypeID(0) and false is returned
func (db *Database) AuthenticatePlayer(login, password string) (PlayerID, error) {

	c := db.session.DB(db.dbName).C(db.playersCollName)
	q := c.Find(bson.M{"nick": login, "password": password})
	r := make(map[string]interface{})

	if err := q.One(r); err != nil {
		return PlayerID(0), DatabaseError{err.Error()}
	} else if errMsg := isError(r); errMsg != "" {
		return PlayerID(0), DatabaseError{errMsg}
	} else if len(r) == 0 {
		return PlayerID(0), InvalidArgumentError{"Wrong player nickname or password"}
	}

	return r["id"].(PlayerID), nil
}

// GetLeaderboard returns the leaderboard
// associated with the given type
func (db *Database) GetLeaderboard(id GameType) (interface{}, error) {
	lbName := db.leaderboardCollPrefix + string(id)
	c := db.session.DB(db.dbName).C(lbName)

	var r interface{}
	var sortingCrit []string
	switch id {
	case DeathMatch:
		r = DMLeaderboard{}
		sortingCrit = dmLeaderboardSortingCriterion
	case CaptureTheFlag:
		r = CTFLeaderboard{}
		sortingCrit = ctfLeaderboardSortingCriterion
	case LastManStanding:
		r = LMSLeaderboard{}
		sortingCrit = lmsLeaderboardSortingCriterion
	case Duel:
		r = DuelLeaderboard{}
		sortingCrit = duelLeaderboardSortingCriterion
	}

	q := c.Find(bson.M{}).Sort(sortingCrit...)

	if err := q.All(r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetPlayerNick queries the database for a nickname
// of a player identified by 'id'
func (db *Database) GetPlayerNick(id PlayerID) (string, error) {
	c := db.session.DB(db.dbName).C(db.playersCollName)
	q := c.Find(bson.M{"playerid": id}).Select(bson.M{"playername": 1})

	var r map[string]interface{}
	if err := q.One(r); err != nil {
		return "", err
	}

	return r["playername"].(string), nil
}

// Close closes the internal session
func (db *Database) Close() {
	db.session.Close()
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
