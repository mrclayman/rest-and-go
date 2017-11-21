package leaderboard

import "github.com/mrclayman/rest-and-go/gameserver/core/player"

// CTFLeaderboardRecord represents a record
// in the Capture-the-Flag game type leaderboard
type CTFLeaderboardRecord struct {
	Player   player.Player `json:"player" bson:"player"`
	Kills    uint          `json:"kills" bson:"kills"`
	Deaths   uint          `json:"deaths" bson:"deaths"`
	Captures uint          `json:"captures" bson:"captures"`
}

// CTFLeaderboard represents a slice of the
// Capture-the-Flag game type leaderboard records
type CTFLeaderboard []CTFLeaderboardRecord
