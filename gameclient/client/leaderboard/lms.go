package leaderboard

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
	"github.com/mrclayman/rest-and-go/gameclient/client/shared"
)

// LMSLeaderboardRecord contains information
// on record in a LMS-type leaderboard
type LMSLeaderboardRecord struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
	Wins   uint          `json:"wins"`
}

// LMSLeaderboard is a slice of LMS-type leaderboard records
type LMSLeaderboard []LMSLeaderboardRecord

// UnmarshalLMSLeaderboard unmarshals a map of elements
// into an instance of the LMSLeaderboardRecord
func UnmarshalLMSLeaderboard(in []map[string]interface{}) (*LMSLeaderboard, error) {
	retval := LMSLeaderboard{}

	for _, lbRecIf := range in {
		lbRec, err := unmarshalLMSLeaderboardRecord(lbRecIf)
		if err != nil {
			return nil, err
		}
		retval = append(retval, lbRec)
	}

	return &retval, nil
}

// unmarshalLMSLeaderboardRecord unmarshals the contents
// of the input map into an instance of LMSLeaderboardRecord
// structure and returns it
func unmarshalLMSLeaderboardRecord(in map[string]interface{}) (LMSLeaderboardRecord, error) {
	p, err := player.FromMap(in)
	if err != nil {
		return LMSLeaderboardRecord{}, err
	}

	var kills, deaths, wins uint
	kills, err = shared.KillCountFromMap(in)
	if err != nil {
		return LMSLeaderboardRecord{}, err
	}

	deaths, err = shared.DeathCountFromMap(in)
	if err != nil {
		return LMSLeaderboardRecord{}, err
	}

	wins, err = shared.WinCountFromMap(in)
	if err != nil {
		return LMSLeaderboardRecord{}, err
	}

	return LMSLeaderboardRecord{
		Player: p,
		Kills:  kills,
		Deaths: deaths,
		Wins:   wins,
	}, nil
}
