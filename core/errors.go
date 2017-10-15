package core

// IntegrityError is used when data integrity
// violation has been encountered within the server
type IntegrityError struct {
	Message string
}

// Error formats the IntegrityError's instance
// message
func (err IntegrityError) Error() string {
	return "Integrity error:" + err.Message
}

// MatchNotFoundError indicates that a match
// with the given ID could not be found
type MatchNotFoundError struct {
	MID MatchID
}

// Error converts the error object into
// a string presentable to the user
func (err MatchNotFoundError) Error() string {
	return "Match not found: " + MatchIDToString(err.MID)
}
