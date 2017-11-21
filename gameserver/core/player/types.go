package player

// ID is the type for unique identification
// of a connected player
type ID int

// InvalidID defines an invalid value
// for the player id
const InvalidID ID = 0

// IDs represents the type for a slice of players' id's
type IDs []ID

// Player structure is an internal representation of
// a connected client that has been successfully
// authenticated
type Player struct {
	ID   ID     `json:"player_id" bson:"playerid"`
	Nick string `json:"player_name" bson:"playername"`
}

// List defines a type for a slice
// of player entities
type List []Player

// Map represents a map of player entities,
// i.e. players connected to the server
type Map map[ID]Player
