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
