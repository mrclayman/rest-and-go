package core

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/leaderboard"
	"github.com/mrclayman/rest-and-go/gameserver/core/match"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
)

func newMatchTable() match.Matches {
	matches := []*match.Match{
		{
			ID:   match.GenerateID(),
			Type: match.DeathMatch,
			Ranks: match.PlayerRanks{
				8535253: leaderboard.DMLeaderboardRecord{
					Player: player.Player{ID: 8535253, Nick: "fatal1ty"},
					Kills:  22,
					Deaths: 6,
				},
				6433858: leaderboard.DMLeaderboardRecord{
					Player: player.Player{ID: 6433858, Nick: "CrimsonDawn"},
					Kills:  14,
					Deaths: 9,
				},
				1412491: leaderboard.DMLeaderboardRecord{
					Player: player.Player{ID: 1412491, Nick: "Dead3y3"},
					Kills:  5,
					Deaths: 11,
				},
			},
		},
		{
			ID:   match.GenerateID(),
			Type: match.LastManStanding,
			Ranks: match.PlayerRanks{
				6735772: leaderboard.LMSLeaderboardRecord{
					Player: player.Player{ID: 6735772, Nick: "Sir3n"},
					Kills:  5,
					Deaths: 7,
				},
				9661327: leaderboard.LMSLeaderboardRecord{
					Player: player.Player{ID: 9661327, Nick: "Camping_Gaz"},
					Kills:  6,
					Deaths: 6,
				},
				8712722: leaderboard.LMSLeaderboardRecord{
					Player: player.Player{ID: 8712722, Nick: "Tweety"},
					Kills:  2,
					Deaths: 4,
				},
				4148994: leaderboard.LMSLeaderboardRecord{
					Player: player.Player{ID: 4148994, Nick: "JigSaw"},
					Kills:  4,
					Deaths: 0,
				},
			},
		},
		{
			ID:   match.GenerateID(),
			Type: match.Duel,
			Ranks: match.PlayerRanks{
				5457676: leaderboard.DuelLeaderboardRecord{
					Player: player.Player{ID: 5457676, Nick: "Howard"},
					Kills:  4,
					Deaths: 5,
				},
				9464779: leaderboard.DuelLeaderboardRecord{
					Player: player.Player{ID: 9464779, Nick: "Kr4zed"},
					Kills:  5,
					Deaths: 4,
				},
			},
		},
	}

	// Store the match instances in the map
	retval := make(match.Matches, len(matches))
	for _, match := range matches {
		retval[match.ID] = match
	}
	return retval
}
