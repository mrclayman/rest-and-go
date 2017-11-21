package client

import "strconv"

// Player structure defines the basic credentials of
// the player connected to the server
type Player struct {
	ID   int    `json:"player_id"`
	Name string `json:"player_name"`
}

// DMPlayerRank aggregates information on a player's
// rank in a match
type DMPlayerRank struct {
	Player Player `json:"player"`
	Kills  uint   `json:"kills"`
	Deaths uint   `json:"deaths"`
}

// DMPlayerRanks defines a list of players' ranks
type DMPlayerRanks []DMPlayerRank

// CTFPlayerRank defines the structure of a player's
// rank in the CTF game type
type CTFPlayerRank struct {
	Player   Player `json:"player"`
	Kills    uint   `json:"kills"`
	Deaths   uint   `json:"deaths"`
	Captures uint   `json:"captures"`
}

// CTFPlayerRanks is a slice of CTF game type
// player rank objects
type CTFPlayerRanks []CTFPlayerRank

// LMSPlayerRank defines the structure of a player's
// rank in the LMS game type
type LMSPlayerRank struct {
	Player Player `json:"player"`
	Kills  uint   `json:"kills"`
	Deaths uint   `json:"deaths"`
	Wins   uint   `json:"wins"`
}

// LMSPlayerRanks is a type for a slice of
// LMS game type player rank objects
type LMSPlayerRanks []LMSPlayerRank

// DuelPlayerRank defines the structure of
// a player's rank in the Duel game type
type DuelPlayerRank LMSPlayerRank

// DuelPlayerRanks defines the type for a slice
// of Duel game type player rank objects
type DuelPlayerRanks []DuelPlayerRank

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
