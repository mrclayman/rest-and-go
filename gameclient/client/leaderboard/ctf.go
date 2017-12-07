package leaderboard

import (
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

// CTFLeaderboardRecord contains information
// on record in a CTF-type leaderboard
type CTFLeaderboardRecord struct {
	Player   player.Player `json:"player"`
	Kills    uint          `json:"kills"`
	Deaths   uint          `json:"deaths"`
	Captures uint          `json:"captures"`
}

// CTFLeaderboard is a slice of CTF leaderboard records
type CTFLeaderboard []CTFLeaderboardRecord

// UnmarshalCTFLeaderboard unmarshals a map of elements
// into an instance of the CTFLeaderboardRecord
/*func UnmarshalCTFLeaderboard(in []map[string]interface{}) (*CTFLeaderboard, error) {
	retval := CTFLeaderboard{}

	for _, lbRecIf := range in {
		lbRec, err := unmarshalCTFLeaderboardRecord(lbRecIf)
		if err != nil {
			return nil, err
		}
		retval = append(retval, lbRec)
	}

	return &retval, nil
}

// unmarshalCTFLeaderboardRecord unmarshals the contents
// of the input map into an instance of CTFLeaderboardRecord
// structure and returns it
func unmarshalCTFLeaderboardRecord(in map[string]interface{}) (CTFLeaderboardRecord, error) {
	p, err := player.FromMap(in)
	if err != nil {
		return CTFLeaderboardRecord{}, err
	}

	var kills, deaths, captures uint
	kills, err = shared.KillCountFromMap(in)
	if err != nil {
		return CTFLeaderboardRecord{}, err
	}

	deaths, err = shared.DeathCountFromMap(in)
	if err != nil {
		return CTFLeaderboardRecord{}, err
	}

	captures, err = shared.CaptureCountFromMap(in)
	if err != nil {
		return CTFLeaderboardRecord{}, err
	}

	return CTFLeaderboardRecord{
		Player:   p,
		Kills:    kills,
		Deaths:   deaths,
		Captures: captures,
	}, nil
}
*/
