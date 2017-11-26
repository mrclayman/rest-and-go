package match

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

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

// IDFromMap retrieves an element with key "match_id"
// from the map in the argument and verifies that the
// element's type is an uint64.
func IDFromMap(m map[string]interface{}) (uint64, error) {
	if v, ok := m["match_id"]; !ok {
		return 0, errors.New("Match does not seem to have 'match_id' key")
	} else if IDNum, ok := v.(json.Number); !ok {
		return 0, errors.New("Match ID does not seem to be a number, but " + reflect.TypeOf(v).Name())
	} else {
		// I have to use ParseUint, because overflows may happen
		// when Number.Int64() is called on an unsigned value
		// sent by the server
		ID, err := strconv.ParseUint(IDNum.String(), 10, 64)
		if err != nil {
			return 0, err
		}
		return ID, nil
	}
}
