package match

import "github.com/mrclayman/rest-and-go/gameserver/core/player"

// DMPlayerRank aggregates information on a player's
// rank in a match
type DMPlayerRank struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// DMPlayerRanks defines a list of players' ranks
type DMPlayerRanks []DMPlayerRank

// CTFPlayerRank defines the structure of a player's
// rank in the CTF game type
type CTFPlayerRank struct {
	Player   player.Player `json:"player"`
	Kills    uint          `json:"kills"`
	Deaths   uint          `json:"deaths"`
	Captures uint          `json:"captures"`
}

// CTFPlayerRanks is a slice of CTF game type
// player rank objects
type CTFPlayerRanks []CTFPlayerRank

// LMSPlayerRank defines the structure of a player's
// rank in the LMS game type
type LMSPlayerRank struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
	Wins   uint          `json:"wins"`
}

// LMSPlayerRanks is a type for a slice of
// LMS game type player rank objects
type LMSPlayerRanks []LMSPlayerRank

// DuelPlayerRank defines the structure of
// a player's rank in the Duel game type
type DuelPlayerRank LMSPlayerRank

// DuelPlayerRanks defines the type for a slice
// of Duel game type player rank objects
type DuelPlayerRanks LMSPlayerRanks
