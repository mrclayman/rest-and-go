package net

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/mrclayman/rest-and-go/gameclient/client/match"
	"github.com/mrclayman/rest-and-go/gameclient/config"
)

const (
	// InvalidMessage defines the ID
	// of an invalid message that should
	// never be dispatched
	InvalidMessage uint16 = 0

	// WeaponFireMessage defines the ID
	// of a message that indicates the
	// player fired a weapon
	WeaponFireMessage uint16 = 1

	// PlayerMoveMessage defines the ID
	// of a message that indicates the
	// player moved to a new position
	PlayerMoveMessage uint16 = 2

	// PlayerListMessage defines the ID
	// of a message that indicates the
	// player requested a list of players
	// present in the match
	PlayerListMessage uint16 = 3

	// QuitMessage defines the ID
	// of a message that indicates the
	// player has requested to quit the match
	QuitMessage uint16 = 4
)

// Message implements the basic structure
// for a message sent on the WebSocket
// interface
type Message struct {
	MessageID uint16      `json:"message_id"`
	PlayerID  int         `json:"player_id"`
	Match     match.ID    `json:"match"`
	Token     string      `json:"token"`
	Data      interface{} `json:"data"`
}

// ConnectSession launches a WebSocket session
// for the player
func ConnectSession() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: config.Cfg.Conn.ServerURL, Path: "/match/room"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func createWeaponFiredMessage(ps PlayerSession, ms MatchSession) *Message {
	msg := Message{
		MessageID: WeaponFireMessage,
		PlayerID:  ps.ID,
		Match:     ms.ID,
		Token:     ms.Token,
		Data:      nil,
	}

	return &msg
}

func createPlayerMoveMessage(ps PlayerSession, ms MatchSession) *Message {
	msg := Message{
		MessageID: PlayerMoveMessage,
		PlayerID:  ps.ID,
		Match:     ms.ID,
		Token:     ms.Token,
		// TODO Allow user input of the coordinates??
		Data: []float64{34.4154367, -123.42362662, 11.23267334},
	}

	return &msg
}

func createPlayerListMessage(ps PlayerSession, ms MatchSession) *Message {
	msg := Message{
		MessageID: PlayerListMessage,
		PlayerID:  ps.ID,
		Match:     ms.ID,
		Token:     ms.Token,
		Data:      nil,
	}

	return &msg
}

func createQuitMatchMessage(ps PlayerSession, ms MatchSession) *Message {
	msg := Message{
		MessageID: QuitMessage,
		PlayerID:  ps.ID,
		Match:     ms.ID,
		Token:     ms.Token,
		Data:      nil,
	}

	return &msg
}

// CreateMessage synthesizes a message of the given ID
func CreateMessage(ps PlayerSession, ms MatchSession, msgID uint16) (*Message, error) {

	switch msgID {
	case WeaponFireMessage:
		return createWeaponFiredMessage(ps, ms), nil
	case PlayerMoveMessage:
		return createPlayerMoveMessage(ps, ms), nil
	case PlayerListMessage:
		return createPlayerListMessage(ps, ms), nil
	case QuitMessage:
		return createQuitMatchMessage(ps, ms), nil
	default:
		return nil, errors.New("Unhandled message type " + strconv.Itoa(int(msgID)))
	}
}
