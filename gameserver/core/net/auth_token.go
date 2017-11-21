package net

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

// GetRandomString returns a random string
// of the given length comprised of alphanumeric
// characters
func GetRandomString(length int) string {
	const chars = "ABCDE0FGHIJ1KLMNO2PQRST3UVWXY4Zabcd5efghi6jklmn7opqrs8tuvwx9yz"
	const charCount = int64(len(chars))
	retval := make([]byte, length)

	for i := 0; i < length; i++ {
		retval[i] = chars[rand.Int63()%charCount]
	}

	return string(retval)
}
