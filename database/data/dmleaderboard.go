package data

// DMLeaderboardRecord defines the basic structure
// of a leaderboard record for the DeathMatch game type
type DMLeaderboardRecord struct {
	Player Player `bson:"player"`
	Kills  uint   `bson:"kills"`
	Deaths uint   `bson:"deaths"`
}

// DMLeaderboardRecords defines a type for
// the DeathMatch game type leaderboard
type DMLeaderboardRecords []DMLeaderboardRecord

// DMLeaderboard is an instance of
// DMLeaderboardRecords and defines a few records
// in the DeathMatch leaderboard
var DMLeaderboard = DMLeaderboardRecords{
	{
		Player: Player{ID: 8535253, Nick: "fatal1ty"},
		Kills:  1220,
		Deaths: 951,
	},
	{
		Player: Player{ID: 9464779, Nick: "Kr4zed"},
		Kills:  1188,
		Deaths: 1324,
	},
	{
		Player: Player{ID: 6992112, Nick: "ne0phyte"},
		Kills:  2242,
		Deaths: 711,
	},
	{
		Player: Player{ID: 1412491, Nick: "Dead3y3"},
		Kills:  1559,
		Deaths: 1332,
	},
	{
		Player: Player{ID: 4219691, Nick: "Mikky"},
		Kills:  1924,
		Deaths: 1422,
	},
	{
		Player: Player{ID: 6433858, Nick: "CrimsonDawn"},
		Kills:  255,
		Deaths: 192,
	},
}
