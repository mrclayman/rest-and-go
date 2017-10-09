package errors

/* RequestError defines an error type used
 * when the request is invalid or otherwise
 * malformed
 */
type RequestError struct {
	Message string
}

func (err RequestError) Error() string {
	return "Request error:" + err.Message
}
