package match

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

// LMSPlayerRank defines the structure of a player's
// rank in the LMS game type
type LMSPlayerRank struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// LMSPlayerRanks is a type for a slice of
// LMS game type player rank objects
type LMSPlayerRanks map[player.ID]LMSPlayerRank

// LMSMatch contains information on
// a LMS type match
type LMSMatch struct {
	Number Number         `json:"match_id"`
	Ranks  LMSPlayerRanks `json:"ranks"`
}

// LMSMatches defines a slice of LMS-type matches
type LMSMatches []*LMSMatch

// unmarshalLMSMatch unmarshals the contents of the
// input map into an instance of LMSMatch
/*func unmarshalLMSMatch(in map[string]interface{}) (*LMSMatch, error) {
	mID, err := IDFromMap(in)
	if err != nil {
		return nil, err
	}

	var playerRanks map[string]interface{}
	playerRanks, err = shared.RanksFromMap(in)
	if err != nil {
		return nil, err
	}

	var ranks LMSPlayerRanks
	for _, r := range playerRanks {
		rm, ok := r.(map[string]interface{})
		if !ok {
			return nil, errors.New("Rank record not a map of values")
		}

		rank, err := unmarshalLMSRankRecord(rm)
		if err != nil {
			return nil, err
		}

		ranks = append(ranks, rank)
	}

	return &LMSMatch{
		Number: Number(mID),
		Ranks:  ranks,
	}, nil
}

// unmarshalLMSRankRecord unmarshals the contents
// of the input map into an instance of
// the LMSPlayerRank structure
func unmarshalLMSRankRecord(in map[string]interface{}) (LMSPlayerRank, error) {
	player, err := player.FromMap(in)
	if err != nil {
		return LMSPlayerRank{}, err
	}

	var kills uint
	kills, err = shared.KillCountFromMap(in)
	if err != nil {
		return LMSPlayerRank{}, err
	}

	var deaths uint
	deaths, err = shared.DeathCountFromMap(in)
	if err != nil {
		return LMSPlayerRank{}, err
	}

	return LMSPlayerRank{
		Player: player,
		Kills:  kills,
		Deaths: deaths,
	}, nil
}
*/
