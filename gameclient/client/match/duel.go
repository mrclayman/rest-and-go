package match

import (
	"errors"

	"github.com/mrclayman/rest-and-go/gameclient/client/player"
	"github.com/mrclayman/rest-and-go/gameclient/client/shared"
)

// DuelPlayerRank defines the structure of
// a player's rank in the Duel game type
type DuelPlayerRank struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// DuelPlayerRanks defines the type for a slice
// of Duel game type player rank objects
type DuelPlayerRanks []DuelPlayerRank

// DuelMatch contains information on
// a Duel type match
type DuelMatch struct {
	Number Number          `json:"match_id"`
	Ranks  DuelPlayerRanks `json:"ranks"`
}

// DuelMatches defines a slice of Duel-type matches
type DuelMatches []*DuelMatch

// unmarshalDuelMatch unmarshals the contents of the
// input map into an instance of DuelMatch
func unmarshalDuelMatch(in map[string]interface{}) (*DuelMatch, error) {
	mID, err := IDFromMap(in)
	if err != nil {
		return nil, err
	}

	var playerRanks map[string]interface{}
	playerRanks, err = shared.RanksFromMap(in)
	if err != nil {
		return nil, err
	}

	var ranks DuelPlayerRanks
	for _, r := range playerRanks {
		rm, ok := r.(map[string]interface{})
		if !ok {
			return nil, errors.New("Rank record not a map of values")
		}

		rank, err := unmarshalDuelRankRecord(rm)
		if err != nil {
			return nil, err
		}

		ranks = append(ranks, rank)
	}

	return &DuelMatch{
		Number: Number(mID),
		Ranks:  ranks,
	}, nil
}

// unmarshalDuelRankRecord unmarshals the contents
// of the input map into an instance of
// the DuelPlayerRank structure
func unmarshalDuelRankRecord(in map[string]interface{}) (DuelPlayerRank, error) {
	player, err := player.FromMap(in)
	if err != nil {
		return DuelPlayerRank{}, err
	}

	var kills uint
	kills, err = shared.KillCountFromMap(in)
	if err != nil {
		return DuelPlayerRank{}, err
	}

	var deaths uint
	deaths, err = shared.DeathCountFromMap(in)
	if err != nil {
		return DuelPlayerRank{}, err
	}

	return DuelPlayerRank{
		Player: player,
		Kills:  kills,
		Deaths: deaths,
	}, nil
}
