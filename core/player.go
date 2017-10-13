package core

// PlayerID is the type for unique identification
// of a connected player
type PlayerID int

// PlayerIDs represents the type for a slice of of players' id's
type PlayerIDs []PlayerID

// AuthToken defines the type of the
// authentication token issued by the
// server upon successful authentication
// of the connecting player
type AuthToken string

// Player structure is an internal representation of
// a connected client that has been successfully
// authenticated
type Player struct {
	ID       PlayerID
	Nickname string
	Token    AuthToken
}

// Players represents a map of player entities,
// i.e. players connected to the server
type Players map[PlayerID]*Player
