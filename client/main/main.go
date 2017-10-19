package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	serverAddress string = "http://localhost:8000"
)

type playerAuthData struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

// ToGet converts the contents of the
// playerAuthData instance into GET request
// arguments
func (data playerAuthData) ToGet() string {
	return "id=" + strconv.Itoa(data.ID) + "&token=" + data.Token
}

type playerLogin struct {
	Nick     string `json:"name"`
	Password string `json:"password"`
}

var players = []playerLogin{
	{"phreak", "c0mm4nd0"},
	{"Kamikaze", "Get'em"},
	{"Lone_Hunter", "SniperFtw"},
	{"ne0phyte", "star4748"},
	{"TheDamned1", "f4llen1"},
	{"SoulScorcher", "Burn'em_all!"},
	{"LittlePony", "pink"},
	{"Mikky", "Come|Get|Some"},
}

func processBody(resp *http.Response, out interface{}) error {
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

	return json.Unmarshal(respData, out)
}

func post(client *http.Client, endpoint string, data []byte, out interface{}) error {
	resp, err := client.Post(serverAddress+endpoint, "application/json",
		bytes.NewReader(data))
	if err != nil {
		return err
	} else if err = processBody(resp, out); err != nil {
		return err
	}

	return nil
}

func get(client *http.Client, endpoint string, auth playerAuthData, out interface{}) error {
	resp, err := client.Get(serverAddress + endpoint + "?" + auth.ToGet())
	if err != nil {
		return err
	} else if err = processBody(resp, out); err != nil {
		return err
	}

	return nil
}

func printLoginQuery() playerLogin {
	fmt.Println("You need to log in first. Pick one of the nicknames below:")
	for i, player := range players {
		fmt.Printf("%v. %v\n", i+1, player.Nick)
	}

	choice := 0
	for {
		fmt.Print("Who do you want to log in as? ")
		if _, err := fmt.Scanf("%v", &choice); err != nil || choice < 1 || choice > len(players) {
			fmt.Println("Please enter a number between 1 and", len(players))
		} else {
			break
		}
	}
	return players[choice-1]
}

func login(client *http.Client) (playerLogin, playerAuthData, error) {
	player := printLoginQuery()
	loginJSON, err := json.Marshal(player)
	if err != nil {
		return playerLogin{}, playerAuthData{}, err
	}

	var authData playerAuthData
	err = post(client, "/login", loginJSON, authData)
	if err != nil {
		return playerLogin{}, playerAuthData{}, err
	}

	return player, authData, nil
}

func printMainMenu() int {
	fmt.Println("These are your options:")
	fmt.Println("1. List active matches")
	fmt.Println("2. List leaderboard for a given game type")
	fmt.Println("3. Join an ongoing match")
	fmt.Println("4. Create a new match of the given type")

	var choice int
	for {
		fmt.Print("\nYour choice? ")
		if _, err := fmt.Scanf("%v", &choice); err != nil || choice < 1 || choice > 4 {
			fmt.Println("Please pick a number between 1 and 4")
		} else {
			break
		}
	}

	return choice
}

func listMatches(client *http.Client, authData playerAuthData) error {
	var matches map[string]interface{}

	if err := get(client, "/matches", authData, matches); err != nil {
		return err
	}

	// TODO List the contents of the matchlist
	return nil
}

func main() {
	client := http.Client{}
	player, authData, err := login(&client)
	if err != nil {
		fmt.Println("Failed to log into the server: " + err.Error())
		return
	}

	fmt.Printf("Logged in as %v\n\n", player.Nick)
	choice := printMainMenu()

	switch choice {
	case 1:
		listMatches(&client, authData)
	default:
	}

}
