package core

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

// Database is a mock-up representation of a
// backend database system containing records
// about players.
type Database struct {
	players      playerTable
	playerIDs    playerNameIDTable
	leaderboards GameLeaderboards
}

// newDatabase creates a pre-filled player database
func newDatabase() *Database {
	retval := &Database{
		players:      newPlayerDBTable(),
		playerIDs:    newPlayerNameIDTable(),
		leaderboards: newLeaderboardTables(),
	}
	return retval
}

// AuthenticatePlayer looks up the player's login name
// in the table of players and compares the password
// provided in the argument with the one stored in the
// table. If there is a match, true and the player's
// internal ID is returned. In the case the player
// name provided is not in the database, or the passwords
// do not match, PlayerTypeID(0) and false is returned
func (db *Database) AuthenticatePlayer(login, password string) (PlayerID, error) {
	var rec *playerRecord
	var ok bool
	if rec, ok = db.players[login]; !ok {
		return PlayerID(0), InvalidArgumentError{"Player nick '" + login + "' not in the database"}
	} else if rec.password != password {
		return PlayerID(0), InvalidArgumentError{"Passwords do not match"}
	}

	return rec.playerID, nil
}

// GetLeaderboard returns the leaderboard
// associated with the given type
func (db *Database) GetLeaderboard(id GameType) (*Leaderboard, bool) {
	leaderboard, ok := db.leaderboards[id]
	return leaderboard, ok
}

// GetPlayerNick queries the database for a nickname
// of a player identified by 'id'
func (db *Database) GetPlayerNick(id PlayerID) (string, bool) {
	nick, ok := db.playerIDs[id]
	return nick, ok
}
