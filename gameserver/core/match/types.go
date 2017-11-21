package match

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
)

// ID defines the type to represent match ID's
type ID uint64

// InvalidID defines an invalid match ID
const InvalidID ID = 0

// PlayerRank is a record of a player's
// performance in a match. Since each game type
// has a slightly different structure for the
// rank/leaderboard record, the generic
// interface type is used here
type PlayerRank interface{}

// PlayerRanks defines a map where each player's
// rank is identified by that player's id
type PlayerRanks map[player.ID]PlayerRank
