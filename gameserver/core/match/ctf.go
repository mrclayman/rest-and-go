package match

import "github.com/mrclayman/rest-and-go/gameserver/core/player"

// CTFRankRecord defines the structure for
// the DeathMatch-type match
type CTFRankRecord struct {
	Player   player.Player `json:"player"`
	Kills    uint          `json:"kills"`
	Deaths   uint          `json:"deaths"`
	Captures uint          `json:"captures"`
}

// CTFRanks defines the type of storage
// of players' ranks for a DeathMatch-type match
type CTFRanks map[player.ID]*CTFRankRecord

// CTFMatch defines the structure for
// a DeathMatch-type match
type CTFMatch struct {
	ID    ID       `json:"match"`
	Ranks CTFRanks `json:"ranks"`
}

// Add adds a player into the match. If the player
// is already present in the match, false is returned.
// Otherwise true is returned.
func (m *CTFMatch) Add(p player.Player) bool {
	if _, present := m.Ranks[p.ID]; present {
		// Player already present
		return false
	}

	m.Ranks[p.ID] = &CTFRankRecord{Player: p}
	return true
}

// Remove removes a player with the given
// ID from the match. True is returned upon successful
// removal, false is returned in case the player
// is not present in the match
func (m *CTFMatch) Remove(ID player.ID) bool {
	if _, present := m.Ranks[ID]; !present {
		// Player not participating in the match
		return false
	}

	delete(m.Ranks, ID)
	return true
}

// CTFMatches defines the ID-keyed
// storage for active DeathMatch matches
type CTFMatches map[ID]*CTFMatch

// NewCTF creates a new match and populates
// it with the given set of players
func NewCTF(pl player.List) *CTFMatch {
	m := &CTFMatch{
		ID: ID{
			Number: GenerateNumber(),
			Type:   CaptureTheFlag,
		},
	}

	for _, p := range pl {
		m.Add(p)
	}

	return m
}
