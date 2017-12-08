package ws

import (
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/mrclayman/rest-and-go/gameclient/client/config"
)

// CreateSession launches a WebSocket session
// for the player
func CreateSession() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: config.Cfg.Conn.ServerURL, Path: "/match/room"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
