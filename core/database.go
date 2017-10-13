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
type playerTable map[string]playerRecord

// newPlayerTable creates a pre-filled table
// of players
func newPlayerTable() *playerTable {
	return &playerTable{
		"phreak":   playerRecord{playerID: 1223145, password: "c0mm4nd0"},
		"fatal1ty": playerRecord{playerID: 8535253, password: "Quake4ever"},
		"How4rd":   playerRecord{playerID: 5457676, password: "Noriko<3"},
		"Sir3n":    playerRecord{playerID: 6735772, password: "Teh n00b!"},
	}
}

// Database is a mock-up representation of a
// backend database system containing records
// about players.
type Database struct {
	players      playerTable
	leaderboards GameLeaderboards
}

// NewDatabase creates a pre-filled player database
func NewDatabase() *Database {
	retval := &Database{players: *newPlayerTable()}
	return retval
}

// AuthenticatePlayer looks up the player's login name
// in the table of players and compares the password
// provided in the argument with the one stored in the
// table. If there is a match, true and the player's
// internal ID is returned. In the case the player
// name provided is not in the database, or the passwords
// do not match, PlayerTypeID(0) and false is returned
func (db *Database) AuthenticatePlayer(login, password string) (PlayerID, bool) {
	if rec, ok := db.players[login]; !ok {
		return PlayerID(0), false
	} else if rec.password != password {
		return PlayerID(0), false
	} else {
		return rec.playerID, true
	}
}

// GetLeaderboardForType returns the leaderboard
// associated with the given type
func (db *Database) GetLeaderboardForType(id GameType) (*Leaderboard, bool) {
	leaderboard, ok := db.leaderboards[id]
	return leaderboard, ok
}
