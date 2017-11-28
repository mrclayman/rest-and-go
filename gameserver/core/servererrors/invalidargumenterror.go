package servererrors

// InvalidArgumentError indicates an
// invalid input to a function
type InvalidArgumentError struct {
	Message string
}

func (err InvalidArgumentError) Error() string {
	return "Invalid argument error: " + err.Message
}
