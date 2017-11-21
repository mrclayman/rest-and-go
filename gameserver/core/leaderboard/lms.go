package leaderboard

import "github.com/mrclayman/rest-and-go/gameserver/core/player"

// LMSLeaderboardRecord represents a record
// in the Last-Man-Standing game type leaderboard
type LMSLeaderboardRecord struct {
	Player player.Player `json:"player" bson:"player"`
	Kills  uint          `json:"kills" bson:"kills"`
	Deaths uint          `json:"deaths" bson:"deaths"`
	Wins   uint          `json:"wins" bson:"wins"`
}

// LMSLeaderboard represents a slice of the
// Last-Man-Standing game type leaderboards
type LMSLeaderboard []LMSLeaderboardRecord
