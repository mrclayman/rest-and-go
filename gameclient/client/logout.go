package client

import (
	"encoding/json"
	"net/http"

	"github.com/mrclayman/rest-and-go/gameclient/client/net"
)

// Logout sends a "quit" message that logs the player out of
// the server
func Logout(c *http.Client, ps net.PlayerSession) error {

	logoutMessageData := map[string]interface{}{
		"id":    ps.ID,
		"token": ps.Token,
	}

	var postData []byte
	var err error
	if postData, err = json.Marshal(logoutMessageData); err != nil {
		return err
	}

	return net.Post(c, "/logout", postData, nil)
}
