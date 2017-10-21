package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/websocket"
)

const (
	invalidMessage uint16 = 0

	weaponFiredMessage uint16 = 1

	playerMovedMessage uint16 = 2

	playerListMessage uint16 = 3

	quitMessage uint16 = 4
)

// message implements the basic structure
// for a message sent on the WebSocket
// interface
type message struct {
	MessageID uint16      `json:"message_id"`
	PlayerID  int         `json:"player_id"`
	MatchID   uint64      `json:"match_id"`
	Token     string      `json:"token"`
	Data      interface{} `json:"data"`
}

func connectSession() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: serverAddress, Path: "/match/room"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func createWeaponFiredMessage(auth PlayerAuthData, sessionData MatchSessionData) *message {
	msg := message{
		MessageID: weaponFiredMessage,
		PlayerID:  auth.ID,
		MatchID:   sessionData.MatchID,
		Token:     sessionData.WSToken,
		Data:      nil,
	}

	return &msg
}

func createPlayerMoveMessage(auth PlayerAuthData, sessionData MatchSessionData) *message {
	msg := message{
		MessageID: playerMovedMessage,
		PlayerID:  auth.ID,
		MatchID:   sessionData.MatchID,
		Token:     sessionData.WSToken,
		// TODO Allow user input of the coordinates??
		Data: []float64{34.4154367, -123.42362662, 11.23267334},
	}

	return &msg
}

func createPlayerListMessage(auth PlayerAuthData, sessionData MatchSessionData) *message {
	msg := message{
		MessageID: playerListMessage,
		PlayerID:  auth.ID,
		MatchID:   sessionData.MatchID,
		Token:     sessionData.WSToken,
		Data:      nil,
	}

	return &msg
}

func createQuitMatchMessage(auth PlayerAuthData, sessionData MatchSessionData) *message {
	msg := message{
		MessageID: quitMessage,
		PlayerID:  auth.ID,
		MatchID:   sessionData.MatchID,
		Token:     sessionData.WSToken,
		Data:      nil,
	}

	return &msg
}

func createMessage(auth PlayerAuthData, sessionData MatchSessionData, msgID uint16) (*message, error) {

	switch msgID {
	case weaponFiredMessage:
		return createWeaponFiredMessage(auth, sessionData), nil
	case playerMovedMessage:
		return createPlayerMoveMessage(auth, sessionData), nil
	case playerListMessage:
		return createPlayerListMessage(auth, sessionData), nil
	case quitMessage:
		return createQuitMatchMessage(auth, sessionData), nil
	default:
		return nil, errors.New("Unhandled message type " + strconv.Itoa(int(msgID)))
	}
}

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

func printPlayerList(playerList map[string]interface{}) {
	fmt.Println(playerList)
	fmt.Println("---------------\nMatch type: ", playerList["match_type"])
	fmt.Println("Ranks:")
	fmt.Println("Player\tKills\tDeaths")
	fmt.Println("---------------")
	for _, rank := range playerList["ranks"].([]interface{}) {
		rankMap, ok := rank.(map[string]interface{})
		if !ok {
			fmt.Println("Cannot print player list, item not a map of values")
			return
		}
		fmt.Println(rankMap["player_name"], "\t", rankMap["kills"], "\t", rankMap["deaths"])
	}
}

func runMatchLoop(c *http.Client, auth PlayerAuthData, sessionData MatchSessionData) error {
	conn, err := connectSession()
	if err != nil {
		return err
	}

	for {
		msgID := getUserAction()
		msg, err := createMessage(auth, sessionData, msgID)

		if err != nil {
			return err
		} else if err = conn.WriteJSON(msg); err != nil {
			return err
		}

		var respData map[string]interface{}
		err = conn.ReadJSON(&respData)
		if err != nil {
			return err
		}

		if msgID == playerListMessage {
			printPlayerList(respData)
		} else if msgID == quitMessage {
			break
		}
	}

	return nil
}
