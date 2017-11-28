package servererrors

// LogicError indicates an illogical problem
// occurring within the system
type LogicError struct {
	Message string
}

// Error converts the error object into
// a string presentable to the client
func (err LogicError) Error() string {
	return "Logic error: " + err.Message
}
