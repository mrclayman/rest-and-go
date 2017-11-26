package leaderboard

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
	"github.com/mrclayman/rest-and-go/gameclient/client/shared"
)

// DMLeaderboardRecord contains information
// on record in a DeathMatch-type
// leaderboard
type DMLeaderboardRecord struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
}

// DMLeaderboard is a slice of DM leaderboard records
type DMLeaderboard []DMLeaderboardRecord

// UnmarshalDMLeaderboard unmarshals a map of elements
// into an instance of the DMLeaderboardRecord
func UnmarshalDMLeaderboard(in []map[string]interface{}) (*DMLeaderboard, error) {
	retval := DMLeaderboard{}

	for _, lbRecIf := range in {
		lbRec, err := unmarshalDMLeaderboardRecord(lbRecIf)
		if err != nil {
			return nil, err
		}
		retval = append(retval, lbRec)
	}

	return &retval, nil
}

// unmarshalDMLeaderboardRecord unmarshals the contents
// of the input map into an instance of DMLeaderboardRecord
// structure and returns it
func unmarshalDMLeaderboardRecord(in map[string]interface{}) (DMLeaderboardRecord, error) {
	p, err := player.FromMap(in)
	if err != nil {
		return DMLeaderboardRecord{}, err
	}

	var kills, deaths uint
	kills, err = shared.KillCountFromMap(in)
	if err != nil {
		return DMLeaderboardRecord{}, err
	}

	deaths, err = shared.DeathCountFromMap(in)
	if err != nil {
		return DMLeaderboardRecord{}, err
	}

	return DMLeaderboardRecord{
		Player: p,
		Kills:  kills,
		Deaths: deaths,
	}, nil
}
