package core

// LeaderboardRecord represents one player's record
// in a leaderboard for some match type
type LeaderboardRecord struct {
	PID         PlayerID
	TotalKills  int
	TotalDeaths int
}

// Leaderboard defines a leaderboard, i.e.
// a list of player records for some specific
// match type
type Leaderboard []LeaderboardRecord

// MatchTypeLeaderboardType defines a map of leaderboards
// identified by a given match type id
type MatchTypeLeaderboardType map[MatchTypeIDType]*Leaderboard
