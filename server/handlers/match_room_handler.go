package handlers

import (
	"log"
	"net/http"
	"strconv"

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
		} else if !h.core.IsInMatch(msg.PlayerID, msg.Token) {
			log.Printf("Player %v not in match", msg.PlayerID)
			return RequestError{"Player not registered in match"}
		}

		// Handle individual messages
		log.Printf("Received message id %v", msg.MsgID)
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
			log.Printf("Terminating communication thread for player %v because of error: %v", msg.PlayerID, err.Error())
			return err
		}
		if msg.MsgID == QuitMessageID {
			log.Printf("Terminating communication thread for player %v because of quit", msg.PlayerID)
			return err
		}
	}
}

func (h *MatchRoomHandler) handleWeaponFiredMessage(conn *websocket.Conn, msg *Message) error {
	log.Printf("TEST: Player %v fired a weapon", msg.PlayerID)
	return conn.WriteJSON(nil)
}

func (h *MatchRoomHandler) handlePlayerMoveMessage(conn *websocket.Conn, msg *Message) error {

	var values []interface{}
	var ok bool

	// We need to parse the position vector from the message, where
	// it is stored as a slice of generic (interface) values
	if values, ok = msg.Data.([]interface{}); !ok {
		return RequestError{"Expected a slice of values as the position vector"}
	} else if len(values) != 3 {
		return RequestError{"The position vector does not have 3 elements, but " + strconv.Itoa(len(values))}
	}

	position := make([]float64, 3)
	for i, value := range values {
		var number float64
		if number, ok = value.(float64); !ok {
			return RequestError{"Expected a 64-bit float in position " + strconv.Itoa(i) + " in the position slice"}
		}
		position[i] = number
	}

	log.Printf("TEST: Player moved to position x = %v, y = %v, z = %v\n", position[0], position[1], position[2])
	return conn.WriteJSON(nil)
}

func (h *MatchRoomHandler) handlePlayerListMessage(conn *websocket.Conn, msg *Message) error {
	log.Printf("TEST: Player %v requested a player list", msg.PlayerID)
	jsonMatch, err := h.core.GetMatchForJSON(msg.MatchID)
	if err != nil {
		return err
	}

	err = conn.WriteJSON(jsonMatch)
	if err != nil {
		log.Println("An error occurred while marshaling JSON player list: " + err.Error())
	} else {
		log.Printf("Player list response dispatched to player %v", msg.PlayerID)
	}
	return err
}

func (h *MatchRoomHandler) handlePlayerQuitMessage(conn *websocket.Conn, msg *Message) error {
	log.Printf("TEST: Player %v is quitting the match", msg.PlayerID)
	if err := h.core.QuitMatch(msg.MatchID, msg.PlayerID); err != nil {
		return RequestError{err.Error()}
	}

	return conn.WriteJSON(nil)
}
