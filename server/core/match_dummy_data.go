package core

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

	// Store the match instances in the map
	retval := make(Matches, len(matches))
	for _, match := range matches {
		retval[match.ID] = match
	}
	return retval
}
