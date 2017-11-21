package client

// DMMatch contains information on a
// DeathMatch type match
type DMMatch struct {
	ID    uint64        `json:"match_id"`
	Type  string        `json:"match_type"`
	Ranks DMPlayerRanks `json:"ranks"`
}

// CTFMatch contains information on
// a CTF type match
type CTFMatch struct {
	ID    uint64         `json:"match_id"`
	Type  string         `json:"match_type"`
	Ranks CTFPlayerRanks `json:"ranks"`
}

// LMSMatch contains information on
// a LMS type match
type LMSMatch struct {
	ID    uint64         `json:"match_id"`
	Type  string         `json:"match_type"`
	Ranks LMSPlayerRanks `json:"ranks"`
}

// DuelMatch contains information on
// a Duel type match
type DuelMatch LMSMatch

// Matchlist defines a slice of match instances
type Matchlist []interface{}

// unmarshalMatchlist
func unmarshalMatchlist(in []map[string]interface{}) Matchlist {

}
