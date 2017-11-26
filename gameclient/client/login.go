package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrclayman/rest-and-go/gameclient/client/net"
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
)

func printLoginQuery() player.Login {
	fmt.Println("You need to log in first. Pick the number of the identity you want to login as:")
	for i, player := range Logins {
		fmt.Printf("%v. %v\n", i+1, player.Nick)
	}

	choice := 0
	numPlayerLogins := len(Logins)
	for {
		fmt.Print("Who do you want to log in as? ")
		if _, err := fmt.Scanf("%v", &choice); err != nil || choice < 1 || choice > numPlayerLogins {
			fmt.Println("Please enter a number between 1 and", numPlayerLogins)
			if err != nil {
				FlushStdin()
			}
		} else {
			break
		}
	}
	return Logins[choice-1]
}

// Login asks the player which identity they wish to
// use for login, and logs the player in
func Login(c *http.Client) (player.Login, net.PlayerSession, error) {
	p := printLoginQuery()
	loginJSON, err := json.Marshal(p)

	if err != nil {
		return player.Login{}, net.PlayerSession{}, err
	}

	var ps net.PlayerSession
	err = net.Post(c, "/login", loginJSON, &ps)
	if err != nil {
		return player.Login{}, net.PlayerSession{}, err
	}

	return p, ps, nil
}
