package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JoinMatch queries the server for active matches and
// lets the player choose a match to join
func JoinMatch(c *http.Client, auth PlayerAuthData) error {
	var err error

	if err = ListMatches(c, auth); err != nil {
		fmt.Println("Cannot join due to previous errors")
		return err
	}

	var matchID uint64
	for {
		fmt.Print("Enter the ID of the match you wish to join: ")
		_, err = fmt.Scanf("%v", &matchID)
		if err != nil {
			fmt.Println("I need an integer")
		} else {
			break
		}
	}

	joinData := map[string]interface{}{
		"player_id":    auth.ID,
		"player_token": auth.Token,
		"match_id":     matchID,
		"game_type":    "",
	}

	var postData []byte
	postData, err = json.Marshal(joinData)
	if err != nil {
		return err
	}

	var sessionData MatchSessionData
	err = post(c, "/match/join", postData, &sessionData)
	if err != nil {
		return err
	}

	return enterMatchLoop(c, &sessionData)
}
