package leaderboard

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

// DMLeaderboardRecord contains information
// on record in a DeathMatch-type
// leaderboard
type DMLeaderboardRecord struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// DMLeaderboard is a slice of DM leaderboard records
type DMLeaderboard []DMLeaderboardRecord
