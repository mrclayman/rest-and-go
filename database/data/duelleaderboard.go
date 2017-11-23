package data

// DuelLeaderboardRecord defines the structure for
// the Duel-type leaderboard record
type DuelLeaderboardRecord LMSLeaderboardRecord

// DuelLeaderboardRecords defines the type for
// a slice of Duel-type leaderboard records
type DuelLeaderboardRecords []DuelLeaderboardRecord

// DuelLeaderboard is an (empty) instance of
// the Duel-type leaderboard
var DuelLeaderboard DuelLeaderboardRecords
