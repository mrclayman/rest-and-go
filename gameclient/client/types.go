package client

import "strconv"

////////////////// Constants /////////////////
const (
	serverAddress string = "claytestgameserver.azurewebsites.net"
)

////////////////// Types ////////////////

// PlayerRank aggregates information on a player's
// rank in a match
type PlayerRank struct {
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name"`
	Kills      uint   `json:"kills"`
	Deaths     uint   `json:"deaths"`
}

// PlayerRanks defines a list of players' ranks
type PlayerRanks []PlayerRank

// Match contains information on a match
type Match struct {
	ID    uint64      `json:"match_id"`
	Type  string      `json:"match_type"`
	Ranks PlayerRanks `json:"ranks"`
}

// Matchlist defines a slice of match instances
type Matchlist []Match

// PlayerAuthData contains information on a player's
// authentication data returned by the server
type PlayerAuthData struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

// ToGet converts the contents of the
// PlayerAuthData instance into GET request
// arguments
func (data PlayerAuthData) ToGet() string {
	return "id=" + strconv.Itoa(data.ID) + "&token=" + data.Token
}

// PlayerLogin contains login information
// for a player
type PlayerLogin struct {
	Nick     string `json:"name"`
	Password string `json:"password"`
}

// PlayerLogins is a slice of player
// login data structures
var PlayerLogins = []PlayerLogin{
	{"phreak", "c0mm4nd0"},
	{"Kamikaze", "Get'em"},
	{"Lone_Hunter", "SniperFtw"},
	{"ne0phyte", "star4748"},
	{"TheDamned1", "f4llen1"},
	{"SoulScorcher", "Burn'em_all!"},
	{"LittlePony", "pink"},
	{"Mikky", "Come|Get|Some"},
}

// LeaderboardRecord contains information
// on record in a leaderboard
type LeaderboardRecord struct {
	ID     int    `json:"player_id"`
	Nick   string `json:"player_name"`
	Kills  uint   `json:"kills"`
	Deaths uint   `json:"deaths"`
}

// Leaderboard is a slice of leaderboard records
type Leaderboard []LeaderboardRecord

// MatchSessionData contains information on the
// match joined and a WebSocket token used
// to communicate with the server while the
// player is participating in the match
type MatchSessionData struct {
	MatchID uint64 `json:"match_id"`
	WSToken string `json:"ws_token"`
}
