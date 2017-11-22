package match

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
)

// DMRankRecord defines the structure for
// the DeathMatch-type match
type DMRankRecord struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// DMRanks defines the type of storage
// of players' ranks for a DeathMatch-type match
type DMRanks map[player.ID]*DMRankRecord

// DMMatch defines the structure for
// a DeathMatch-type match
type DMMatch struct {
	ID    Number  `json:"match_id"`
	Ranks DMRanks `json:"ranks"`
}

// Add adds a player into the match. If the player
// is already present in the match, false is returned.
// Otherwise true is returned.
func (m *DMMatch) Add(p player.Player) bool {
	if _, present := m.Ranks[p.ID]; present {
		// Player already present
		return false
	}

	m.Ranks[p.ID] = &DMRankRecord{Player: p}
	return true
}

// Remove removes a player with the given
// ID from the match. True is returned upon successful
// removal, false is returned in case the player
// is not present in the match
func (m *DMMatch) Remove(ID player.ID) bool {
	if _, present := m.Ranks[ID]; !present {
		// Player not participating in the match
		return false
	}

	delete(m.Ranks, ID)
	return true
}

// DMMatches defines the number-keyed
// storage for active DeathMatch matches
type DMMatches map[Number]*DMMatch
