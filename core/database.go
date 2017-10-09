package core

// playerRecordType is an internal representation
// of a player record in the database
type playerRecordType struct {
	PlayerID PlayerIDType
	password string
}

// PlayerTableType is a exactly that, a table of
// player records. Keys are the players' login
// names.
type PlayerTableType map[string]playerRecordType

// NewPlayerTable creates a pre-filled table
// of players
func NewPlayerTable() *PlayerTableType {
	return &PlayerTableType{
		"phreak":   playerRecordType{PlayerID: 1223145, password: "c0mm4nd0"},
		"fatal1ty": playerRecordType{PlayerID: 8535253, password: "Quake4ever"},
		"How4rd":   playerRecordType{PlayerID: 5457676, password: "Noriko<3"},
		"Sir3n":    playerRecordType{PlayerID: 6735772, password: "Teh n00b!"},
	}
}

// Database is a mock-up representation of
// backend database system containing records
// about players. The type has been intentionally
// left exported in case someone wants an empty
// database and not a pre-filled one.
type Database struct {
	players PlayerTableType
}

// NewDatabase creates a pre-filled player database
func NewDatabase() *Database {
	retval := &Database{players: *NewPlayerTable()}
	return retval
}

// AuthenticatePlayer looks up the player's login name
// in the table of players and compares the password
// provided in the argument with the one stored in the
// table. If there is a match, true and the player's
// internal ID= is returned. In the case the player
// name provided is not in the database, or the passwords
// do not match, PlayerTypeID(0) and false is returned
func (db *Database) AuthenticatePlayer(playerLogin, playerPassword string) (PlayerIDType, bool) {
	if playerRec, ok := db.players[playerLogin]; !ok {
		return PlayerIDType(0), false
	} else if playerRec.password != playerPassword {
		return PlayerIDType(0), false
	} else {
		return playerRec.PlayerID, true
	}
}
