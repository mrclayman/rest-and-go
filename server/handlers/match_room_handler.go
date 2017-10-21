package handlers

import (
	"fmt"
	"log"
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
	log.Println("Received match room request, will spawn a separate goroutine to handle communication")

	// Obtain the connection object from the request
	conn, err := h.upgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Printf("Failed to obtain connection object from request: %v", err.Error())
		http.Error(resp, "Failed to obtain connection from request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	go h.handleWebSockConnection(conn)
}

// handleWebSockConnection processes all messages coming through
// the WebSocket connection
func (h *MatchRoomHandler) handleWebSockConnection(conn *websocket.Conn) error {
	defer conn.Close()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Failed to read message on WS interface: %v", err.Error())
			return err
		}

		if !isValidMessageID(msg.MsgID) {
			log.Printf("WS message has unknown id %v", msg.MsgID)
			return RequestError{"Invalid WS message id"}
		} else if msg.Token == core.InvalidWebSocketToken {
			log.Println("Invalid WS token provided by player")
			return RequestError{"Invalid WS token"}
		} else if !h.core.IsInMatch(msg.PID, msg.Token) {
			log.Printf("Player %v not in match", msg.PID)
			return RequestError{"Player not registered in match"}
		}

		// Handle individual messages
		switch msg.MsgID {
		case WeaponFiredMessageID:
			err = h.handleWeaponFiredMessage(conn, &msg)
		case MoveMessageID:
			err = h.handlePlayerMoveMessage(conn, &msg)
		case PlayerListMessageID:
			err = h.handlePlayerListMessage(conn, &msg)
		case QuitMessageID:
			err = h.handlePlayerQuitMessage(conn, &msg)
		}

		if err != nil {
			log.Println(err.Error())
		}
		if msg.MsgID == QuitMessageID {
			return err
		}
	}
}

func (h *MatchRoomHandler) handleWeaponFiredMessage(conn *websocket.Conn, msg *Message) error {
	// TODO Log data from the message on the output
	fmt.Println("TEST: Weapon fired")
	return conn.WriteJSON(nil)
}

func (h *MatchRoomHandler) handlePlayerMoveMessage(conn *websocket.Conn, msg *Message) error {
	// TODO Log data from the message on the output
	fmt.Println("TEST: Player moved")
	return conn.WriteJSON(nil)
}

func (h *MatchRoomHandler) handlePlayerListMessage(conn *websocket.Conn, msg *Message) error {
	fmt.Println("TEST: Player requested a player list")
	return conn.WriteJSON(nil)
}

func (h *MatchRoomHandler) handlePlayerQuitMessage(conn *websocket.Conn, msg *Message) error {
	fmt.Println("TEST: Player is quitting the match")
	return conn.WriteJSON(nil)
}
