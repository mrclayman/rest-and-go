package core

var (
	dmLeaderboardSortingCriterion   = []string{"-kills", "deaths"}
	ctfLeaderboardSortingCriterion  = []string{"-captures", "-kills", "deaths"}
	lmsLeaderboardSortingCriterion  = []string{"-wins", "-kills", "deaths"}
	duelLeaderboardSortingCriterion = lmsLeaderboardSortingCriterion
)

// DMLeaderboardRecord represents a record
// in the Deathmatch game type leaderboard
type DMLeaderboardRecord struct {
	PlayerID   PlayerID `json:"player_id" bson:"playerid"`
	PlayerName string   `json:"player_name" bson:"playername"`
	Kills      uint     `json:"kills" bson:"kills"`
	Deaths     uint     `json:"deaths" bson:"deaths"`
}

// DMLeaderboard represents a slice of
// DeathMatch leaderboard records
type DMLeaderboard []DMLeaderboardRecord

// CTFLeaderboardRecord represents a record
// in the Capture-the-Flag game type leaderboard
type CTFLeaderboardRecord struct {
	PlayerID   PlayerID `json:"player_id" bson:"playerid"`
	PlayerName string   `json:"player_name" bson:"playername"`
	Kills      uint     `json:"kills" bson:"kills"`
	Deaths     uint     `json:"deaths" bson:"deaths"`
	Captures   uint     `json:"captures" bson:"captures"`
}

// CTFLeaderboard represents a slice of the
// Capture-the-Flag game type leaderboard records
type CTFLeaderboard []CTFLeaderboardRecord

// LMSLeaderboardRecord represents a record
// in the Last-Man-Standing game type leaderboard
type LMSLeaderboardRecord struct {
	PlayerID   PlayerID `json:"player_id" bson:"playerid"`
	PlayerName string   `json:"player_name" bson:"playername"`
	Kills      uint     `json:"kills" bson:"kills"`
	Deaths     uint     `json:"deaths" bson:"deaths"`
	Wins       uint     `json:"wins" bson:"wins"`
}

// LMSLeaderboard represents a slice of the
// Last-Man-Standing game type leaderboards
type LMSLeaderboard []LMSLeaderboardRecord

// DuelLeaderboardRecord defines the structure
// of the Duel game type leaderboard (currently,
// it is shared with the LMS game type record)
type DuelLeaderboardRecord LMSLeaderboardRecord

// DuelLeaderboard represents a slice of the
// Duel game type leaderboard
type DuelLeaderboard LMSLeaderboard
