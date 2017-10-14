package core

import "fmt"

// Core is the data core of the application. Its members
// are accessible to HTTP request handlers to provide the
// client with the required information
type Core struct {
	db      *Database
	players Players
	matches Matches
	// TODO Add existing matches, leaderboards,...
}

// NewCore creates and returns a new core object
func NewCore() *Core {
	return &Core{
		db:      newDatabase(),             // Pre-fill the database
		players: newConnectedPlayerTable(), // Pre-fill the connected player table
		matches: newMatchTable(),           // Pre-fill the match list
	}
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
	c.players[id] = &Player{ID: id, Nick: nick, Token: token}
}

// IsLoggedIn checks that a player with the given id
// is active in the system and provided the correct
// authentication token
func (c *Core) IsLoggedIn(id PlayerID, token AuthToken) bool {
	player, ok := c.players[id]
	return ok && player.Token == token
}

// GetPlayerNick returns the nickname for the player
// identified by id. If the player has not been found
// in the list of connected players, empty string and
// an appropriate boolean flag are returned
func (c *Core) GetPlayerNick(id PlayerID) (string, bool) {
	var player *Player
	var ok bool
	if player, ok = c.players[id]; !ok {
		return "", ok
	}

	return player.Nick, ok
}

// GetMatchlist returns a transformed matchlist
// suitable for presentation to the client. I could
// just pass the whole structure to the JSON marshaller
// but I wanted to present the player with the other
// players' nicknames and not just their id's
func (c *Core) GetMatchlist() []map[string]interface{} {
	retval := make([]map[string]interface{}, len(c.matches))
	i := 0
	for _, match := range c.matches {
		matchItem := make(map[string]interface{})
		matchItem["match_id"] = match.ID
		matchItem["match_type"] = match.Type

		matchRanks := make([]map[string]interface{}, len(match.Ranks))
		j := 0
		for _, rank := range match.Ranks {
			playerRank := make(map[string]interface{})
			playerRank["player_id"] = rank.Player

			playerNick, ok := c.GetPlayerNick(rank.Player)
			if !ok {
				fmt.Println("Couldn't find nickname for player", rank.Player)
				continue
			}
			playerRank["player_name"] = playerNick
			playerRank["kills"] = rank.Kills
			playerRank["deaths"] = rank.Deaths
			matchRanks[j] = playerRank
			j++
		}
		matchItem["ranks"] = matchRanks
		retval[i] = matchItem
		i++
	}

	return retval
}
