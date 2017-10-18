package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mrclayman/rest-and-go/server/core"
)

// MatchRoomHandler handles WebSocket requests
// from a player that has previously joined
// a match
type MatchRoomHandler struct {
	core     *core.Core
	upgrader websocket.Upgrader
}

// NewMatchRoomHandler returns a pointer to a
// new instance of the match room request handler
func NewMatchRoomHandler(c *core.Core) *MatchRoomHandler {
	return &MatchRoomHandler{
		core:     c,
		upgrader: websocket.Upgrader{},
	}
}

// ProcessRequest processes the WebSocket request
// from the player and creates an appropriate response
func (h *MatchRoomHandler) ProcessRequest(resp http.ResponseWriter, req *http.Request) {
	// Obtain the connection object from the request
	conn, err := h.upgrader.Upgrade(resp, req, nil)
	if err != nil {
		http.Error(resp, "Failed to obtain connection from request: "+err.Error(), http.StatusInternalServerError)
	}

	go h.handleWebSockConnection(conn)
}

// handleWebSockConnection processes all messages coming through
// the WebSocket connection
func (h *MatchRoomHandler) handleWebSockConnection(conn *websocket.Conn) error {
	defer conn.Close()

	for {
		msgType, bytes, err := conn.ReadMessage()
		if err != nil {
			return err
		} else if msgType != websocket.TextMessage {
			return RequestError{"Expected a binary message"}
		}

		var msg Message
		if err = GetJSONFromBytes(bytes, msg); err != nil {
			return err
		} else if !isValidMessageID(msg.MsgID) {
			return RequestError{"Invalid message id"}
		} else if msg.Token == core.InvalidWebSocketToken {
			return RequestError{"Invalid WS token"}
		} else if !h.core.IsInMatch(msg.PID, msg.Token) {
			return RequestError{"Player not registered in match"}
		}

		// Handle individual messages
		switch msg.MsgID {
		case WeaponFiredMessageID:
			h.handleWeaponFiredMessage(conn, &msg)
		case MoveMessageID:
			h.handlePlayerMoveMessage(conn, &msg)
		case PlayerListMessageID:
			h.handlePlayerListMessage(conn, &msg)
		case QuitMessageID:
			h.handlePlayerQuitMessage(conn, &msg)
			return nil
		}
	}
}

func (h *MatchRoomHandler) handleWeaponFiredMessage(conn *websocket.Conn, msg *Message) {
	// TODO Log data from the message on the output
	fmt.Println("TEST: Weapon fired")
}

func (h *MatchRoomHandler) handlePlayerMoveMessage(conn *websocket.Conn, msg *Message) {
	// TODO Log data from the message on the output
}

func (h *MatchRoomHandler) handlePlayerListMessage(conn *websocket.Conn, msg *Message) {

}

func (h *MatchRoomHandler) handlePlayerQuitMessage(conn *websocket.Conn, msg *Message) {

}
