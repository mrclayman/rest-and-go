package core

import (
	"math/rand"
	"strconv"
)

// GenerateAuthenticationToken generates a random number and
// casts it to a string
func GenerateAuthenticationToken() AuthToken {
	token := rand.Uint64()
	return AuthToken(strconv.FormatUint(token, 10))
}
