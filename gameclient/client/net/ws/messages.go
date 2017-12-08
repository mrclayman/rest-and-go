package ws

import (
	"errors"
	"strconv"

	"github.com/mrclayman/rest-and-go/gameclient/client/match"
	"github.com/mrclayman/rest-and-go/gameclient/client/net"
)

const (
	// InvalidMessageID defines the ID
	// of an invalid message that should
	// never be dispatched
	InvalidMessageID uint16 = 0

	// WeaponFireMessageID defines the ID
	// of a message that indicates the
	// player fired a weapon
	WeaponFireMessageID uint16 = 1

	// PlayerMoveMessageID defines the ID
	// of a message that indicates the
	// player moved to a new position
	PlayerMoveMessageID uint16 = 2

	// PlayerListMessageID defines the ID
	// of a message that indicates the
	// player requested a list of players
	// present in the match
	PlayerListMessageID uint16 = 3

	// QuitMessageID defines the ID
	// of a message that indicates the
	// player has requested to quit the match
	QuitMessageID uint16 = 4
)

// WeaponFireMessage is used to notify server
// of a user firing a weapon
type WeaponFireMessage struct {
	MessageID uint16   `json:"message_id"`
	PlayerID  int      `json:"player_id"`
	Match     match.ID `json:"match"`
	Token     string   `json:"token"`
}

// newWeaponFireMessage is used to synthesize
// an instance of the WeaponFireMessage structure
func newWeaponFireMessage(ps net.PlayerSession, ms net.MatchSession) *WeaponFireMessage {
	return &WeaponFireMessage{
		MessageID: WeaponFireMessageID,
		PlayerID:  ps.ID,
		Match:     ms.ID,
		Token:     ms.Token,
	}
}

// PlayerMoveMessage is used to indicate to the
// server that the player has changed position
// The new position is transmitted to the server
type PlayerMoveMessage struct {
	MessageID uint16     `json:"message_id"`
	PlayerID  int        `json:"player_id"`
	Match     match.ID   `json:"match"`
	Token     string     `json:"token"`
	Data      [3]float64 `json:"data"`
}

// newPlayerMoveMessage is used to synthesize
// an instance of the PlayerMoveMessage structure
func newPlayerMoveMessage(ps net.PlayerSession, ms net.MatchSession) *PlayerMoveMessage {
	return &PlayerMoveMessage{
		MessageID: PlayerMoveMessageID,
		PlayerID:  ps.ID,
		Match:     ms.ID,
		Token:     ms.Token,
		Data:      [3]float64{34.4154367, -123.42362662, 11.23267334},
	}
}

// PlayerListMessage is used to indicate to the
// server that the player wishes to obtain the
// list of
type PlayerListMessage struct {
	MessageID uint16   `json:"message_id"`
	PlayerID  int      `json:"player_id"`
	Match     match.ID `json:"match"`
	Token     string   `json:"token"`
}

// newPlayerListMessage synthesizes a new
// PlayerListMessage instance
func newPlayerListMessage(ps net.PlayerSession, ms net.MatchSession) *PlayerListMessage {
	return &PlayerListMessage{
		MessageID: PlayerListMessageID,
		PlayerID:  ps.ID,
		Match:     ms.ID,
		Token:     ms.Token,
	}
}

// QuitMatchMessage indicates to the server
// that the player wishes to leave the match
type QuitMatchMessage struct {
	MessageID uint16   `json:"message_id"`
	PlayerID  int      `json:"player_id"`
	Match     match.ID `json:"match"`
	Token     string   `json:"token"`
}

// newQuitMatchMessage synthesizes a QuitMessage
// structure instance
func newQuitMatchMessage(ps net.PlayerSession, ms net.MatchSession) *QuitMatchMessage {
	return &QuitMatchMessage{
		MessageID: QuitMessageID,
		PlayerID:  ps.ID,
		Match:     ms.ID,
		Token:     ms.Token,
	}
}

// NewMessage creates a new message based
// on the ID provided in the argument
func NewMessage(ID uint16, ps net.PlayerSession, ms net.MatchSession) (interface{}, error) {
	switch ID {
	case WeaponFireMessageID:
		return newWeaponFireMessage(ps, ms), nil
	case PlayerMoveMessageID:
		return newPlayerMoveMessage(ps, ms), nil
	case PlayerListMessageID:
		return newPlayerListMessage(ps, ms), nil
	case QuitMessageID:
		return newQuitMatchMessage(ps, ms), nil
	default:
		return nil, errors.New("Unhandled message ID " + strconv.FormatUint(uint64(ID), 10))
	}
}
