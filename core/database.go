package core

// playerRecord is an internal representation
// of a player record in the database
type playerRecord struct {
	playerID PlayerIDType
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

// Database is a mock-up representation of
// backend database system containing records
// about players. The type has been intentionally
// left exported in case someone wants an empty
// database and not a pre-filled one.
type Database struct {
	players      playerTable
	leaderboards MatchTypeLeaderboardType
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
func (db *Database) AuthenticatePlayer(login, password string) (PlayerIDType, bool) {
	if rec, ok := db.players[login]; !ok {
		return PlayerIDType(0), false
	} else if rec.password != password {
		return PlayerIDType(0), false
	} else {
		return rec.playerID, true
	}
}

func (db *Database) GetLeaderboardForType(id MatchTypeIDType) (*LeaderboardType, bool) {
	return nil, db.leaderboards[id]
}
