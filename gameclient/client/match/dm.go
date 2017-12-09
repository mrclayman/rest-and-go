package match

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

// DMPlayerRank aggregates information on a player's
// rank in a match
type DMPlayerRank struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// DMPlayerRanks defines a list of players' ranks
type DMPlayerRanks map[player.ID]DMPlayerRank

// DMMatch contains information on a
// DeathMatch type match
type DMMatch struct {
	Number Number        `json:"match_id"`
	Ranks  DMPlayerRanks `json:"ranks"`
}

// DMMatches defines a slice of DeathMatch-type matches
type DMMatches []*DMMatch
