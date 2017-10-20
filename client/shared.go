package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func processResponse(resp *http.Response, out interface{}) error {
	var respData []byte
	var err error

	respData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	} else if resp.StatusCode/100 != 2 {
		return errors.New(string(respData))
	} else if resp.Header.Get("Content-Type") != "application/json" {
		return errors.New("Unexpected response content type")
	}

	//fmt.Println(string(respData))
	err = json.Unmarshal(respData, out)
	if err != nil {
		fmt.Println("Failed to unmarshal data:", err.Error())
	}

	return err
}

func post(client *http.Client, endpoint string, data []byte, out interface{}) error {
	resp, err := client.Post(serverAddress+endpoint, "application/json",
		bytes.NewReader(data))
	if err != nil {
		return err
	} else if err = processResponse(resp, out); err != nil {
		return err
	}

	return nil
}

func get(client *http.Client, endpoint string, auth PlayerAuthData, out interface{}) error {
	url := serverAddress + endpoint + "?" + auth.ToGet()
	resp, err := client.Get(url)
	if err != nil {
		return err
	} else if err = processResponse(resp, out); err != nil {
		return err
	}

	return nil
}

func connectSession(sessionData *MatchSessionData) (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: serverAddress, Path: "/match/room"}
}

func enterMatchLoop(c *http.Client, sessionData *MatchSessionData) error {

}
