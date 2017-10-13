package core

import (
	"math/rand"
	"strconv"
)

// GenerateAuthenticationToken generates a random number and
// casts it to a string
func GenerateAuthenticationToken() string {
	token := rand.Uint64()
	return strconv.FormatUint(token, 10)
}
