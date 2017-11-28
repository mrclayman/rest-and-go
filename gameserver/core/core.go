package core

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/auth"
	"github.com/mrclayman/rest-and-go/gameserver/core/database"
	"github.com/mrclayman/rest-and-go/gameserver/core/servererrors"
	"github.com/mrclayman/rest-and-go/gameserver/core/match"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
	"gopkg.in/mgo.v2"
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
func NewCore(di *mgo.DialInfo) (*Core, error) {
	db, err := database.New(di)
	if err != nil {
		return nil, servererrors.InvalidArgumentError{Message: "Could not connect to database, invalid database URL"}
	}

	return &Core{
		db:               db,                             // Pre-fill the database
		players:          newConnectedPlayerTable(),      // Pre-fill the connected player table
		matches:          newMatchRegistry(),             // Pre-fill the match list
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
		return player.InvalidID, servererrors.LogicError{Message: "Duplicate login by player"}
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
		return player.Player{}, servererrors.InvalidArgumentError{Message: "Player with ID " + player.IDToString(ID) + " does not appear to be connected"}
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
	default:
		err = servererrors.InvalidArgumentError{Message: "Invalid game type '" + match.GameTypeToString(ID.Type) + "' in GetMatchForJSON"}
	}
	if err != nil {
		return nil, err
	}

	return retval, nil
}

// GetLeaderboardForJSON returns a transformed leaderboard
// associated with the given game type.
func (c *Core) GetLeaderboardForJSON(gt match.GameType) (interface{}, error) {
	var retval interface{}
	var err error

	switch gt {
	case match.DeathMatch:
		retval, err = c.db.GetDMLeaderboard()
	case match.CaptureTheFlag:
		retval, err = c.db.GetCTFLeaderboard()
	case match.LastManStanding:
		retval, err = c.db.GetLMSLeaderboard()
	case match.Duel:
		retval, err = c.db.GetDuelLeaderboard()
	default:
		err = servererrors.InvalidArgumentError{Message: "Unhandled game type '" + match.GameTypeToString(gt) + "' in GetLeaderboardForJSON()"}
	}

	if err != nil {
		return nil, err
	}

	return retval, nil
}

// JoinMatch lets a player with id 'pID' join a match 'mID',
// or create a new match of game type 'gt' if mID = InvalidMatchID.
// If 'mID' identifies a non-existent match, MatchNotFoundError
// is returned
func (c *Core) JoinMatch(mID match.ID, pID player.ID, token auth.WebSocketToken) (match.Number, error) {

	var p player.Player
	var err error
	if p, err = c.GetActivePlayer(pID); err != nil {
		return match.InvalidNumber, err
	}

	n := match.InvalidNumber
	switch mID.Type {
	case match.DeathMatch:
		n, err = c.joinDMMatch(mID.Number, p)
	case match.CaptureTheFlag:
		n, err = c.joinCTFMatch(mID.Number, p)
	case match.LastManStanding:
		n, err = c.joinLMSMatch(mID.Number, p)
	case match.Duel:
		n, err = c.joinDuelMatch(mID.Number, p)
	default:
		err = servererrors.InvalidArgumentError{Message: "Unhandled game type '" + match.GameTypeToString(mID.Type) + "' in JoinMatch()"}
	}

	if err != nil {
		return match.InvalidNumber, err
	}

	c.playerWSTokens[p.ID] = token

	return n, nil
}

// joinDMMatch either adds the player to an existing match
// or creates a new one
func (c *Core) joinDMMatch(n match.Number, p player.Player) (match.Number, error) {
	var m *match.DMMatch
	if n != match.InvalidNumber {
		// The player wants to join an existing match
		var err error
		m, err = c.matches.GetDM(n)
		if err != nil {
			return match.InvalidNumber, servererrors.InvalidArgumentError{Message: "DM match not found: " + match.NumberToString(n)}
		}
		m.Add(p)
	} else {
		// The player wants to create a new match
		pl := player.List{p}
		m = c.matches.NewDM(pl)
	}
	return m.Number, nil
}

// joinCTFMatch either adds the player to an existing match
// or creates a new one
func (c *Core) joinCTFMatch(n match.Number, p player.Player) (match.Number, error) {
	var m *match.CTFMatch
	if n != match.InvalidNumber {
		// The player wants to join an existing match
		var err error
		m, err = c.matches.GetCTF(n)
		if err != nil {
			return match.InvalidNumber, servererrors.InvalidArgumentError{Message: "CTF match not found: " + match.NumberToString(n)}
		}
		m.Add(p)
	} else {
		// The player wants to create a new match
		pl := player.List{p}
		m = c.matches.NewCTF(pl)
	}
	return m.Number, nil
}

