package data

// CTFLeaderboardRecord defines the structure
// of a record in a "Capture the Flag" leaderboard
// table
type CTFLeaderboardRecord struct {
	Player   Player `bson:"player"`
	Kills    uint   `bson:"kills"`
	Deaths   uint   `bson:"deaths"`
	Captures uint   `bson:"captures"`
}

// CTFLeaderboardRecords defines a type that is
// a list of CTF leaderboard records
type CTFLeaderboardRecords []CTFLeaderboardRecord

// CTFLeaderboard is an instance
// of CTFLeadeCTFLeaderboardRecords type and defines
// some basic CTF leaderboard items
var CTFLeaderboard = CTFLeaderboardRecords{
	{
		Player:   Player{ID: 4648464, Nick: "Lone_Hunter"},
		Kills:    456,
		Deaths:   111,
		Captures: 208,
	},
	{
		Player:   Player{ID: 5747548, Nick: "TheDamned1"},
		Kills:    794,
		Deaths:   552,
		Captures: 432,
	},
	{
		Player:   Player{ID: 1321878, Nick: "SoulScorcher"},
		Kills:    991,
		Deaths:   814,
		Captures: 555,
	},
	{
		Player:   Player{ID: 6735772, Nick: "Sir3n"},
		Kills:    614,
		Deaths:   516,
		Captures: 381,
	},
	{
		Player:   Player{ID: 6433858, Nick: "CrimsonDawn"},
		Kills:    416,
		Deaths:   311,
		Captures: 89,
	},
	{
		Player:   Player{ID: 5723425, Nick: "LittlePony"},
		Kills:    1551,
		Deaths:   1196,
		Captures: 899,
	},
	{
		Player:   Player{ID: 9661327, Nick: "Camping_Gaz"},
		Kills:    1441,
		Deaths:   1101,
		Captures: 1003,
	},
}
