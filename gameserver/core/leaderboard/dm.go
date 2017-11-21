package leaderboard

import "github.com/mrclayman/rest-and-go/gameserver/core/player"

// DMLeaderboardRecord represents a record
// in the Deathmatch game type leaderboard
type DMLeaderboardRecord struct {
	Player player.Player `json:"player" bson:"player"`
	Kills  uint          `json:"kills" bson:"kills"`
	Deaths uint          `json:"deaths" bson:"deaths"`
}

// DMLeaderboard represents a slice of
// DeathMatch leaderboard records
type DMLeaderboard []DMLeaderboardRecord
