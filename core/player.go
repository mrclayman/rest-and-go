package core

import "strconv"

// PlayerID is the type for unique identification
// of a connected player
type PlayerID int

// InvalidPlayerID defines an invalid value
// for the player id
const InvalidPlayerID PlayerID = 0

// PlayerIDs represents the type for a slice of players' id's
type PlayerIDs []PlayerID

// StringToPlayerID converts a string into
// an equivalent value of the type PlayerID
func StringToPlayerID(strID string) (PlayerID, error) {
	id, err := strconv.Atoi(strID)
	return PlayerID(id), err
}

// Player structure is an internal representation of
// a connected client that has been successfully
// authenticated
type Player struct {
	ID    PlayerID
	Nick  string
	Token AuthToken
}

// Players represents a map of player entities,
// i.e. players connected to the server
type Players map[PlayerID]*Player

func newConnectedPlayerTable() Players {
	// The list below corresponds to the complete
	// list of players already participating in matches
	players := []*Player{
		{
			ID:    8535253,
			Nick:  "fatal1ty",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    5457676,
			Nick:  "How4rd",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    9464779,
			Nick:  "Kr4zed",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    6735772,
			Nick:  "Sir3n",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    6433858,
			Nick:  "CrimsonDawn",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    4148994,
			Nick:  "JigSaw",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    9661327,
			Nick:  "Camping_Gaz",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    1412491,
			Nick:  "Dead3y3",
			Token: GenerateAuthenticationToken(),
		},
		{
			ID:    8712722,
			Nick:  "Tweety",
			Token: GenerateAuthenticationToken(),
		},
	}

	retval := make(Players, len(players))
	for _, player := range players {
		retval[player.ID] = player
	}

	return retval
}
