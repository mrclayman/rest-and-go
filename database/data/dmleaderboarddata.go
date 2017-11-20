package data

// DMLeaderboardRecord defines the basic structure
// of a leaderboard record for the DeathMatch game type
type DMLeaderboardRecord struct {
	PlayerID   int
	PlayerName string
	Kills      uint
	Deaths     uint
}

// DMLeaderboardRecords defines a type for
// the DeathMatch game type leaderboard
type DMLeaderboardRecords []DMLeaderboardRecord

// DMLeaderboard is an instance of
// DMLeaderboardRecords and defines a few records
// in the DeathMatch leaderboard
var DMLeaderboard = DMLeaderboardRecords{
	{
		PlayerID:   8535253,
		PlayerName: "fatal1ty",
		Kills:      1220,
		Deaths:     951,
	},
	{
		PlayerID:   9464779,
		PlayerName: "Kr4zed",
		Kills:      1188,
		Deaths:     1324,
	},
	{
		PlayerID:   6992112,
		PlayerName: "ne0phyte",
		Kills:      2242,
		Deaths:     711,
	},
	{
		PlayerID:   1412491,
		PlayerName: "Dead3y3",
		Kills:      1559,
		Deaths:     1332,
	},
	{
		PlayerID:   4219691,
		PlayerName: "Mikky",
		Kills:      1924,
		Deaths:     1422,
	},
	{
		PlayerID:   6433858,
		PlayerName: "CrimsonDawn",
		Kills:      255,
		Deaths:     192,
	},
}
