package core

import (
	"math/rand"
)

// MatchID defines the type to represent match ID's
type MatchID uint64

// Match structure represents an ongoing match
// and the participating players
type Match struct {
	ID    MatchID
	Type  GameType
	Ranks PlayerMatchRanks
}

// Matches defines a map of match instances
type Matches map[MatchID]*Match

// PlayerMatchRank defines the type that
// represents a player in a match, including
// their number of kills and deaths
type PlayerMatchRank struct {
	Player     PlayerID
	KillCount  int
	DeathCount int
}

// PlayerMatchRanks defines a map where each player's
// rank is identified by that player's id
type PlayerMatchRanks map[PlayerID]*PlayerMatchRank

// AddPlayer adds a player into the match. True is returned
// upon successful add, false is returned in case the
// player is already present in the match
func (m *Match) AddPlayer(player PlayerID) bool {
	if _, present := m.Ranks[player]; present {
		// Player already present
		return false
	}

	m.Ranks[player] = &PlayerMatchRank{Player: player}
	return true
}

// RemovePlayer removes a player with the given
// ID from the match. True is returned upon successful
// removal, false is returned in case the player
// is not present in the match
func (m *Match) RemovePlayer(id PlayerID) bool {
	if _, present := m.Ranks[id]; !present {
		// Player not participating in the match
		return false
	}

	delete(m.Ranks, id)
	return true
}

// Generates a random match ID
func generateMatchID() MatchID {
	return MatchID(rand.Uint64())
}

// NewMatchNoPlayers creates a new match with no players
func NewMatchNoPlayers(gt GameType) *Match {
	matchID := generateMatchID()
	return &Match{ID: matchID, Type: gt}
}

// NewMatchWithPlayers creates a new match and populates
// it with the given set of players
func NewMatchWithPlayers(gt GameType, ids PlayerIDs) *Match {
	matchID := generateMatchID()
	ranks := make(PlayerMatchRanks, len(ids))
	for _, id := range ids {
		ranks[id] = &PlayerMatchRank{Player: id}
	}
	return &Match{ID: matchID, Type: gt, Ranks: ranks}
}
