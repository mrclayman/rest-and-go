package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type playerAuthData struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

const (
	serverAddress string = "http://localhost:8000"
)

type player struct {
	Nick     string `json:"name"`
	Password string `json:"password"`
}

var players = []player{
	{"phreak", "c0mm4nd0"},
	{"Kamikaze", "Get'em"},
	{"Lone_Hunter", "SniperFtw"},
	{"ne0phyte", "star4748"},
	{"TheDamned1", "f4llen1"},
	{"SoulScorcher", "Burn'em_all!"},
	{"LittlePony", "pink"},
	{"Mikky", "Come|Get|Some"},
}

func post(client *http.Client, endpoint string, data []byte) ([]byte, error) {
	resp, err := client.Post(serverAddress+endpoint, "application/json",
		bytes.NewReader(data))
	if err != nil {
		return []byte{}, err
	} else if resp.Header.Get("Content-Type") != "application/json" {
		return []byte{}, errors.New("Unexpected response content type")
	}

	var respData []byte
	respData, err = ioutil.ReadAll(resp.Body)
	return respData, err
}

func loginQuery() player {
	fmt.Println("Who do you want to log in as?")
	for i, player := range players {
		fmt.Printf("%v. %v\n", i+1, player.Nick)
	}
	choice := 0
	for {
		_, err := fmt.Scanf("%v", &choice)
		if err != nil || choice < 1 || choice > len(players) {
			fmt.Println("Please enter a number between 1 and", len(players))
			fmt.Println("Who do you want to log in as?")
		} else {
			break
		}
	}
	return players[choice]
}

func login(client *http.Client) (playerAuthData, error) {
	player := loginQuery()
	loginJSON, err := json.Marshal(player)
	if err != nil {
		return playerAuthData{}, err
	}

	var resp []byte
	resp, err = post(client, "/login", loginJSON)
	if err != nil {
		return playerAuthData{}, err
	}

	var authData playerAuthData
	if err = json.Unmarshal(resp, &authData); err != nil {
		return playerAuthData{}, errors.New("Failed to unmarshal server's reply to our login request: " + err.Error())
	}

	return authData, nil
}

func main() {
	client := new(http.Client)
	loginCreds, err := login(client)
	if err != nil {
		log.Fatal("Failed to log into the server: " + err.Error())
	}

	fmt.Println("ID:", loginCreds.ID, ", Token:", loginCreds.Token)
}
