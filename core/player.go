package core

// PlayerIDType is the type for unique identification
// of a connected player
type PlayerIDType int

// PlayerType structure is an internal representation of
// a connected client that has been successfully
// authenticated
type PlayerType struct {
	PlayerID  PlayerIDType
	Nickname  string
	AuthToken string
}

// PlayersType represents a map of player entities
type PlayersType map[PlayerIDType]*PlayerType
