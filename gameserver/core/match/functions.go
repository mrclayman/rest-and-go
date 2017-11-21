package match

import (
	"math/rand"
	"strconv"

	"github.com/mrclayman/rest-and-go/gameserver/core/errors"
	"github.com/mrclayman/rest-and-go/gameserver/core/leaderboard"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
)

// GenerateID generates a random match ID
// TODO Hide once the dummy data used in the server are no longer needed
func GenerateID() ID {
	return ID(rand.Uint64())
}

// IDToString converts a match id into
// a string
func IDToString(ID ID) string {
	return strconv.FormatUint(uint64(ID), 10)
}

// New creates a new match and populates
// it with the given set of players
func New(gt GameType, pl player.List) (*Match, error) {
	ID := GenerateID()
	ranks := make(PlayerRanks)

	for _, p := range pl {
		var r interface{}
		var err error

		if r, err = createNewLeaderboardRecord(gt, p); err != nil {
			return nil, err
		}

		ranks[p.ID] = r
	}

	return &Match{
		ID:    ID,
		Type:  gt,
		Ranks: ranks,
	}, nil
}

// createNewLeaderboardRecord creates a new leaderboard record
// for a player based on the game type
func createNewLeaderboardRecord(gt GameType, p player.Player) (interface{}, error) {
	switch gt {
	case DeathMatch:
		return leaderboard.DMLeaderboardRecord{Player: p}, nil
	case CaptureTheFlag:
		return leaderboard.CTFLeaderboardRecord{Player: p}, nil
	case LastManStanding:
		return leaderboard.LMSLeaderboardRecord{Player: p}, nil
	case Duel:
		return leaderboard.DuelLeaderboardRecord{Player: p}, nil
	default:
		return nil, errors.InvalidArgumentError{Message: "Unhandled game type '" + GameTypeToString(gt) + "'"}
	}
}
