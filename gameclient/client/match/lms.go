package match

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

// LMSPlayerRank defines the structure of a player's
// rank in the LMS game type
type LMSPlayerRank struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// LMSPlayerRanks is a type for a slice of
// LMS game type player rank objects
type LMSPlayerRanks map[player.ID]LMSPlayerRank

// LMSMatch contains information on
// a LMS type match
type LMSMatch struct {
	Number Number         `json:"match_id"`
	Ranks  LMSPlayerRanks `json:"ranks"`
}

// LMSMatches defines a slice of LMS-type matches
type LMSMatches []*LMSMatch
