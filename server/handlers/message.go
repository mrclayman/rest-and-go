package handlers

import "github.com/mrclayman/rest-and-go/server/core"

// MessageID identifies an incoming message on the
// WebSocket interface
type MessageID uint16

const (
	// InvalidMessageID identifies an invalid value
	// for a message ID
	InvalidMessageID MessageID = iota

	// WeaponFiredMessageID identifies a message about
	// a player firing a weapon
	WeaponFiredMessageID MessageID = iota

	// MoveMessageID identifies a message about
	// a player moving to a new position
	MoveMessageID MessageID = iota

	// PlayerListMessageID identifies a message about
	// a player querying the player list
	PlayerListMessageID MessageID = iota

	// QuitMessageID indicates that a player
	// wishes to quit the match
	QuitMessageID MessageID = iota
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
	MatchID  core.MatchID        `json:"match_id"`
	PlayerID core.PlayerID       `json:"player_id"`
	Token    core.WebSocketToken `json:"token"`
	Data     interface{}         `json:"data"`
}
