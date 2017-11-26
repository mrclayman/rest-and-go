package handlers

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/auth"
	"github.com/mrclayman/rest-and-go/gameserver/core/match"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
)

// MessageID identifies an incoming message on the
// WebSocket interface
type MessageID uint16

const (
	// InvalidMessageID identifies an invalid value
	// for a message ID
	InvalidMessageID MessageID = 0

	// WeaponFiredMessageID identifies a message about
	// a player firing a weapon
	WeaponFiredMessageID MessageID = 1

	// MoveMessageID identifies a message about
	// a player moving to a new position
	MoveMessageID MessageID = 2

	// PlayerListMessageID identifies a message about
	// a player querying the player list
	PlayerListMessageID MessageID = 3

	// QuitMessageID indicates that a player
	// wishes to quit the match
	QuitMessageID MessageID = 4
)

// IsValidMessageID checks that the id
// in the argument identifies a valid message
func isValidMessageID(mid MessageID) bool {
	switch mid {
	case WeaponFiredMessageID:
		fallthrough
	case MoveMessageID:
		fallthrough
	case PlayerListMessageID:
		fallthrough
	case QuitMessageID:
		return true
	default:
		return false
	}
}

// Message represents a message received on the
// WebSocket interface
type Message struct {
	MsgID    MessageID           `json:"message_id"`
	MatchID  match.ID            `json:"match"`
	PlayerID player.ID           `json:"player_id"`
	Token    auth.WebSocketToken `json:"token"`
	Data     interface{}         `json:"data"`
}
