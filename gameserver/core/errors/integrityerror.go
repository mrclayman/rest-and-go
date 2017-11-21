package errors

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
