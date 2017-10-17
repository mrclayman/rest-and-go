package core

// WebSocketToken defines the type for player's
// WebSock token used to identify a player
// participating in an active match
type WebSocketToken string

const (
	// WebSocketTokenLength denotes the length
	// of the WebSocket communication token
	WebSocketTokenLength = 32

	// InvalidWebSocketToken defines an invalid
	// value for a WebSock token
	InvalidWebSocketToken WebSocketToken = ""
)

// GenerateWebSocketToken generates a new
// WebSock token value
func GenerateWebSocketToken() WebSocketToken {
	return WebSocketToken(GetRandomString(WebSocketTokenLength))
}
