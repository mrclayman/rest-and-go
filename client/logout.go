package client

import "net/http"
import "encoding/json"

// Logout sends a "quit" message that logs the player out of
// the server
func Logout(c *http.Client, auth PlayerAuthData) error {

	logoutMessageData := map[string]interface{}{
		"id":    auth.ID,
		"token": auth.Token,
	}

	var postData []byte
	var err error
	if postData, err = json.Marshal(logoutMessageData); err != nil {
		return err
	}

	return post(c, "/logout", postData, nil)
}
