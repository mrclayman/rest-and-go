package match

import "github.com/mrclayman/rest-and-go/gameserver/core/player"

// Match structure represents an ongoing match
// and the participating players
type Match struct {
	ID    ID
	Type  GameType
	Ranks PlayerRanks
}

// Matches defines a map of match instances
type Matches map[ID]*Match

// Add adds a player into the match. If the player
// is already present in the match, false and nil error
// are returned. In case an error occurs, false and the
// error object are returned. Otherwise true and nil
// error are returned.
func (m *Match) Add(p player.Player) (bool, error) {
	if _, present := m.Ranks[p.ID]; present {
		// Player already present
		return false, nil
	}

	var r interface{}
	var err error
	if r, err = createNewLeaderboardRecord(m.Type, p); err != nil {
		return false, err
	}
	m.Ranks[p.ID] = r

	return true, nil
}

// Remove removes a player with the given
// ID from the match. True is returned upon successful
// removal, false is returned in case the player
// is not present in the match
func (m *Match) Remove(ID player.ID) bool {
	if _, present := m.Ranks[ID]; !present {
		// Player not participating in the match
		return false
	}

	delete(m.Ranks, ID)
	return true
}
