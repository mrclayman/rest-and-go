package match

// Number defines the type that holds
// the value of the number of the match
type Number uint64

// InvalidNumber defines the value for
// an invalid match number
const InvalidNumber Number = 0

// ID uniquely identifies a match as a
// combination of number and game type
type ID struct {
	Number Number `json:"id"`
	Type   string `json:"type"`
}
