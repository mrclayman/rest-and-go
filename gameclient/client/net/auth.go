package net

import (
	"strconv"

	"github.com/mrclayman/rest-and-go/gameclient/client/match"
)

// PlayerSession contains information on a player's
// authentication data returned by the server
type PlayerSession struct {
	ID    int    `json:"player_id"`
	Token string `json:"token"`
}

// ToGet converts the contents of the
// PlayerSession instance into GET request
// arguments
func (ps PlayerSession) ToGet() string {
	return "id=" + strconv.Itoa(ps.ID) + "&token=" + ps.Token
}

// MatchSession contains information on the
// match joined and a WebSocket token used
// to communicate with the server while the
// player is participating in the match
type MatchSession struct {
	ID    match.ID `json:"match"`
	Token string   `json:"ws_token"`
}
