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

// newPlayerDBTable creates a pre-filled table
// of players
func newPlayerDBTable() playerTable {
	return playerTable{
		"phreak":       &playerRecord{playerID: 1223145, password: "c0mm4nd0"},
		"fatal1ty":     &playerRecord{playerID: 8535253, password: "Quake4ever"}, // Already connected
		"How4rd":       &playerRecord{playerID: 5457676, password: "Noriko<3"},   // Already connected
		"Kr4zed":       &playerRecord{playerID: 9464779, password: "3328425913"}, // Already connected
		"Sir3n":        &playerRecord{playerID: 6735772, password: "Teh n00b!"},  // Already connected
		"Kamikaze":     &playerRecord{playerID: 1159734, password: "Get'em"},
		"Lone_Hunter":  &playerRecord{playerID: 4648464, password: "SniperFtw"},
		"ne0phyte":     &playerRecord{playerID: 6992112, password: "star4748"},
		"CrimsonDawn":  &playerRecord{playerID: 6433858, password: "necr0mancer"}, // Already connected
		"TheDamned1":   &playerRecord{playerID: 5747548, password: "f4llen1"},
		"SoulScorcher": &playerRecord{playerID: 1321878, password: "Burn'em_all!"},
		"LittlePony":   &playerRecord{playerID: 5723425, password: "pink"},
		"JigSaw":       &playerRecord{playerID: 4148994, password: "pure_3vil"},    // Already connected
		"Camping_Gaz":  &playerRecord{playerID: 9661327, password: "somepassw0rd"}, // Already connected
		"Dead3y3":      &playerRecord{playerID: 1412491, password: "5t4lk3r"},      // Already connected
		"Tweety":       &playerRecord{playerID: 8712722, password: "Silvester"},    // Already connected
		"Mikky":        &playerRecord{playerID: 4219691, password: "Come|Get|Some"},
	}
}

// playerNameIDTable maps player's id's to their nicknames
type playerNameIDTable map[PlayerID]string

// newPlayerNameIDTable returns a pre-initialized map
// between players' ID's and their nicknames
func newPlayerNameIDTable() playerNameIDTable {
	return playerNameIDTable{
		1223145: "phreak",
		8535253: "fatal1ty",
		5457676: "How4rd",
		9464779: "Kr4zed",
		6735772: "Sir3n",
		1159734: "Kamikaze",
		4648464: "Lone_Hunter",
		6992112: "ne0phyte",
		6433858: "CrimsonDawn",
		5747548: "TheDamned1",
		1321878: "SoulScorcher",
		5723425: "LittlePony",
		4148994: "JigSaw",
		9661327: "Camping_Gaz",
		1412491: "Dead3y3",
		8712722: "Tweety",
		4219691: "Mikky",
	}
}

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
func (db *Database) AuthenticatePlayer(login, password string) (PlayerID, bool) {
	if rec, ok := db.players[login]; !ok {
		return PlayerID(0), false
	} else if rec.password != password {
		return PlayerID(0), false
	} else {
		return rec.playerID, true
	}
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
