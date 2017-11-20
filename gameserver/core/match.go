package core

import (
	"math/rand"
	"strconv"
)

// MatchID defines the type to represent match ID's
type MatchID uint64

// InvalidMatchID defines an invalid match ID
const InvalidMatchID MatchID = 0

// MatchIDToString converts a match id into
// a string
func MatchIDToString(mid MatchID) string {
	return strconv.FormatUint(uint64(mid), 10)
}

// Match structure represents an ongoing match
// and the participating players
type Match struct {
	ID    MatchID
	Type  GameType
	Ranks PlayerMatchRanks
}

// Matches defines a map of match instances
type Matches map[MatchID]*Match

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

// PlayerMatchRank is a record of a player's
// performance in a match. Since each game type
// has a slightly different structure for the
// rank/leaderboard record, the generic
// interface type is used here
type PlayerMatchRank interface{}

// PlayerMatchRanks defines a map where each player's
// rank is identified by that player's id
type PlayerMatchRanks map[PlayerID]PlayerMatchRank

// Add adds a player into the match. True is returned
// upon successful add, false is returned in case the
// player is already present in the match
func (m *Match) Add(player PlayerID) bool {
	if _, present := m.Ranks[player]; present {
		// Player already present
		return false
	}

	m.Ranks[player] = &PlayerMatchRank{Player: player}
	return true
}

// Remove removes a player with the given
// ID from the match. True is returned upon successful
// removal, false is returned in case the player
// is not present in the match
func (m *Match) Remove(id PlayerID) bool {
	if _, present := m.Ranks[id]; !present {
		// Player not participating in the match
		return false
	}

	delete(m.Ranks, id)
	return true
}
