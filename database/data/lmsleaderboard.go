package data

// LMSLeaderboardRecord defines the basic structure
// for the Last Man Standing leaderboard record
type LMSLeaderboardRecord struct {
	Player Player `bson:"player"`
	Kills  uint   `bson:"kills"`
	Deaths uint   `bson:"deaths"`
	Wins   uint   `bson:"wins"`
}

// LMSLeaderboardRecords defines the type for an
// array of LMS game type leaderboard records
type LMSLeaderboardRecords []LMSLeaderboardRecord

// LMSLeaderboard instances the LMSLeaderboardRecords
// type and defines some basic values for
// the LMS game type leaderboard
var LMSLeaderboard = LMSLeaderboardRecords{
	{
		Player: Player{ID: 5723425, Nick: "LittlePony"},
		Kills:  814,
		Deaths: 726,
		Wins:   391,
	},
	{
		Player: Player{ID: 1412491, Nick: "Dead3y3"},
		Kills:  1944,
		Deaths: 1671,
		Wins:   1222,
	},
	{
		Player: Player{ID: 4648464, Nick: "Lone_Hunter"},
		Kills:  641,
		Deaths: 716,
		Wins:   177,
	},
	{
		Player: Player{ID: 9464779, Nick: "Kr4zed"},
		Kills:  1516,
		Deaths: 1496,
		Wins:   661,
	},
	{
		Player: Player{ID: 5457676, Nick: "How4rd"},
		Kills:  1744,
		Deaths: 1758,
		Wins:   1321,
	},
	{
		Player: Player{ID: 8712722, Nick: "Tweety"},
		Kills:  110,
		Deaths: 151,
		Wins:   11,
	},
}
