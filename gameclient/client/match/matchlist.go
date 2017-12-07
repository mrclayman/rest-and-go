package match

import (
	"encoding/json"
)

// Matchlist defines a structure that holds
// lists of existing matches for all game types
type Matchlist struct {
	DM   DMMatches   `json:"dm"`
	CTF  CTFMatches  `json:"ctf"`
	LMS  LMSMatches  `json:"lms"`
	Duel DuelMatches `json:"duel"`
}

// MatchlistUnmarshaler unmarshals the contents
// of the Matchlist structure from JSON byte slice
type MatchlistUnmarshaler struct {
	Matchlist Matchlist
}

// UnmarshalJSON unmarshals the contents of a server's
// response into the internal structure of match lists
func (m *MatchlistUnmarshaler) UnmarshalJSON(in []byte) error {
	return json.Unmarshal(in, &m.Matchlist)
	// Read opening curly brace
	/*	var t json.Token
		var err error
		if t, err = d.Token(); err != nil {
			return err
		}

		for d.More() {
			var gt string
			var ok bool

			// Get the game type identifier
			t, err = d.Token()
			gt, ok = t.(string)
			if !ok {
				return errors.New("Token not a string (game type)")
			}
			fmt.Printf("Processing match list for type %v\n", gt)

					if err = d.Decode(&gt); err != nil {
					fmt.Println("Failed to read game type identifier")
					return err
				}

			switch gt {
			case DeathMatch:
				err = m.unmarshalDMMatchlist(d)
			case CaptureTheFlag:
				err = m.unmarshalCTFMatchlist(d)
			case LastManStanding:
				err = m.unmarshalLMSMatchlist(d)
			case Duel:
				err = m.unmarshalDuelMatchlist(d)
			}

			if err != nil {
				return err
			}
		}

		return nil
	*/
}

/*
func (m *Matchlist) unmarshalDMMatchlist(d *json.Decoder) error {
	matches := make(DMMatches, 0, 5)
	if err := d.Decode(&matches); err != nil {
		return err
	}

	m.DM = matches
	return nil
}

func (m *Matchlist) unmarshalCTFMatchlist(d *json.Decoder) error {
	matches := make(CTFMatches, 0, 5)
	if err := d.Decode(&matches); err != nil {
		return err
	}

	m.CTF = matches
	return nil
}

func (m *Matchlist) unmarshalLMSMatchlist(d *json.Decoder) error {
	matches := make(LMSMatches, 0, 5)
	if err := d.Decode(&matches); err != nil {
		return err
	}

	m.LMS = matches
	return nil
}

func (m *Matchlist) unmarshalDuelMatchlist(d *json.Decoder) error {
	matches := make(DuelMatches, 0, 5)
	if err := d.Decode(&matches); err != nil {
		return err
	}

	m.Duel = matches
	return nil
}

/*
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
*/
