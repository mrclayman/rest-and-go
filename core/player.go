package core

/* Player structure is an internal representation of
 * a connected client that has been successfully
 * authenticated
 */
type Player struct {
	Nickname string
	AuthToken string
}

// Players represents a slice of player entities
type Players []Player
