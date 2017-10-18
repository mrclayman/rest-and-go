package core

// newConnectedPlayerTable creates a table of
// connected players
func newConnectedPlayerTable() Players {
	// The list below corresponds to the complete
	// list of players already participating in matches
	// (see function newMatchTable() in match_dummy_data.go)
	players := []*Player{
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

	retval := make(Players, len(players))
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
		8535253: GenerateAuthenticationToken(),
		5457676: GenerateAuthenticationToken(),
		9464779: GenerateAuthenticationToken(),
		6735772: GenerateAuthenticationToken(),
		6433858: GenerateAuthenticationToken(),
		4148994: GenerateAuthenticationToken(),
		9661327: GenerateAuthenticationToken(),
		1412491: GenerateAuthenticationToken(),
		8712722: GenerateAuthenticationToken(),
	}
}

// newPlayerInMatchTokenTable returns a table of
// WebSocket tokens
func newPlayerInMatchTokenTable() PlayerWSTokens {
	return PlayerWSTokens{
		8535253: GenerateWebSocketToken(),
		5457676: GenerateWebSocketToken(),
		9464779: GenerateWebSocketToken(),
		6735772: GenerateWebSocketToken(),
		6433858: GenerateWebSocketToken(),
		4148994: GenerateWebSocketToken(),
		9661327: GenerateWebSocketToken(),
		1412491: GenerateWebSocketToken(),
		8712722: GenerateWebSocketToken(),
	}
}