// joinLMSMatch either adds the player to an existing match
// or creates a new one
func (c *Core) joinLMSMatch(n match.Number, p player.Player) (match.Number, error) {
	var m *match.LMSMatch
	if n != match.InvalidNumber {
		// The player wants to join an existing match
		var err error
		m, err = c.matches.GetLMS(n)
		if err != nil {
			return match.InvalidNumber, servererrors.InvalidArgumentError{Message: "LMS match not found: " + match.NumberToString(n)}
		}
		m.Add(p)
	} else {
		// The player wants to create a new match
		pl := player.List{p}
		m = c.matches.NewLMS(pl)
	}
	return m.Number, nil
}

// joinDuelMatch either adds the player to an existing match
// or creates a new one
func (c *Core) joinDuelMatch(n match.Number, p player.Player) (match.Number, error) {
	var m *match.DuelMatch
	if n != match.InvalidNumber {
		// The player wants to join an existing match
		var err error
		m, err = c.matches.GetDuel(n)
		if err != nil {
			return match.InvalidNumber, servererrors.InvalidArgumentError{Message: "Duel match not found: " + match.NumberToString(n)}
		}
		m.Add(p)
	} else {
		// The player wants to create a new match
		pl := player.List{p}
		m = c.matches.NewDuel(pl)
	}
	return m.Number, nil
}

// QuitMatch removes a player from the given match.
// If the match turns out to be empty, it is removed
// from the match set as well
func (c *Core) QuitMatch(mID match.ID, pID player.ID) error {

	var err error
	switch mID.Type {
	case match.DeathMatch:
		err = c.quitDMMatch(mID.Number, pID)
	case match.CaptureTheFlag:
		err = c.quitCTFMatch(mID.Number, pID)
	case match.LastManStanding:
		err = c.quitLMSMatch(mID.Number, pID)
	case match.Duel:
		err = c.quitDuelMatch(mID.Number, pID)
	}

	if err != nil {
		return err
	}
	delete(c.playerWSTokens, pID)

	return nil
}

// quitDMMatch removes a player with the given ID
// from a DeathMatch-type match with the given number
// In case the match could not be found, an error is raised.
// If the match ends up being empty after the player's removal
// it is dropped from the match registry
func (c *Core) quitDMMatch(n match.Number, pID player.ID) error {
	m, err := c.matches.GetDM(n)
	if err != nil {
		return err
	}

	if ok := m.Remove(pID); !ok {
		return servererrors.InvalidArgumentError{Message: "Player " + player.IDToString(pID) + " not present in match " + match.NumberToString(n)}
	}

	if len(m.Ranks) == 0 {
		c.matches.DropDM(m.Number)
	}

	return nil
}

// quitCTFMatch removes a player with the given ID
// from a CTF-type match with the given number
// In case the match could not be found, an error is raised.
// If the match ends up being empty after the player's removal
// it is dropped from the match registry
func (c *Core) quitCTFMatch(n match.Number, pID player.ID) error {
	m, err := c.matches.GetCTF(n)
	if err != nil {
		return err
	}

	if ok := m.Remove(pID); !ok {
		return servererrors.InvalidArgumentError{Message: "Player " + player.IDToString(pID) + " not present in match " + match.NumberToString(n)}
	}

	if len(m.Ranks) == 0 {
		c.matches.DropCTF(m.Number)
	}

	return nil
}

// quitLMSMatch removes a player with the given ID
// from a LMS-type match with the given number
// In case the match could not be found, an error is raised.
// If the match ends up being empty after the player's removal
// it is dropped from the match registry
func (c *Core) quitLMSMatch(n match.Number, pID player.ID) error {
	m, err := c.matches.GetLMS(n)
	if err != nil {
		return err
	}

	if ok := m.Remove(pID); !ok {
		return servererrors.InvalidArgumentError{Message: "Player " + player.IDToString(pID) + " not present in match " + match.NumberToString(n)}
	}

	if len(m.Ranks) == 0 {
		c.matches.DropLMS(m.Number)
	}

	return nil
}

// quitDuelMatch removes a player with the given ID
// from a DeathMatch-type match with the given number
// In case the match could not be found, an error is raised.
// If the match ends up being empty after the player's removal
// it is dropped from the match registry
func (c *Core) quitDuelMatch(n match.Number, pID player.ID) error {
	m, err := c.matches.GetDuel(n)
	if err != nil {
		return err
	}

	if ok := m.Remove(pID); !ok {
		return servererrors.InvalidArgumentError{Message: "Player " + player.IDToString(pID) + " not present in match " + match.NumberToString(n)}
	}

	if len(m.Ranks) == 0 {
		c.matches.DropDuel(m.Number)
	}

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
		return servererrors.InvalidArgumentError{Message: "Player " + player.IDToString(pID) + " not connected"}
	}

	delete(c.players, pID)
	return nil
}

// Cleanup performs finalization of
// of the core object
func (c *Core) Cleanup() {
	c.db.Close()
}
