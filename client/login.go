package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func printLoginQuery() PlayerLogin {
	fmt.Println("You need to log in first. Pick the number of the identity you want to login as:")
	for i, player := range PlayerLogins {
		fmt.Printf("%v. %v\n", i+1, player.Nick)
	}

	choice := 0
	numPlayerLogins := len(PlayerLogins)
	for {
		fmt.Print("Who do you want to log in as? ")
		if _, err := fmt.Scanf("%v", &choice); err != nil || choice < 1 || choice > numPlayerLogins {
			fmt.Println("Please enter a number between 1 and", numPlayerLogins)
		} else {
			break
		}
	}
	return PlayerLogins[choice-1]
}

// Login asks the player which identity they wish to
// use for login, and logs the player in
func Login(c *http.Client) (PlayerLogin, PlayerAuthData, error) {
	player := printLoginQuery()
	loginJSON, err := json.Marshal(player)

	if err != nil {
		return PlayerLogin{}, PlayerAuthData{}, err
	}

	var authData PlayerAuthData
	err = post(c, "/login", loginJSON, &authData)
	if err != nil {
		return PlayerLogin{}, PlayerAuthData{}, err
	}

	return player, authData, nil
}
