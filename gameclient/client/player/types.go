package player

// ID is the type of a player's id
type ID int

// Player structure defines the basic credentials of
// the player after getting connected to the server
type Player struct {
	ID   ID     `json:"id"`
	Nick string `json:"nick"`
}

// Login contains login information
// of a player to be sent to the server
type Login struct {
	Nick     string `json:"name"`
	Password string `json:"password"`
}
