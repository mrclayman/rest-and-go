package match

import "github.com/mrclayman/rest-and-go/gameserver/core/player"

// LMSRankRecord defines the structure for
// the DeathMatch-type match
type LMSRankRecord struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// LMSRanks defines the type of storage
// of players' ranks for a DeathMatch-type match
type LMSRanks map[player.ID]*LMSRankRecord

// LMSMatch defines the structure for
// a DeathMatch-type match
type LMSMatch struct {
	Number Number   `json:"match_id"`
	Ranks  LMSRanks `json:"ranks"`
}

// Add adds a player into the match. If the player
// is already present in the match, false is returned.
// Otherwise true is returned.
func (m *LMSMatch) Add(p player.Player) bool {
	if _, present := m.Ranks[p.ID]; present {
		// Player already present
		return false
	}

	m.Ranks[p.ID] = &LMSRankRecord{Player: p}
	return true
}

// Remove removes a player with the given
// ID from the match. True is returned upon successful
// removal, false is returned in case the player
// is not present in the match
func (m *LMSMatch) Remove(ID player.ID) bool {
	if _, present := m.Ranks[ID]; !present {
		// Player not participating in the match
		return false
	}

	delete(m.Ranks, ID)
	return true
}

// newLMS creates a new LMS-type match
func newLMS(pl player.List) *LMSMatch {
	m := &LMSMatch{
		Number: GenerateNumber(),
		Ranks:  make(LMSRanks),
	}

	for _, p := range pl {
		m.Add(p)
	}

	return m
}

// LMSMatches defines the number-keyed
// storage for active DeathMatch matches
type LMSMatches map[Number]*LMSMatch
