package match

import (
	"errors"
	"reflect"

	"github.com/mrclayman/rest-and-go/gameclient/client/shared"
)

// Matchlist defines a structure that holds
// lists of existing matches for all game types
type Matchlist struct {
	DM   DMMatches
	CTF  CTFMatches
	LMS  LMSMatches
	Duel DuelMatches
}

// UnmarshalJSON unmarshals input byteslice into
// a map structure, from which individual match
// type instances are synthesized
func (m *Matchlist) UnmarshalJSON(in []byte) error {
	// TODO Maybe I should do piecemeal unmarshaling
	// That way I should be able to leave the unmarshaling
	// to the sub-structures themselves.
	inMap := make(map[string]interface{})
	err := shared.DecodeJSON(in, &inMap)
	if err != nil {
		return err
	}

	for k, v := range inMap {
		// fmt.Printf("Key-value: K: %v, V: %v\n", k, v)
		if reflect.TypeOf(v).Kind() == reflect.Slice && len(v.([]interface{})) == 0 {
			continue
		}

		// fmt.Println("Type of non-empty v:", reflect.TypeOf(v))
		vMapSlice, ok := v.([]interface{})
		if !ok {
			return errors.New("Item in game type matchlist not a slice of maps")
		}

		switch k {
		case DeathMatch:
			err = m.unmarshalDMMatches(vMapSlice)
		case CaptureTheFlag:
			err = m.unmarshalCTFMatches(vMapSlice)
		case LastManStanding:
			err = m.unmarshalLMSMatches(vMapSlice)
		case Duel:
			err = m.unmarshalDuelMatches(vMapSlice)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Matchlist) unmarshalDMMatches(in []interface{}) error {
	for _, mMapIf := range in {
		mMap, ok := mMapIf.(map[string]interface{})
		if !ok {
			return errors.New("Item in slice not a map of values")
		}

		match, err := unmarshalDMMatch(mMap)
		if err != nil {
			return err
		}
		m.DM = append(m.DM, match)
	}
	return nil
}

func (m *Matchlist) unmarshalCTFMatches(in []interface{}) error {
	for _, mMapIf := range in {
		mMap, ok := mMapIf.(map[string]interface{})
		if !ok {
			return errors.New("Item in slice not a map of values")
		}

		match, err := unmarshalCTFMatch(mMap)
		if err != nil {
			return err
		}
		m.CTF = append(m.CTF, match)
	}
	return nil
}

func (m *Matchlist) unmarshalLMSMatches(in []interface{}) error {
	for _, mMapIf := range in {
		mMap, ok := mMapIf.(map[string]interface{})
		if !ok {
			return errors.New("Item in slice not a map of values")
		}

		match, err := unmarshalLMSMatch(mMap)
		if err != nil {
			return err
		}
		m.LMS = append(m.LMS, match)
	}
	return nil
}

func (m *Matchlist) unmarshalDuelMatches(in []interface{}) error {
	for _, mMapIf := range in {
		mMap, ok := mMapIf.(map[string]interface{})
		if !ok {
			return errors.New("Item in slice not a map of values")
		}

		match, err := unmarshalDuelMatch(mMap)
		if err != nil {
			return err
		}
		m.Duel = append(m.Duel, match)
	}
	return nil
}
