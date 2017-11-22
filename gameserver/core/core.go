package core

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/auth"
	"github.com/mrclayman/rest-and-go/gameserver/core/database"
	"github.com/mrclayman/rest-and-go/gameserver/core/errors"
	"github.com/mrclayman/rest-and-go/gameserver/core/match"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
)

// PlayerAuthTokens aggregates authentication
// tokens of connected players
type PlayerAuthTokens map[player.ID]auth.AuthToken

// PlayerWSTokens aggregates WebSocket tokens
// of players participating in matches
type PlayerWSTokens map[player.ID]auth.WebSocketToken

// Core is the data core of the application. Its members
// are accessible to HTTP request handlers to provide the
// client with the required information
type Core struct {
	db               *database.Database
	players          player.Map
	matches          match.Registry
	playerAuthTokens PlayerAuthTokens
	playerWSTokens   PlayerWSTokens
}

// NewCore creates and returns a new core object
// with pre-filled member structures
func NewCore(dbURL string) (*Core, error) {
	db, err := database.New(dbURL)
	if err != nil {
		return nil, errors.InvalidArgumentError{Message: "Could not connect to database, invalid database URL"}
	}

	return &Core{
		db:               db,                             // Pre-fill the database
		players:          newConnectedPlayerTable(),      // Pre-fill the connected player table
		matches:          newMatchTable(),                // Pre-fill the match list
		playerAuthTokens: newConnectedPlayerTokenTable(), // Pre-fill the player auth token table
		playerWSTokens:   newPlayerInMatchTokenTable(),   // Pre-fill the player-in-match token table
	}, nil
}

// AuthenticatePlayer verifies the login credentials
// provided by the user with the database. Upon successful
// verification, the player's internal ID and nil are
// returned, otherwise an invalid player ID and the error
// object are returned
func (c *Core) AuthenticatePlayer(name, pass string) (player.ID, error) {
	id, err := c.db.AuthenticatePlayer(name, pass)
	if err != nil {
		return player.InvalidID, err
	} else if _, ok := c.players[id]; ok {
		return player.InvalidID, errors.LogicError{Message: "Duplicate login by player"}
	}

	return id, nil
}

// AddConnected adds a newly connected player to the system
func (c *Core) AddConnected(ID player.ID, nick string, token auth.AuthToken) {
	c.players[ID] = player.Player{ID: ID, Nick: nick}
	c.playerAuthTokens[ID] = token
}

// IsLoggedIn checks that a player with the given id
// is active in the system and provided the correct
// authentication token
func (c *Core) IsLoggedIn(ID player.ID, token auth.AuthToken) bool {
	playerToken, ok := c.playerAuthTokens[ID]
	return ok && playerToken == token
}

// IsInMatch checks that a player is in match by
// comparing the provided WebSocket token with
// the one stored in the system
func (c *Core) IsInMatch(ID player.ID, token auth.WebSocketToken) bool {
	playerToken, ok := c.playerWSTokens[ID]
	return ok && playerToken == token
}

// GetPlayerNick returns the nickname for the player
// identified by id. If the player has not been found
// in the database, empty string and an appropriate
// boolean flag are returned
func (c *Core) GetPlayerNick(ID player.ID) (string, error) {
	return c.db.GetPlayerNick(ID)
}

// GetActivePlayer looks into the registry of
// connected player and tries to find the entry
// corresponding to the ID provided in the argument
func (c *Core) GetActivePlayer(ID player.ID) (player.Player, error) {
	var p player.Player
	var ok bool

	if p, ok = c.players[ID]; !ok {
		return player.Player{}, errors.InvalidArgumentError{Message: "Player with ID " + player.IDToString(ID) + " does not appear to be connected"}
	}

	return p, nil
}

// GetMatchlistForJSON returns a transformed matchlist
// suitable for presentation to the client.
func (c *Core) GetMatchlistForJSON() map[match.GameType]interface{} {
	retval := make(map[match.GameType]interface{})

	retval[match.DeathMatch] = c.matches.GetAllDM()
	retval[match.CaptureTheFlag] = c.matches.GetAllCTF()
	retval[match.LastManStanding] = c.matches.GetAllLMS()
	retval[match.Duel] = c.matches.GetAllDuel()
	return retval
}

// GetMatchForJSON returns a serialized version of a
// match instance with the given ID.
func (c *Core) GetMatchForJSON(ID match.ID) (interface{}, error) {
	var retval interface{}
	var err error
	switch ID.Type {
	case match.DeathMatch:
		retval, err = c.matches.GetDM(ID.Number)
	case match.CaptureTheFlag:
		retval, err = c.matches.GetCTF(ID.Number)
	case match.LastManStanding:
		retval, err = c.matches.GetLMS(ID.Number)
	case match.Duel:
		retval, err = c.matches.GetDuel(ID.Number)
	}
	if err != nil {
		return nil, err
	}

	return retval, nil
}

// GetLeaderboardForJSON returns a transformed leaderboard
// associated with the given game type.
func (c *Core) GetLeaderboardForJSON(gt match.GameType) (interface{}, error) {
	return c.db.GetLeaderboard(gt)
}

// JoinMatch lets a player with id 'pID' join a match 'mID',
// or create a new match of game type 'gt' if mID = InvalidMatchID.
// If 'mID' identifies a non-existent match, MatchNotFoundError
// is returned
// TODO Continue here!!
func (c *Core) JoinMatch(mID match.ID, pID player.ID, token auth.WebSocketToken, gt match.GameType) (match.ID, error) {
	var m *match.Match
	var ok bool

	var p player.Player
	var err error

	if p, err = c.GetActivePlayer(pID); err != nil {
		return match.InvalidID, err
	}

	if mID != match.InvalidID {
		// The player wants to join an existing match
		m, ok = c.matches[mID]
		if !ok {
			return match.InvalidID, errors.InvalidArgumentError{Message: "Match not found: " + match.IDToString(mID)}
		}

		m.Add(p)
	} else {
		// The player wants to create a new match
		if gt == match.InvalidGameType {
			return match.InvalidID, errors.InvalidArgumentError{Message: "Game type specification required if no match ID defined"}
		}

		pl := player.List{p}
		var err error
		if m, err = match.New(gt, pl); err != nil {
			return match.InvalidID, err
		}

		c.matches[m.ID] = m
	}

	c.playerWSTokens[p.ID] = token

	return m.ID, nil
}

// QuitMatch removes a player from the given match.
// If the match turns out to be empty, it is removed
// from the match set as well
func (c *Core) QuitMatch(mID match.ID, pID player.ID) error {

	m, ok := c.matches[mID]
	if !ok {
		return errors.InvalidArgumentError{Message: "Invalid match ID " + match.IDToString(mID)}
	}

	delete(m.Ranks, pID)
	if len(m.Ranks) == 0 {
		delete(c.matches, mID)
	}
	delete(c.playerWSTokens, pID)

	return nil
}

// QuitPlayer performs logout procedure, removing
// the player from the table of connected players
func (c *Core) QuitPlayer(pID player.ID) error {

	// I know that just attempting to delete the player ID
	// from the player map wouldn't hurt even if it wasn't
	// there, but I wanted to indicate that something
	// is wrong if the logout attempt has been made
	// without the player being in the system to begin with
	if _, ok := c.players[pID]; !ok {
		return errors.InvalidArgumentError{Message: "Player " + player.IDToString(pID) + " not connected"}
	}

	delete(c.players, pID)
	return nil
}

// Cleanup performs finalization of
// of the core object
func (c *Core) Cleanup() {
	c.db.Close()
}
