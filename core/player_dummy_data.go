package core

// newConnectedPlayerTable creates a table of
// connected players
func newConnectedPlayerTable() Players {
	// The list below corresponds to the complete
	// list of players already participating in matches
	// (see function newMatchTable() in match_dummy_data.go)
	players := []*Player{
		{
			ID:    8535253,
			Nick:  "fatal1ty",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    5457676,
			Nick:  "How4rd",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    9464779,
			Nick:  "Kr4zed",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    6735772,
			Nick:  "Sir3n",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    6433858,
			Nick:  "CrimsonDawn",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    4148994,
			Nick:  "JigSaw",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    9661327,
			Nick:  "Camping_Gaz",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    1412491,
			Nick:  "Dead3y3",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    8712722,
			Nick:  "Tweety",
			Token: GenerateAuthenticationToken(),
		},
	}

	retval := make(Players, len(players))
	for _, player := range players {
		retval[player.ID] = player
	}

	return retval
}
