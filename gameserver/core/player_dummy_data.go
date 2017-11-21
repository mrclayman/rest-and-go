package core

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/auth"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
)

// newConnectedPlayerTable creates a table of
// connected players
func newConnectedPlayerTable() player.Map {
	// The list below corresponds to the complete
	// list of players already participating in matches
	// (see function newMatchTable() in match_dummy_data.go)
	players := []player.Player{
		{
			ID:   8535253,
			Nick: "fatal1ty",
		},
		{
			ID:   5457676,
			Nick: "How4rd",
		},
		{
			ID:   9464779,
			Nick: "Kr4zed",
		},
		{
			ID:   6735772,
			Nick: "Sir3n",
		},
		{
			ID:   6433858,
			Nick: "CrimsonDawn",
		},
		{
			ID:   4148994,
			Nick: "JigSaw",
		},
		{
			ID:   9661327,
			Nick: "Camping_Gaz",
		},
		{
			ID:   1412491,
			Nick: "Dead3y3",
		},
		{
			ID:   8712722,
			Nick: "Tweety",
		},
	}

	retval := make(player.Map, len(players))
	for _, player := range players {
		retval[player.ID] = player
	}

	return retval
}

// newConnectedPlayerTokenTable creates a table
// with player id's and their respective authentication
// tokens
func newConnectedPlayerTokenTable() PlayerAuthTokens {
	return PlayerAuthTokens{
		8535253: auth.GenerateAuthenticationToken(),
		5457676: auth.GenerateAuthenticationToken(),
		9464779: auth.GenerateAuthenticationToken(),
		6735772: auth.GenerateAuthenticationToken(),
		6433858: auth.GenerateAuthenticationToken(),
		4148994: auth.GenerateAuthenticationToken(),
		9661327: auth.GenerateAuthenticationToken(),
		1412491: auth.GenerateAuthenticationToken(),
		8712722: auth.GenerateAuthenticationToken(),
	}
}

// newPlayerInMatchTokenTable returns a table of
// WebSocket tokens
func newPlayerInMatchTokenTable() PlayerWSTokens {
	return PlayerWSTokens{
		8535253: auth.GenerateWebSocketToken(),
		5457676: auth.GenerateWebSocketToken(),
		9464779: auth.GenerateWebSocketToken(),
		6735772: auth.GenerateWebSocketToken(),
		6433858: auth.GenerateWebSocketToken(),
		4148994: auth.GenerateWebSocketToken(),
		9661327: auth.GenerateWebSocketToken(),
		1412491: auth.GenerateWebSocketToken(),
		8712722: auth.GenerateWebSocketToken(),
	}
}
