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

func newMatchTable() Matches {

	matches := []*Match{
		{
			ID:   generateMatchID(),
			Type: DeathMatch,
			Ranks: PlayerMatchRanks{
				8535253: &PlayerMatchRank{
					Player: 8535253, // fatal1ty
					Kills:  22,
					Deaths: 6,
				},
				6433858: &PlayerMatchRank{
					Player: 6433858, // CrimsonDawn
					Kills:  14,
					Deaths: 9,
				},
				1412491: &PlayerMatchRank{
					Player: 1412491, // Dead3y3
					Kills:  5,
					Deaths: 11,
				},
			},
		},
		{
			ID:   generateMatchID(),
			Type: LastManStanding,
			Ranks: PlayerMatchRanks{
				6735772: &PlayerMatchRank{
					Player: 6735772, // Sir3n
					Kills:  5,
					Deaths: 7,
				},
				9661327: &PlayerMatchRank{
					Player: 9661327, // Camping_Gaz
					Kills:  6,
					Deaths: 6,
				},
				8712722: &PlayerMatchRank{
					Player: 8712722, // Tweety
					Kills:  2,
					Deaths: 4,
				},
				4148994: &PlayerMatchRank{
					Player: 4148994, // JigSaw
					Kills:  4,
					Deaths: 0,
				},
			},
		},
		{
			ID:   generateMatchID(),
			Type: Duel,
			Ranks: PlayerMatchRanks{
				5457676: {
					Player: 5457676, // Howard
					Kills:  4,
					Deaths: 5,
				},
				9464779: {
					Player: 9464779, // Kr4zed
					Kills:  5,
					Deaths: 4,
				},
			},
		},
	}

	// TODO Store the match instances in the map
	retval := make(Matches, len(matches))
	for _, match := range matches {
		retval[match.ID] = match
	}
	return retval
}

// PlayerMatchRank is a record of a player's
// performance in a match. Structurally, it is
// equivalent to a record in the leaderboards
type PlayerMatchRank LeaderboardRecord

// PlayerMatchRanks defines a map where each player's
// rank is identified by that player's id
type PlayerMatchRanks map[PlayerID]*PlayerMatchRank

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
