package core

import "sort"

// Core is the data core of the application. Its members
// are accessible to HTTP request handlers to provide the
// client with the required information
type Core struct {
	db               *Database
	players          Players
	matches          Matches
	playerAuthTokens PlayerAuthTokens
	playerWSTokens   PlayerWSTokens
}

// NewCore creates and returns a new core object
// with pre-filled member structures
func NewCore() *Core {
	return &Core{
		db:               newDatabase(),                  // Pre-fill the database
		players:          newConnectedPlayerTable(),      // Pre-fill the connected player table
		matches:          newMatchTable(),                // Pre-fill the match list
		playerAuthTokens: newConnectedPlayerTokenTable(), // Pre-fill the player auth token table
		playerWSTokens:   newPlayerInMatchTokenTable(),   // Pre-fill the player-in-match token table
	}
}

// AuthenticatePlayer verifies the login credentials
// provided by the user with the database. Upon successful
// verification, 'true' and the player's internal ID are
// returned, otherwise 'false' and a zero player ID is returned
func (c *Core) AuthenticatePlayer(name, pass string) (PlayerID, error) {
	id, err := c.db.AuthenticatePlayer(name, pass)
	if err != nil {
		return InvalidPlayerID, err
	} else if _, ok := c.players[id]; ok {
		return InvalidPlayerID, LogicError{"Duplicate login by player"}
	}

	return id, nil
}

// AddConnected adds a newly connected player to the system
func (c *Core) AddConnected(id PlayerID, nick string, token AuthToken) {
	c.players[id] = &Player{ID: id, Nick: nick}
	c.playerAuthTokens[id] = token
}

// IsLoggedIn checks that a player with the given id
// is active in the system and provided the correct
// authentication token
func (c *Core) IsLoggedIn(id PlayerID, token AuthToken) bool {
	playerToken, ok := c.playerAuthTokens[id]
	return ok && playerToken == token
}

// IsInMatch checks that a player is in match by
// comparing the provided WebSocket token with
// the one stored in the system
func (c *Core) IsInMatch(id PlayerID, token WebSocketToken) bool {
	playerToken, ok := c.playerWSTokens[id]
	return ok && playerToken == token
}

// GetPlayerNick returns the nickname for the player
// identified by id. If the player has not been found
// in the database, empty string and an appropriate
// boolean flag are returned
func (c *Core) GetPlayerNick(id PlayerID) (string, bool) {
	nick, ok := c.db.GetPlayerNick(id)
	return nick, ok
}

// GetMatchlistForJSON returns a transformed matchlist
// suitable for presentation to the client. I could
// just pass the whole structure to the JSON marshaller
// but I wanted to present the player with the other
// players' nicknames (those are not stored along with the
// players' match ranks) and not just their id's
func (c *Core) GetMatchlistForJSON() ([]map[string]interface{}, error) {
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

// GetLeaderboardForJSON returns a transformed leaderboard
// associated with the given game type.
func (c *Core) GetLeaderboardForJSON(gt GameType) ([]map[string]interface{}, error) {
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

// JoinMatch lets a player with id 'pid' join a match 'mid',
// or create a new match of game type 'gt' if mid = InvalidMatchID.
// If 'mid' identifies a non-existent match, MatchNotFoundError
// is returned
func (c *Core) JoinMatch(mid MatchID, pid PlayerID, token WebSocketToken, gt GameType) (MatchID, error) {
	var match *Match
	var ok bool

	if mid != InvalidMatchID {
		// The player wants to join an existing match
		match, ok = c.matches[mid]
		if !ok {
			return InvalidMatchID, InvalidArgumentError{"Match not found:" + MatchIDToString(mid)}
		}

		match.Add(pid)
	} else {
		// The player wants to create a new match
		if gt == InvalidGameType {
			return InvalidMatchID, InvalidArgumentError{"Game type specification required if no match ID defined"}
		}

		ids := PlayerIDs{pid}
		match = NewMatchWithPlayers(gt, ids)
		c.matches[match.ID] = match
	}

	c.playerWSTokens[pid] = token

	return match.ID, nil
}

// QuitMatch removes a player from the given match.
// If the match turns out to be empty, it is removed
// from the match set
func (c *Core) QuitMatch(mid MatchID, pid PlayerID) error {

	match, ok := c.matches[mid]
	if !ok {
		return InvalidArgumentError{"Invalid match ID " + MatchIDToString(mid)}
	}
	delete(match.Ranks, pid)
	if len(match.Ranks) == 0 {
		delete(c.matches, mid)
	}
	delete(c.playerWSTokens, pid)

	return nil
}
