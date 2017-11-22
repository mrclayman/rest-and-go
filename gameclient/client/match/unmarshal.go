package match

import (
	"encoding/json"
	"errors"

	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

type matchlistUnmarshaler struct {
	Matchlist Matchlist
}

// UnmarshalJSON unmarshals the contents
// of the byte slice and synthesizes a slice of
// appropriate *Match instances
func (mu matchlistUnmarshaler) UnmarshalJSON(in []byte) error {
	var jsonData []map[string]interface{}

	if err := json.Unmarshal(in, jsonData); err != nil {
		return err
	}

	for _, m := range jsonData {
		if mID, err := getMatchID(m); err != nil {
			return err
		} else if mType, err := getMatchType(m); err != nil {
			return err
		} else {
			switch mType {
			case DeathMatch:
				err = mu.unmarshalDMMatch(m, mID, mType)
			case CaptureTheFlag:
				err = mu.unmarshalCTFMatch(m, mID, mType)
			case LastManStanding:
				err = mu.unmarshalLMSMatch(m, mID, mType)
			case Duel:
				err = mu.unmarshalDuelMatch(m, mID, mType)
			default:
				return errors.New("Unhandled game type '" + mType + "' in matchlist unmarshaler")
			}

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (mu matchlistUnmarshaler) unmarshalDMMatch(in map[string]interface{}, mID uint64, mType string) error {

	//var ranks DMPlayerRanks
	rl, err := getRanks(in)
	if err != nil {
		return err
	}

	for _, i := range rl {

	}
	return nil
}

func (mu matchlistUnmarshaler) unmarshalCTFMatch(in map[string]interface{}, mID uint64, mType string) error {

	return nil
}

func (mu matchlistUnmarshaler) unmarshalLMSMatch(in map[string]interface{}, mID uint64, mType string) error {

	return nil
}

func (mu matchlistUnmarshaler) unmarshalDuelMatch(in map[string]interface{}, mID uint64, mType string) error {

	return nil
}

func getMatchID(m map[string]interface{}) (uint64, error) {
	if v, ok := m["match_id"]; !ok {
		return 0, errors.New("Match does not seem to have 'match_id' key")
	} else if ID, ok := v.(uint64); !ok {
		return 0, errors.New("Match ID does not seem to be a uint64")
	} else {
		return ID, nil
	}
}

func getMatchType(m map[string]interface{}) (string, error) {
	if v, ok := m["match_type"]; !ok {
		return "", errors.New("Match does not seem to have 'match_type' key")
	} else if t, ok := v.(string); !ok {
		return "", errors.New("Match type does not seem to be a string")
	} else {
		return t, nil
	}
}

func getRanks(m map[string]interface{}) ([]map[string]interface{}, error) {
	if r, ok := m["ranks"]; !ok {
		return nil, errors.New("Match does not seem to have a list of ranks")
	} else if v, ok := r.([]map[string]interface{}); !ok {
		return nil, errors.New("Ranks do not appear to be a list of structures")
	} else {
		return v, nil
	}
}

func getPlayerFromRankMap(rm map[string]interface{}) (player.Player, error) {
	p, ok := rm["player"]
	if !ok {
		return player.Player{}, errors.New("Rank structure does not seem to have player info")
	} else if pm, ok := p.(map[string]interface{}); !ok {
		return player.Player{}, errors.New("Player info is not a map")
	} else if pIDIf, ok := pm["player_id"]; !ok {
		return player.Player{}, errors.New("Player info map does not have player ID")
	} else if pID, ok := pIDIf.(int); !ok {
		return player.Player{}, errors.New("Player ID is not an int")
	} else if pNameIf, ok := pm["player_name"]; !ok {
		return player.Player{}, errors.New("Player info map does not have player name")
	} else if pName, ok := pNameIf.(string); !ok {
		return player.Player{}, errors.New("Player name is not a string")
	} else {
		return player.Player{ID: pID, Name: pName}, nil
	}
}

func getKillCountFromRankMap(rm map[string]interface{}) (uint, error) {
	if kcIf, ok := rm["kills"]; !ok {
		return 0, errors.New("Kill count not in rank map")
	} else if kc, ok := kcIf.(uint); !ok {
		return 0, errors.New("Kill count not an uint")
	} else {
		return kc, nil
	}
}

func getDeathCountFromRankMap(rm map[string]interface{}) (uint, error) {
	if dcIf, ok := rm["deaths"]; !ok {
		return 0, errors.New("Death count not in rank map")
	} else if dc, ok := dcIf.(uint); !ok {
		return 0, errors.New("Death count not an uint")
	} else {
		return dc, nil
	}
}

func getInt(m map[string]interface{}, key string) (int64, error) {
	if vIf, ok := m[key]; !ok {
		return 0, errors.New("Map does not have item with key '" + key + "'")
	} else if v, ok :=
}
