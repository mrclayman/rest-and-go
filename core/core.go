package core

import "sort"

// Core is the data core of the application. Its members
// are accessible to HTTP request handlers to provide the
// client with the required information
type Core struct {
	db      *Database
	players Players
	matches Matches
}

// NewCore creates and returns a new core object
// with pre-filled member structures
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
// in the database, empty string and an appropriate
// boolean flag are returned
func (c *Core) GetPlayerNick(id PlayerID) (string, bool) {
	nick, ok := c.db.GetPlayerNick(id)
	return nick, ok
}

// GetMatchlist returns a transformed matchlist
// suitable for presentation to the client. I could
// just pass the whole structure to the JSON marshaller
// but I wanted to present the player with the other
// players' nicknames (those are not stored along with the
// players' match ranks) and not just their id's
func (c *Core) GetMatchlist() ([]map[string]interface{}, error) {
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
				return nil, IntegrityError{"Could not find nickname for player" + PlayerIDToString(rank.Player)}
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

	return retval, nil
}

// GetLeaderboard returns a transformed leaderboard
// associated with the given game type.
func (c *Core) GetLeaderboard(gt GameType) ([]map[string]interface{}, error) {
	board, ok := c.db.GetLeaderboard(gt)
	if !ok {
		return nil, IntegrityError{"Leaderboard has not been created for type " + GameTypeToString(gt)}
	}

	sortedBoard := *board
	sort.Sort(sortedBoard)
	retval := make([]map[string]interface{}, len(sortedBoard))
	for i, record := range sortedBoard {
		item := make(map[string]interface{})
		item["player_id"] = record.Player

		playerNick, ok := c.GetPlayerNick(record.Player)
		if !ok {
			return nil, IntegrityError{"Could not find nickname for player" + PlayerIDToString(record.Player)}
		}
		item["player_name"] = playerNick
		item["kills"] = record.Kills
		item["deaths"] = record.Deaths
		retval[i] = item
	}
	return retval, nil
}
