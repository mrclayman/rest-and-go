package data

// LMSLeaderboardRecord defines the basic structure
// for the Last Man Standing leaderboard record
type LMSLeaderboardRecord struct {
	PlayerID   int
	PlayerName string
	Kills      uint
	Deaths     uint
	Wins       uint
}

// LMSLeaderboardRecords defines the type for an
// array of LMS game type leaderboard records
type LMSLeaderboardRecords []LMSLeaderboardRecord

// LMSLeaderboard instances the LMSLeaderboardRecords
// type and defines some basic values for
// the LMS game type leaderboard
var LMSLeaderboard = LMSLeaderboardRecords{
	{
		PlayerID:   5723425,
		PlayerName: "LittlePony",
		Kills:      814,
		Deaths:     726,
		Wins:       391,
	},
	{
		PlayerID:   1412491,
		PlayerName: "Dead3y3",
		Kills:      1944,
		Deaths:     1671,
		Wins:       1222,
	},
	{
		PlayerID:   4648464,
		PlayerName: "Lone_Hunter",
		Kills:      641,
		Deaths:     716,
		Wins:       177,
	},
	{
		PlayerID:   9464779,
		PlayerName: "Kr4zed",
		Kills:      1516,
		Deaths:     1496,
		Wins:       661,
	},
	{
		PlayerID:   5457676,
		PlayerName: "How4rd",
		Kills:      1744,
		Deaths:     1758,
		Wins:       1321,
	},
	{
		PlayerID:   8712722,
		PlayerName: "Tweety",
		Kills:      110,
		Deaths:     151,
		Wins:       11,
	},
}
