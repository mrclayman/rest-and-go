package main

import (
	"fmt"
	"net/http"

	"github.com/mrclayman/rest-and-go/client"
)

////////////// Functions ////////////////

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

func main() {
	cl := http.Client{}
	player, authData, err := client.Login(&cl)
	if err != nil {
		fmt.Println("Failed to log into the server: " + err.Error())
		return
	}

	fmt.Printf("Logged in as %v\n\n", player.Nick)

	for {
		choice := printMainMenu()
		var err error

		switch choice {
		case 1:
			fmt.Println("Listing matches")
			err = client.ListMatches(&cl, authData)
		case 2:
			fmt.Println("Getting leaderboards")
			err = client.GetLeaderboard(&cl, authData)
		case 3:
			fmt.Println("Joining a match")
			err = client.JoinMatch(&cl, authData)
		default:
			break
		}

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
