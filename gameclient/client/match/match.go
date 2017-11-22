package match

const (
	// DeathMatch indicates the match is of
	// type  "Deathmatch"
	DeathMatch string = "dm"

	// CaptureTheFlag indicates the match is
	// of type "Capture the Flag"
	CaptureTheFlag string = "ctf"

	// LastManStanding indicates the match is
	// of type "Last Man Standing"
	LastManStanding string = "lms"

	// Duel indicates the match is of type "Duel"
	Duel string = "duel"
)

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
