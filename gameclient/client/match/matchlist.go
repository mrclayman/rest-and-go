package match

// Matchlist defines a structure that holds
// lists of existing matches for all game types
type Matchlist struct {
	DM   DMMatches   `json:"dm"`
	CTF  CTFMatches  `json:"ctf"`
	LMS  LMSMatches  `json:"lms"`
	Duel DuelMatches `json:"duel"`
}
