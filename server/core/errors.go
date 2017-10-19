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

// InvalidArgumentError indicates an
// invalid input to a function
type InvalidArgumentError struct {
	Message string
}

func (err InvalidArgumentError) Error() string {
	return "Invalid argument error: " + err.Message
}

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
