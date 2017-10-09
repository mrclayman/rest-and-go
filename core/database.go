package core

// playerRecord is an internal representation
// of a player record in the database
type playerRecord struct {
	password string
}

// playerTable is a exactly that, a table of
// player records. Keys are the players' login
// names.
type playerTable map[string]playerRecord

// NewPlayerTable creates a pre-filled table
// of players
func NewPlayerTable() playerTable {
	return playerTable{
		"phreak":   playerRecord{password: "c0mm4nd0"},
		"fatal1ty": playerRecord{password: "Quake4ever"},
		"How4rd":   playerRecord{password: "Noriko<3"},
		"Sir3n":    playerRecord{password: "Teh n00b!"},
	}
}

// Database is a mock-up representation of
// backend database system containing records
// about players. The type has been intentionally
// left exported in case someone wants an empty
// database and not a pre-filled one.
type Database struct {
	players playerTable
}

// NewDatabase creates a pre-filled player database
func NewDatabase() *Database {
	retval := new(Database)
	retval.players = NewPlayerTable()
	return retval
}

// AuthenticatePlayer looks up the player's login name
// in the table of players and compares the password
// provided in the argument with the one stored in the
// table. If there is a match, true is returned. In the
// case the player name provided is not in the database,
// or the passwords do not match, false is returned
func (db *Database) AuthenticatePlayer(playerLogin, playerPassword string) bool {
	if playerRec, ok := db.players[playerLogin]; !ok {
		return false
	} else if playerRec.password != playerPassword {
		return false
	}

	return true
}
