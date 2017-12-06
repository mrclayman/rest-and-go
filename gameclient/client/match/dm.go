package match

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

// DMPlayerRank aggregates information on a player's
// rank in a match
type DMPlayerRank struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// DMPlayerRanks defines a list of players' ranks
type DMPlayerRanks map[player.ID]DMPlayerRank

// DMMatch contains information on a
// DeathMatch type match
type DMMatch struct {
	Number Number        `json:"match_id"`
	Ranks  DMPlayerRanks `json:"ranks"`
}

// DMMatches defines a slice of DeathMatch-type matches
type DMMatches []*DMMatch

// unmarshalDMMatch unmarshals the contents of the
// input map into an instance of DMMatch
/*func unmarshalDMMatch(in map[string]interface{}) (*DMMatch, error) {
	mID, err := IDFromMap(in)
	if err != nil {
		return nil, err
	}

	var playerRanks map[string]interface{}
	playerRanks, err = shared.RanksFromMap(in)
	if err != nil {
		return nil, err
	}

	var ranks DMPlayerRanks
	for _, r := range playerRanks {
		rm, ok := r.(map[string]interface{})
		if !ok {
			return nil, errors.New("Rank record not a map of values")
		}

		rank, err := unmarshalDMRankRecord(rm)
		if err != nil {
			return nil, err
		}

		ranks = append(ranks, rank)
	}

	return &DMMatch{
		Number: Number(mID),
		Ranks:  ranks,
	}, nil
}

// unmarshalDMRankRecord unmarshals the contents
// of the input map into an instance of
// the DMPlayerRank structure
func unmarshalDMRankRecord(in map[string]interface{}) (DMPlayerRank, error) {
	player, err := player.FromMap(in)
	if err != nil {
		return DMPlayerRank{}, err
	}

	var kills uint
	kills, err = shared.KillCountFromMap(in)
	if err != nil {
		return DMPlayerRank{}, err
	}

	var deaths uint
	deaths, err = shared.DeathCountFromMap(in)
	if err != nil {
		return DMPlayerRank{}, err
	}

	return DMPlayerRank{
		Player: player,
		Kills:  kills,
		Deaths: deaths,
	}, nil
}
*/
