package client

////////////////// Constants /////////////////
const (
	// Azure server address
	//serverAddress string = "claytestgameserver.azurewebsites.net"

	serverAddress = "localhost:8000"
)

////////////////// Types ////////////////



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
