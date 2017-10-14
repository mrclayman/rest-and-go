package core

// LeaderboardRecord represents one player's record
// in a leaderboard for some game type.
type LeaderboardRecord struct {
	Player PlayerID
	Kills  uint
	Deaths uint
}

// Leaderboard defines a leaderboard, i.e.
// a list of player records for some specific
// game type
type Leaderboard []LeaderboardRecord

// GameLeaderboards defines a map of leaderboards
// identified by a given game type id
type GameLeaderboards map[GameType]*Leaderboard
