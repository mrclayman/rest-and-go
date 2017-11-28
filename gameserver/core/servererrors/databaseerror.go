package servererrors

// DatabaseError indicates an issue
// with the backend database
type DatabaseError struct {
	Message string
}

// Error converts the error object
// into a string presentable to the client
func (err DatabaseError) Error() string {
	return "Database error: " + err.Message
}
