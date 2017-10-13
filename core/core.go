package core

// Core is the data core of the application. Its members
// are accessible to HTTP request handlers to provide the
// client with the required information
type Core struct {
	db      *Database
	Players Players
	// TODO Add existing matches, leaderboards,...
}

// NewCore creates and returns a new core object
func NewCore() *Core {
	return &Core{db: NewDatabase(), Players: make(Players)}
}

// AuthenticatePlayer verifies the login credentials
// provided by the user with the database. Upon successful
// verification, 'true' and the player's internal ID are
// returned, otherwise 'false' and a zero player ID is returned
func (c *Core) AuthenticatePlayer(name, pass string) (PlayerID, bool) {
	return c.db.AuthenticatePlayer(name, pass)
}

// AddConnected adds a newly connected player to the system
func (c *Core) AddConnected(id PlayerID, nick string, token AuthToken) {
	c.Players[id] = &Player{ID: id, Nickname: nick, Token: token}
}
