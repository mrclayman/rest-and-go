package match

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

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
type CTFPlayerRanks map[player.ID]CTFPlayerRank

// CTFMatch contains information on
// a CTF type match
type CTFMatch struct {
	Number Number         `json:"match_id"`
	Ranks  CTFPlayerRanks `json:"ranks"`
}

// CTFMatches defines a slice of CTF-type matches
type CTFMatches []*CTFMatch
