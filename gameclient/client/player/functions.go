package player

// FromMap retrieves player info under the key "player"
// from the map in the argument and synthesizes an instance
// of the Player structure from them
/*func FromMap(rankRec map[string]interface{}) (Player, error) {
	var pm map[string]interface{}

	pIf, ok := rankRec["player"]
	if !ok {
		return Player{}, errors.New("Map does not seem to have player info")
	} else if pm, ok = pIf.(map[string]interface{}); !ok {
		return Player{}, errors.New("Player info is not a map")
	} else if IDIf, ok := pm["id"]; !ok {
		return Player{}, errors.New("Player info map does not have player ID")
	} else if IDNum, ok := IDIf.(json.Number); !ok {
		return Player{}, errors.New("Player ID is not a number")
	} else if ID, err := IDNum.Int64(); err != nil {
		return Player{}, err
	} else if nickIf, ok := pm["nick"]; !ok {
		return Player{}, errors.New("Player info map does not have player nickname")
	} else if nick, ok := nickIf.(string); !ok {
		return Player{}, errors.New("Player nickname is not a string")
	} else {
		return Player{ID: int(ID), Nick: nick}, nil
	}
}*/
