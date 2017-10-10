package core

import (
	"math/rand"
)

// MatchID defines the type to represent match ID's
type MatchID uint64

// GameType is a type designating the type of a game/match
type GameType string

// DeathMatch indicates the match is of type "deathmatch"
const DeathMatch GameType = "dm"

// CaptureTheFlag indicates the match is of type "capture the flag"
const CaptureTheFlag GameType = "ctf"

// LastManStanding indicates the match is of type "last man standing"
const LastManStanding GameType = "lms"

// Match structure represents an ongoing match
// and the participating players
type Match struct {
	ID    MatchID
	Type  GameType
	Ranks PlayerMatchRanksType
}

// Matches defines a map of match instances
type Matches map[MatchID]*Match

// PlayerMatchRankType defines the type that
// represents a player in a match, including
// their number of kills and deaths
type PlayerMatchRankType struct {
	PlayerID   PlayerIDType
	KillCount  int
	DeathCount int
}

// PlayerMatchRanksType defines a map where each player's
// rank is identified by that player's id
type PlayerMatchRanksType map[PlayerIDType]*PlayerMatchRankType

// AddPlayer adds a player into the match. True is returned
// upon successful add, false is returned in case the
// player is already present in the match
func (match *Match) AddPlayer(playerID PlayerIDType) bool {
	if _, present := match.PlayerRanks[playerID]; present {
		// Player already present
		return false
	}

	match.PlayerRanks[playerID] = &PlayerMatchRankType{PlayerID: playerID}
	return true
}

// RemovePlayer removes a player with the given
// ID from the match. True is returned upon successful
// removal, false is returned in case the player
// is not present in the match
func (match *Match) RemovePlayer(playerID PlayerIDType) bool {
	if _, present := match.PlayerRanks[playerID]; !present {
		// Player not participating in the match
		return false
	}

	delete(match.PlayerRanks, playerID)
	return true
}

// Generates a random match ID
func generateMatchID() MatchID {
	return MatchID(rand.Uint64())
}

// NewMatchNoPlayers creates a new match with no players
func NewMatchNoPlayers(matchType GameTypeID) *MatchType {
	matchID := generateMatchID()
	return &MatchType{Match: matchID, MatchTypeID: matchType}
}

// NewMatchWithPlayers creates a new match and populates
// it with the given set of players
func NewMatchWithPlayers(matchType GameTypeID, playerIDs PlayerIDsType) *MatchType {
	matchID := generateMatchID()
	playerRanks := make(PlayerMatchRanksType, len(playerIDs))
	for _, playerID := range playerIDs {
		playerRanks[playerID] = &PlayerMatchRankType{PlayerID: playerID}
	}
	return &MatchType{MatchID: matchID, MatchTypeID: matchType, PlayerRanks: playerRanks}
}
