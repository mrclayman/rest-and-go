package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrclayman/rest-and-go/gameclient/client/match"
	"github.com/mrclayman/rest-and-go/gameclient/client/net"
	"github.com/mrclayman/rest-and-go/gameclient/client/net/ws"
)

func getUserAction() uint16 {
	fmt.Println("Pick next action:")
	fmt.Println("1. Fire a weapon")
	fmt.Println("2. Move to a new position")
	fmt.Println("3. Get a list of players")
	fmt.Println("4. Quit the match")
	fmt.Println("------")

	var choice uint16
	for {
		fmt.Print("What do you want to do next? ")
		_, err := fmt.Scanf("%v", &choice)

		if err != nil || choice < 1 || choice > 4 {
			fmt.Println("Please enter an integer value in the interval between 1 and 4")
			if err != nil {
				FlushStdin()
			}
		} else {
			break
		}
	}

	return choice
}

func printDMPlayerList(data []byte) {
	var m match.DMMatch
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Println(err.Error())
	} else {
		printDMMatch(&m)
	}
}

func printCTFPlayerList(data []byte) {
	var m match.CTFMatch
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Println(err.Error())
	} else {
		printCTFMatch(&m)
	}
}

func printLMSPlayerList(data []byte) {
	var m match.LMSMatch
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Println(err.Error())
	} else {
		printLMSMatch(&m)
	}
}

func printDuelPlayerList(data []byte) {
	var m match.DuelMatch
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Println(err.Error())
	} else {
		printDuelMatch(&m)
	}
}

func printPlayerList(gt string, data []byte) {
	fmt.Printf("---------------\nMatch type: %v\n---------------\n", gt)

	switch gt {
	case match.DeathMatch:
		printDMPlayerList(data)
	case match.CaptureTheFlag:
		printCTFPlayerList(data)
	case match.LastManStanding:
		printLMSPlayerList(data)
	case match.Duel:
		printDuelPlayerList(data)
	}
}

// runMatchLoop obtains input from the player
// generates messages based on the input, sends
// them to the server and processes its responses
func runMatchLoop(c *http.Client, ps net.PlayerSession, ms net.MatchSession) error {
	conn, err := ws.CreateSession()
	if err != nil {
		return err
	}

	for {
		msgID := getUserAction()
		msg, err := ws.NewMessage(msgID, ps, ms)

		if err != nil {
			return err
		} else if err = conn.WriteJSON(msg); err != nil {
			return err
		}

		var data []byte
		_, data, err = conn.ReadMessage()
		if err != nil {
			return err
		}

		if msgID == ws.PlayerListMessageID {
			printPlayerList(ms.ID.Type, data)
		} else if msgID == ws.QuitMessageID {
			break
		}
	}

	return nil
}
