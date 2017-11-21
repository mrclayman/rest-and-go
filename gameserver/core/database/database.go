package database

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/errors"
	"github.com/mrclayman/rest-and-go/gameserver/core/leaderboard"
	"github.com/mrclayman/rest-and-go/gameserver/core/match"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// AuthenticatePlayer looks up the player's login name
// in the table of players and compares the password
// provided in the argument with the one stored in the
// table. If there is a match, true and the player's
// internal ID is returned. In the case the player
// name provided is not in the database, or the passwords
// do not match, PlayerTypeID(0) and false is returned
func (db *Database) AuthenticatePlayer(login, password string) (player.ID, error) {

	c := db.session.DB(db.dbName).C(db.playersCollName)
	q := c.Find(bson.M{"nick": login, "password": password})
	r := make(map[string]interface{})

	if err := q.One(r); err != nil {
		if err.Error() == errNotFound {
			return player.ID(0), errors.InvalidArgumentError{Message: "Wrong player nickname or password"}
		}
		return player.ID(0), errors.DatabaseError{Message: err.Error()}
	} else if errMsg := isError(r); errMsg != "" {
		return player.ID(0), errors.DatabaseError{Message: errMsg}
	}

	return player.ID(r["id"].(int)), nil
}

// GetLeaderboard returns the leaderboard
// associated with the given type
func (db *Database) GetLeaderboard(ID match.GameType) (interface{}, error) {
	lbName := db.leaderboardCollPrefix + match.GameTypeToString(ID)
	c := db.session.DB(db.dbName).C(lbName)

	var r interface{}
	var sortingCrit []string
	switch ID {
	case match.DeathMatch:
		r = leaderboard.DMLeaderboard{}
		sortingCrit = dmLeaderboardSortingCriterion
	case match.CaptureTheFlag:
		r = leaderboard.CTFLeaderboard{}
		sortingCrit = ctfLeaderboardSortingCriterion
	case match.LastManStanding:
		r = leaderboard.LMSLeaderboard{}
		sortingCrit = lmsLeaderboardSortingCriterion
	case match.Duel:
		r = leaderboard.DuelLeaderboard{}
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
func (db *Database) GetPlayerNick(id player.ID) (string, error) {
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
