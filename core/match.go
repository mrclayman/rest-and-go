package core

import (
	"math/rand"
)

// MatchIDType defines the type to represent match ID's
type MatchIDType uint64

// MatchTypeIDType is a type designating the type
// of a match
type MatchTypeIDType string

// DeathMatchTypeID indicates the match is of type "deathmatch"
const DeathMatchTypeID MatchTypeIDType = "deathmatch"

// CaptureTheFlagTypeID indicates the match is of type "capture the flag"
const CaptureTheFlagTypeID MatchTypeIDType = "capture_the_flag"

// LastManStandingTypeID indicates the match is of type "last man standing"
const LastManStandingTypeID MatchTypeIDType = "last_man_standing"

// MatchType structure represents an ongoing match
// and the participating players
type MatchType struct {
	MatchID     MatchIDType
	MatchTypeID MatchTypeIDType
	Players     PlayersType
}

// MatchesType defines a map of match instances
type MatchesType map[MatchIDType]*MatchType

// AddPlayer adds a player into the match. True is returned
// upon successful add, false is returned in case the
// player is already present in the match
func (match *MatchType) AddPlayer(player *PlayerType) bool {
	if _, present := match.Players[player.PlayerID]; present {
		// Player already present
		return false
	}

	match.Players[player.PlayerID] = player
	return true
}

// RemovePlayer removes a player with the given
// ID from the match. True is returned upon successful
// removal, false is returned in case the player
// is not present in the match
func (match *MatchType) RemovePlayer(playerID PlayerIDType) bool {
	if _, present := match.Players[playerID]; !present {
		// Player not participating in the match
		return false
	}

	delete(match.Players, playerID)
	return true
}

// Generates a random match ID
func generateMatchID() MatchIDType {
	return MatchIDType(rand.Uint64())
}

// NewMatchNoPlayers creates a new match with no players
func NewMatchNoPlayers(matchType MatchTypeIDType) *MatchType {
	matchID := generateMatchID()
	return &MatchType{MatchID: matchID, MatchTypeID: matchType}
}

// NewMatchWithPlayers creates a new match and populates
// it with the given set of players
func NewMatchWithPlayers(matchType MatchTypeIDType, players PlayersType) *MatchType {
	matchID := generateMatchID()
	return &MatchType{MatchID: matchID, MatchTypeID: matchType, Players: players}
}
