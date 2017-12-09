package leaderboard

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

// LMSLeaderboardRecord contains information
// on record in a LMS-type leaderboard
type LMSLeaderboardRecord struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
	Wins   uint          `json:"wins"`
}

// LMSLeaderboard is a slice of LMS-type leaderboard records
type LMSLeaderboard []LMSLeaderboardRecord
