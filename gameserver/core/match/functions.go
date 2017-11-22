package match

import (
	"math/rand"
	"strconv"
)

// GenerateNumber generates a random match ID
// TODO Hide once the dummy data used in the server are no longer needed
func GenerateNumber() Number {
	return Number(rand.Uint64())
}

// NumberToString converts a match
// number into a string
func NumberToString(ID Number) string {
	return strconv.FormatUint(uint64(ID), 10)
}
