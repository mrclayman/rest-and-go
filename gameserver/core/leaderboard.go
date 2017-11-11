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

// Len returns the number of records in a leaderboard
func (l Leaderboard) Len() int {
	return len(l)
}

// Swap swaps the records in the leaderboard
// at positions i and j
func (l Leaderboard) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Less compares the kill count record
// in records at positions i and j. Since
// we want to sort the leadeboard in
// descending order, we check that
// kill count at l[i] > kill count at l[j]
func (l Leaderboard) Less(i, j int) bool {
	return l[i].Kills > l[j].Kills
}

// GameLeaderboards defines a map of leaderboards
// identified by a given game type id
type GameLeaderboards map[GameType]*Leaderboard

