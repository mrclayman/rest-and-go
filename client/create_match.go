package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getGameType() string {
	fmt.Println("Available game types:")
	fmt.Println("1. Deathmatch")
	fmt.Println("2. Capture the Flag")
	fmt.Println("3. Last Man Standing")
	fmt.Println("4. Duel")
	fmt.Println()
	fmt.Print("Pick one by entering its number: ")

	var choice int
	for {
		_, err := fmt.Scanf("%v", &choice)
		if err != nil {
			fmt.Print("I need an integer. Try again: ")
			FlushStdin()
			continue
		}

		switch choice {
		case 1:
			return "dm"
		case 2:
			return "ctf"
		case 3:
			return "lms"
		case 4:
			return "duel"
		default:
			fmt.Print("The value must be in the interval 1 - 4. Enter another number: ")
		}
	}
}

// CreateMatch creates for the player a new match of the given type
func CreateMatch(c *http.Client, auth PlayerAuthData) error {

	gtype := getGameType()

	createData := map[string]interface{}{
		"player_id": auth.ID,
		"token":     auth.Token,
		"match_id":  0,
		"game_type": gtype,
	}

	var postData []byte
	var err error
	if postData, err = json.Marshal(createData); err != nil {
		return err
	}

	var sessionData MatchSessionData
	if err = post(c, "/match/join", postData, &sessionData); err != nil {
		return err
	}

	return runMatchLoop(c, auth, sessionData)
}
