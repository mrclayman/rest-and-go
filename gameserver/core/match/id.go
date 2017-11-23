package match

// ID structure represents an ongoing match
// and the participating players
type ID struct {
	Number Number   `json:"id"`
	Type   GameType `json:"type"`
}

// InvalidID defines an invalid ID value
var InvalidID ID
