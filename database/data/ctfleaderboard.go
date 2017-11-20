package data

// CTFLeaderboardRecord defines the structure
// of a record in a "Capture the Flag" leaderboard
// table
type CTFLeaderboardRecord struct {
	PlayerID   int
	PlayerName string
	Kills      uint
	Deaths     uint
	Captures   uint
}

// CTFLeaderboardRecords defines a type that is
// a list of CTF leaderboard records
type CTFLeaderboardRecords []CTFLeaderboardRecord

// CTFLeaderboard is an instance
// of CTFLeadeCTFLeaderboardRecords type and defines
// some basic CTF leaderboard items
var CTFLeaderboard = CTFLeaderboardRecords{
	{
		PlayerID:   4648464,
		PlayerName: "Lone_Hunter",
		Kills:      456,
		Deaths:     111,
		Captures:   208,
	},
	{
		PlayerID:   5747548,
		PlayerName: "TheDamned1",
		Kills:      794,
		Deaths:     552,
		Captures:   432,
	},
	{
		PlayerID:   1321878,
		PlayerName: "SoulScorcher",
		Kills:      991,
		Deaths:     814,
		Captures:   555,
	},
	{
		PlayerID:   6735772,
		PlayerName: "Sir3n",
		Kills:      614,
		Deaths:     516,
		Captures:   381,
	},
	{
		PlayerID:   6433858,
		PlayerName: "CrimsonDawn",
		Kills:      416,
		Deaths:     311,
		Captures:   89,
	},
	{
		PlayerID:   5723425,
		PlayerName: "LittlePony",
		Kills:      1551,
		Deaths:     1196,
		Captures:   899,
	},
	{
		PlayerID:   9661327,
		PlayerName: "Camping_Gaz",
		Kills:      1441,
		Deaths:     1101,
		Captures:   1003,
	},
}
