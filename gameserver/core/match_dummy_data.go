package core

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/match"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
)

func newMatchRegistry() match.Registry {
	dmMatchSlice := []*match.DMMatch{
		&match.DMMatch{
			Number: match.GenerateNumber(),
			Ranks: match.DMRanks{
				8535253: &match.DMRankRecord{
					Player: player.Player{ID: 8535253, Nick: "fatal1ty"},
					Kills:  22,
					Deaths: 6,
				},
				6433858: &match.DMRankRecord{
					Player: player.Player{ID: 6433858, Nick: "CrimsonDawn"},
					Kills:  14,
					Deaths: 9,
				},
				1412491: &match.DMRankRecord{
					Player: player.Player{ID: 1412491, Nick: "Dead3y3"},
					Kills:  5,
					Deaths: 11,
				},
			},
		},
	}

	lmsMatchSlice := []*match.LMSMatch{
		&match.LMSMatch{Number: match.GenerateNumber(),
			Ranks: match.LMSRanks{
				6735772: &match.LMSRankRecord{
					Player: player.Player{ID: 6735772, Nick: "Sir3n"},
					Kills:  5,
					Deaths: 7,
				},
				9661327: &match.LMSRankRecord{
					Player: player.Player{ID: 9661327, Nick: "Camping_Gaz"},
					Kills:  6,
					Deaths: 6,
				},
				8712722: &match.LMSRankRecord{
					Player: player.Player{ID: 8712722, Nick: "Tweety"},
					Kills:  2,
					Deaths: 4,
				},
				4148994: &match.LMSRankRecord{
					Player: player.Player{ID: 4148994, Nick: "JigSaw"},
					Kills:  4,
					Deaths: 0,
				},
			},
		},
	}

	duelMatchSlice := []*match.DuelMatch{
		&match.DuelMatch{Number: match.GenerateNumber(),
			Ranks: match.DuelRanks{
				5457676: &match.DuelRankRecord{
					Player: player.Player{ID: 5457676, Nick: "Howard"},
					Kills:  4,
					Deaths: 5,
				},
				9464779: &match.DuelRankRecord{
					Player: player.Player{ID: 9464779, Nick: "Kr4zed"},
					Kills:  5,
					Deaths: 4,
				},
			},
		},
	}

	dmMatches := make(match.DMMatches)
	for _, m := range dmMatchSlice {
		dmMatches[m.Number] = m
	}

	lmsMatches := make(match.LMSMatches)
	for _, m := range lmsMatchSlice {
		lmsMatches[m.Number] = m
	}

	duelMatches := make(match.DuelMatches)
	for _, m := range duelMatchSlice {
		duelMatches[m.Number] = m
	}

	return match.Registry{
		DM:   dmMatches,
		LMS:  lmsMatches,
		Duel: duelMatches,
	}
}
