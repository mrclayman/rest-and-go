package match

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

// DuelPlayerRank defines the structure of
// a player's rank in the Duel game type
type DuelPlayerRank struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// DuelPlayerRanks defines the type for a slice
// of Duel game type player rank objects
type DuelPlayerRanks map[player.ID]DuelPlayerRank

// DuelMatch contains information on
// a Duel type match
type DuelMatch struct {
	Number Number          `json:"match_id"`
	Ranks  DuelPlayerRanks `json:"ranks"`
}

// DuelMatches defines a slice of Duel-type matches
type DuelMatches []*DuelMatch
