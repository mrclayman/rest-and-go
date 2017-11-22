package player

import "strconv"

// Player structure defines the basic credentials of
// the player connected to the server
type Player struct {
	ID   int    `json:"player_id"`
	Name string `json:"player_name"`
}

// AuthData contains information on a player's
// authentication data returned by the server
type AuthData struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

// ToGet converts the contents of the
// AuthData instance into GET request
// arguments
func (data AuthData) ToGet() string {
	return "id=" + strconv.Itoa(data.ID) + "&token=" + data.Token
}

// Login contains login information
// for a player
type Login struct {
	Nick     string `json:"name"`
	Password string `json:"password"`
}
