package shared

import (
	"encoding/json"
	"errors"
	"reflect"
)

// RanksFromMap retrieves an element with key "ranks"
// from the map in the argument and verifies that the
// element's type is a map of string-interface{} pairs.
func RanksFromMap(m map[string]interface{}) (map[string]interface{}, error) {
	if r, ok := m["ranks"]; !ok {
		return nil, errors.New("Match does not seem to have a list of ranks")
	} else if rm, ok := r.(map[string]interface{}); !ok {
		return nil, errors.New("Ranks do not appear to be a list of values, but instead " + reflect.TypeOf(rm).Name())
	} else {
		return rm, nil
	}
}

// KillCountFromMap retrieves an element with key "kills"
// from the map in the argument and verifies that the
// element's type is an uint.
func KillCountFromMap(rm map[string]interface{}) (uint, error) {
	if kcIf, ok := rm["kills"]; !ok {
		return 0, errors.New("Kill count not in rank map")
	} else if kcNum, ok := kcIf.(json.Number); !ok {
		return 0, errors.New("Kill count not a number")
	} else {
		kc, err := kcNum.Int64()
		if err != nil {
			return 0, err
		}
		return uint(kc), nil
	}
}

// DeathCountFromMap retrieves an element with key "deaths"
// from the map in the argument and verifies that the
// element's type is an uint.
func DeathCountFromMap(rm map[string]interface{}) (uint, error) {
	if dcIf, ok := rm["deaths"]; !ok {
		return 0, errors.New("Death count not in rank map")
	} else if dcNum, ok := dcIf.(json.Number); !ok {
		return 0, errors.New("Death count not a number")
	} else {
		dc, err := dcNum.Int64()
		if err != nil {
			return 0, err
		}
		return uint(dc), nil
	}
}

// CaptureCountFromMap retrieves an element with key "captures"
// from the map in the argument and verifies that the
// element's type is an uint.
func CaptureCountFromMap(rm map[string]interface{}) (uint, error) {
	if ccIf, ok := rm["captures"]; !ok {
		return 0, errors.New("Capture count not in rank map")
	} else if ccNum, ok := ccIf.(json.Number); !ok {
		return 0, errors.New("Capture count not a number")
	} else {
		cc, err := ccNum.Int64()
		if err != nil {
			return 0, err
		}
		return uint(cc), nil
	}
}

// WinCountFromMap retrieves an element with key "wins"
// from the map in the argument and verifies that the
// element's type is an uint64.
func WinCountFromMap(rm map[string]interface{}) (uint, error) {
	if wcIf, ok := rm["wins"]; !ok {
		return 0, errors.New("Win count not in rank map")
	} else if wcNum, ok := wcIf.(json.Number); !ok {
		return 0, errors.New("Win count not a number")
	} else {
		wc, err := wcNum.Int64()
		if err != nil {
			return 0, err
		}
		return uint(wc), nil
	}
}
