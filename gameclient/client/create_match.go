package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrclayman/rest-and-go/gameclient/client/match"
	"github.com/mrclayman/rest-and-go/gameclient/client/net"
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
			return match.DeathMatch
		case 2:
			return match.CaptureTheFlag
		case 3:
			return match.LastManStanding
		case 4:
			return match.Duel
		default:
			fmt.Print("The value must be in the interval 1 - 4. Enter another number: ")
		}
	}
}

// CreateMatch creates for the player a new match of the given type
func CreateMatch(c *http.Client, ps net.PlayerSession) error {

	gtype := getGameType()

	createData := map[string]interface{}{
		"player_id": ps.ID,
		"token":     ps.Token,
		"match":     match.ID{Number: match.InvalidNumber, Type: gtype},
	}

	var postData []byte
	var err error
	if postData, err = json.Marshal(createData); err != nil {
		return err
	}

	var ms net.MatchSession
	if err = net.Post(c, "/match/join", postData, &ms); err != nil {
		return err
	}

	return runMatchLoop(c, ps, ms)
}
