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

// newLeaderboardTables generates a pre-filled
// map of leaderboard tables, one for each
// game type
func newLeaderboardTables() GameLeaderboards {
	retval := GameLeaderboards{
		DeathMatch: &Leaderboard{
			{
				Player: 8535253, // fatal1ty
				Kills:  1220,
				Deaths: 951,
			},
			{
				Player: 9464779, // Kr4zed
				Kills:  1188,
				Deaths: 1324,
			},
			{
				Player: 6992112, // ne0phyte
				Kills:  2242,
				Deaths: 711,
			},
			{
				Player: 1412491, // Dead3y3
				Kills:  1559,
				Deaths: 1332,
			},
			{
				Player: 4219691, // Mikky
				Kills:  1924,
				Deaths: 1422,
			},
			{
				Player: 6433858, // CrimsonDawn
				Kills:  255,
				Deaths: 192,
			},
		},
		CaptureTheFlag: &Leaderboard{
			{
				Player: 4648464, // Lone_Hunter
				Kills:  456,
				Deaths: 111,
			},
			{
				Player: 5747548, // TheDamned1
				Kills:  794,
				Deaths: 552,
			},
			{
				Player: 1321878, // SoulScorcher
				Kills:  991,
				Deaths: 814,
			},
			{
				Player: 6735772, // Sir3n
				Kills:  614,
				Deaths: 516,
			},
			{
				Player: 6433858, // CrimsonDawn
				Kills:  416,
				Deaths: 311,
			},
			{
				Player: 5723425, // LittlePony
				Kills:  1551,
				Deaths: 1196,
			},
			{
				Player: 9661327, // Camping_Gaz
				Kills:  1441,
				Deaths: 1101,
			},
		},
		LastManStanding: &Leaderboard{
			{
				Player: 5723425, // LittlePony
				Kills:  814,
				Deaths: 726,
			},
			{
				Player: 1412491, // Dead3y3
				Kills:  1944,
				Deaths: 1671,
			},
			{
				Player: 4648464, // Lone_Hunter
				Kills:  641,
				Deaths: 716,
			},
			{
				Player: 9464779, // Kr4zed
				Kills:  1516,
				Deaths: 1496,
			},
			{
				Player: 5457676, // How4rd
				Kills:  1744,
				Deaths: 1758,
			},
			{
				Player: 8712722, // Tweety
				Kills:  110,
				Deaths: 151,
			},
		},
		Duel: &Leaderboard{},
	}

	return retval
}
