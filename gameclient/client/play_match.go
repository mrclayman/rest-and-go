package client

import (
	"fmt"
	"net/http"

	"github.com/mrclayman/rest-and-go/gameclient/client/net"
	"github.com/mrclayman/rest-and-go/gameclient/client/player"
	"github.com/mrclayman/rest-and-go/gameclient/client/shared"
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

func printPlayerList(mt string, playerList map[string]interface{}) {
	fmt.Println("---------------\nMatch type:", mt)
	fmt.Println("Ranks:")
	fmt.Println("Player\tKills\tDeaths")
	fmt.Println("---------------")
	for _, rank := range playerList["ranks"].(map[string]interface{}) {
		rankMap, ok := rank.(map[string]interface{})
		if !ok {
			fmt.Println("Cannot print player list, item not a map of values")
			return
		}

		p, err := player.FromMap(rankMap)
		if err != nil {
			fmt.Println("Could not obtain player struct:", err.Error())
		}
		fmt.Println(p.Nick, "\t", rankMap["kills"], "\t", rankMap["deaths"])
	}
	fmt.Println("---------------")
	fmt.Println()
}

func runMatchLoop(c *http.Client, ps net.PlayerSession, ms net.MatchSession) error {
	conn, err := net.CreateSession()
	if err != nil {
		return err
	}

	for {
		msgID := getUserAction()
		msg, err := net.CreateMessage(ps, ms, msgID)

		if err != nil {
			return err
		} else if err = conn.WriteJSON(msg); err != nil {
			return err
		}

		var respData map[string]interface{}
		var data []byte

		// I cannot use ReadJSON() because I need
		// special handling of numeric values
		_, data, err = conn.ReadMessage()
		if err != nil {
			return err
		} else if err = shared.DecodeJSON(data, &respData); err != nil {
			return err
		}

		if msgID == net.PlayerListMessage {
			printPlayerList(ms.ID.Type, respData)
		} else if msgID == net.QuitMessage {
			break
		}
	}

	return nil
}
