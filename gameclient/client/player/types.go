package player

// Player structure defines the basic credentials of
// the player after getting connected to the server
type Player struct {
	ID   int    `json:"id"`
	Nick string `json:"nick"`
}

// Login contains login information
// of a player to be sent to the server
type Login struct {
	Nick     string `json:"name"`
	Password string `json:"password"`
}
