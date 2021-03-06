package database

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
	"github.com/mrclayman/rest-and-go/gameserver/core/servererrors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AuthenticatePlayer looks up the player's login name
// in the table of players and compares the password
// provided in the argument with the one stored in the
// table. If there is a match, true and the player's
// internal ID is returned. In the case the player
// name provided is not in the database, or the passwords
// do not match, PlayerTypeID(0) and false is returned
func (db *Database) AuthenticatePlayer(login, password string) (player.ID, error) {

	//serverlog.Logger.Println("Verifying player login:", login, ",", password)
	c := db.session.DB(db.dbName).C(db.playersCollName)
	q := c.Find(bson.M{"nick": login, "password": password}).Select(bson.M{"id": 1})
	r := make(map[string]interface{})
	var refreshed, done bool

	for !done {
		if err := q.One(r); err != nil {
			if err == mgo.ErrNotFound {
				return player.ID(0), servererrors.InvalidArgumentError{Message: "Wrong player nickname or password"}
			} else if !refreshed {
				db.session.Refresh()
				refreshed = true
				continue
			} else {
				return player.ID(0), servererrors.DatabaseError{Message: err.Error()}
			}
		} else if errMsg := isError(r); errMsg != "" {
			return player.ID(0), servererrors.DatabaseError{Message: errMsg}
		}
		done = true
	}
	return player.ID(r["id"].(int)), nil
}

// GetPlayerNick queries the database for a nickname
// of a player identified by 'id'
func (db *Database) GetPlayerNick(id player.ID) (string, error) {
	c := db.session.DB(db.dbName).C(db.playersCollName)
	q := c.Find(bson.M{"playerid": id}).Select(bson.M{"playername": 1})
	var r map[string]interface{}

	var refreshed, done bool
	for !done {
		if err := q.One(r); err != nil {
			if err == mgo.ErrNotFound {
				return "", servererrors.InvalidArgumentError{Message: "Unknown player ID"}
			} else if !refreshed {
				db.session.Refresh()
				refreshed = true
				continue
			} else {
				return "", err
			}
		}
		done = true
	}
	return r["playername"].(string), nil
}
