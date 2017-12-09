package leaderboard

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

// CTFLeaderboardRecord contains information
// on record in a CTF-type leaderboard
type CTFLeaderboardRecord struct {
	Player   player.Player `json:"player"`
	Kills    uint          `json:"kills"`
	Deaths   uint          `json:"deaths"`
	Captures uint          `json:"captures"`
}

// CTFLeaderboard is a slice of CTF leaderboard records
type CTFLeaderboard []CTFLeaderboardRecord

