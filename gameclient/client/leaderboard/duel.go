package leaderboard

import "github.com/mrclayman/rest-and-go/gameclient/client/player"

// DuelLeaderboardRecord contains information
// on record in a Duel-type leaderboard
type DuelLeaderboardRecord struct {
	Player player.Player `json:"player"`
	Kills  uint          `json:"kills"`
	Deaths uint          `json:"deaths"`
	Wins   uint          `json:"wins"`
}

// DuelLeaderboard is a slice of Duel-type leaderboard records
type DuelLeaderboard []DuelLeaderboardRecord

// UnmarshalDuelLeaderboard unmarshals a map of elements
// into an instance of the DuelLeaderboardRecord
/*func UnmarshalDuelLeaderboard(in []map[string]interface{}) (*DuelLeaderboard, error) {
	retval := DuelLeaderboard{}

	for _, lbRecIf := range in {
		lbRec, err := unmarshalDuelLeaderboardRecord(lbRecIf)
		if err != nil {
			return nil, err
		}
		retval = append(retval, lbRec)
	}

	return &retval, nil
}

// unmarshalDuelLeaderboardRecord unmarshals the contents
// of the input map into an instance of DuelLeaderboardRecord
// structure and returns it
func unmarshalDuelLeaderboardRecord(in map[string]interface{}) (DuelLeaderboardRecord, error) {
	p, err := player.FromMap(in)
	if err != nil {
		return DuelLeaderboardRecord{}, err
	}

	var kills, deaths, wins uint
	kills, err = shared.KillCountFromMap(in)
	if err != nil {
		return DuelLeaderboardRecord{}, err
	}

	deaths, err = shared.DeathCountFromMap(in)
	if err != nil {
		return DuelLeaderboardRecord{}, err
	}

	wins, err = shared.WinCountFromMap(in)
	if err != nil {
		return DuelLeaderboardRecord{}, err
	}

	return DuelLeaderboardRecord{
		Player: p,
		Kills:  kills,
		Deaths: deaths,
		Wins:   wins,
	}, nil
}
*/
