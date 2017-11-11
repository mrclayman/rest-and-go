package client

import (
	"fmt"
	"net/http"
)

// ListMatches queries the server using player's credentials
// and prints out to stdout the list of active matches
func ListMatches(c *http.Client, auth PlayerAuthData) error {
	var matches Matchlist

	if err := get(c, "/matches", auth, &matches); err != nil {
		return err
	}

	fmt.Println("---------------------------------------------")
	for _, match := range matches {
		fmt.Println("Match ID:", match.ID)
		fmt.Println("Game type:", match.Type)
		fmt.Println("Ranks:\nPlayer\t\tKills\tDeaths\n---")
		for _, rank := range match.Ranks {
			fmt.Println(rank.PlayerName, "\t\t", rank.Kills, "\t", rank.Deaths)
		}
		fmt.Println("---------------------------------------------")
	}
	return nil
}
