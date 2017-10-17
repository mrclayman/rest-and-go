package core

import "strconv"

// PlayerID is the type for unique identification
// of a connected player
type PlayerID int

// InvalidPlayerID defines an invalid value
// for the player id
const InvalidPlayerID PlayerID = 0

// PlayerIDs represents the type for a slice of players' id's
type PlayerIDs []PlayerID

// StringToPlayerID converts a string into
// an equivalent value of the type PlayerID
func StringToPlayerID(strID string) (PlayerID, error) {
	id, err := strconv.Atoi(strID)
	return PlayerID(id), err
}

// PlayerIDToString converts a player's id
// into its string equivalent
func PlayerIDToString(id PlayerID) string {
	return strconv.Itoa(int(id))
}

// Player structure is an internal representation of
// a connected client that has been successfully
// authenticated
type Player struct {
	ID   PlayerID
	Nick string
}

// PlayerAuthTokens aggregates authentication
// tokens of connected players
type PlayerAuthTokens map[PlayerID]AuthToken

// PlayerWSTokens aggregates WebSocket tokens
// of players participating in matches
type PlayerWSTokens map[PlayerID]WebSocketToken

// Players represents a map of player entities,
// i.e. players connected to the server
type Players map[PlayerID]*Player
