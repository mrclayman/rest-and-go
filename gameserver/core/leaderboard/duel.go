package leaderboard

import "github.com/mrclayman/rest-and-go/gameserver/core/player"

// DuelLeaderboardRecord defines the structure
// of the Duel game type leaderboard (currently,
// it is shared with the LMS game type record)
type DuelLeaderboardRecord struct {
	Player player.Player `json:"player" bson:"player"`
	Kills  uint          `json:"kills" bson:"kills"`
	Deaths uint          `json:"deaths" bson:"deaths"`
	Wins   uint          `json:"wins" bson:"wins"`
}

// DuelLeaderboard represents a slice of the
// Duel game type leaderboard
type DuelLeaderboard []DuelLeaderboardRecord
