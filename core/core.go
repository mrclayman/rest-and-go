package core

// Core is the data core of the application. Its members
// are accessible to HTTP request handlers to provide the
// client with the required information
type Core struct {
	db *Database
	// TODO Add existing matches, leaderboards,...
}

// NewCore creates and returns a new core object
func NewCore() *Core {
	return &Core{db: NewDatabase()}
}
