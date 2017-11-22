package match

// Registry aggregates ongoing matches
// of all types
type Registry struct {
	DM   DMMatches
	CTF  CTFMatches
	LMS  LMSMatches
	Duel DuelMatches
}
