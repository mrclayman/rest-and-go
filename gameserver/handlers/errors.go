package handlers

// RequestError defines an error type used
// when the request is invalid or otherwise
// malformed
type RequestError struct {
	Message string
}

// Error provides the means of getting
// the error message from the RequestError
// structure instance
func (err RequestError) Error() string {
	return "Request error: " + err.Message
}

// InternalError indicates that an internal
// server error has been encountered and
// request processing is no longer possible
type InternalError struct {
	Message string
}

// Error provides the means of getting
// the error message from InternalError
// structure instance
func (err InternalError) Error() string {
	return "Internal server error: " + err.Message
}
