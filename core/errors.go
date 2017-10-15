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
