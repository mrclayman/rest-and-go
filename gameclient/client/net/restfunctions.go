package net

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mrclayman/rest-and-go/gameclient/client/shared"
	"github.com/mrclayman/rest-and-go/gameclient/config"
)

func processResponse(resp *http.Response, out interface{}) error {
	respData, err := ioutil.ReadAll(resp.Body)
	respDataLen := len(respData)
	if err != nil {
		return err
	} else if resp.StatusCode/100 != 2 {
		return errors.New(string(respData))
	} else if respDataLen == 0 {
		return nil
	} else if resp.Header.Get("Content-Type") != "application/json" {
		return errors.New("Unexpected response content type")
	}

	//fmt.Println(string(respData))
	if out != nil {
		if err != nil {
			fmt.Println("Failed to unmarshal data:", err.Error())
		}
	}

	err = shared.DecodeJSON(respData, out)
	return err
}

// Post sends a POST request to the server, then parses
// the reply and stores it in the 'out' argument
func Post(client *http.Client, endpoint string, data []byte, out interface{}) error {
	resp, err := client.Post(restAPIProtocol+config.Cfg.Conn.ServerURL+endpoint, "application/json",
		bytes.NewReader(data))
	if err != nil {
		return err
	} else if err = processResponse(resp, out); err != nil {
		return err
	}

	return nil
}

// Get sends a GET request to the server, then parses
// the reply and stores it in the 'out' argument
func Get(client *http.Client, endpoint string, ps PlayerSession, out interface{}) error {
	url := restAPIProtocol + config.Cfg.Conn.ServerURL + endpoint + "?" + ps.ToGet()
	resp, err := client.Get(url)
	if err != nil {
		return err
	} else if err = processResponse(resp, out); err != nil {
		return err
	}

	return nil
}
