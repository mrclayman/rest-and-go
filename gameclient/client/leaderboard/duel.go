package leaderboard

import "github.com/mrclayman/rest-and-go/gameclient/client/player"

// DuelLeaderboardRecord contains information
// on record in a Duel-type leaderboard
type DuelLeaderboardRecord struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
	Wins   uint          `json:"wins"`
}

// DuelLeaderboard is a slice of Duel-type leaderboard records
type DuelLeaderboard []DuelLeaderboardRecord

