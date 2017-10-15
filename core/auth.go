package core

import (
	"math/rand"
	"strconv"
)

// AuthToken defines the type of the
// authentication token issued by the
// server upon successful authentication
// of the connecting player
type AuthToken string

// InvalidAuthToken defines an invalid value
// for the authentication token of a player
const InvalidAuthToken AuthToken = ""

// GenerateAuthenticationToken generates a random number and
// casts it to a string
func GenerateAuthenticationToken() AuthToken {
	token := rand.Uint64()
	return AuthToken(strconv.FormatUint(token, 10))
}

// StringToAuthToken converts a string into
// the equivalent authentication token value
func StringToAuthToken(str string) AuthToken {
	return AuthToken(str)
}

// WebSockToken defines the type for player's
// WebSock token used to identify a player
// participating in an active match
type WebSockToken string

// InvalidWebSockToken defines an invalid
// value for a WebSock token
const InvalidWebSockToken WebSockToken = ""

// GenerateWebSockToken generates a new
// WebSock token value
func GenerateWebSockToken() WebSockToken {
	token := rand.Uint64()
	return WebSockToken(strconv.FormatUint(token, 10))
}
