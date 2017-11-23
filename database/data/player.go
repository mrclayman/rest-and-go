package data

// Player is a representation of a player
// in a leaderboard
type Player struct {
	ID   int    `bson:"id"`
	Nick string `bson:"nick"`
}
