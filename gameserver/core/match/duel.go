package match

import "github.com/mrclayman/rest-and-go/gameserver/core/player"

// DuelRankRecord defines the structure for
// the DeathMatch-type match
type DuelRankRecord struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// DuelRanks defines the type of storage
// of players' ranks for a DeathMatch-type match
type DuelRanks map[player.ID]*DuelRankRecord

// DuelMatch defines the structure for
// a DeathMatch-type match
type DuelMatch struct {
	Number Number    `json:"match_id"`
	Ranks  DuelRanks `json:"ranks"`
}

// Add adds a player into the match. If the player
// is already present in the match, false is returned.
// Otherwise true is returned.
func (m *DuelMatch) Add(p player.Player) bool {
	if _, present := m.Ranks[p.ID]; present {
		// Player already present
		return false
	}

	m.Ranks[p.ID] = &DuelRankRecord{Player: p}
	return true
}

// Remove removes a player with the given
// ID from the match. True is returned upon successful
// removal, false is returned in case the player
// is not present in the match
func (m *DuelMatch) Remove(ID player.ID) bool {
	if _, present := m.Ranks[ID]; !present {
		// Player not participating in the match
		return false
	}

	delete(m.Ranks, ID)
	return true
}

// newDuel creates a new Duel-type match
func newDuel(pl player.List) *DuelMatch {
	m := &DuelMatch{
		Number: GenerateNumber(),
		Ranks:  make(DuelRanks),
	}

	for _, p := range pl {
		m.Add(p)
	}

	return m
}

// DuelMatches defines the number-keyed
// storage for active DeathMatch matches
type DuelMatches map[Number]*DuelMatch
