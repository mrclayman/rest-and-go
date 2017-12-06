package client

import (
	"fmt"
	"net/http"

	"github.com/mrclayman/rest-and-go/gameclient/client/match"
	"github.com/mrclayman/rest-and-go/gameclient/client/net"
)

// ListMatches queries the server using player's credentials
// and prints out to stdout the list of active matches
func ListMatches(c *http.Client, ps net.PlayerSession) (*match.Matchlist, error) {
	var mu match.MatchlistUnmarshaler

	if err := net.Get(c, "/matches", ps, &mu); err != nil {
		return nil, err
	}

	listDMMatches(&mu.Matchlist.DM)
	listCTFMatches(&mu.Matchlist.CTF)
	listLMSMatches(&mu.Matchlist.LMS)
	listDuelMatches(&mu.Matchlist.Duel)
	return &mu.Matchlist, nil
}

func listDMMatches(ml *match.DMMatches) {
	fmt.Println("---------------------------------------------")
	fmt.Println("Game type:", match.DeathMatch)
	if len(*ml) == 0 {
		fmt.Println("< No matches to display >")
		fmt.Println("---------------------------------------------")
		return
	}

	for _, match := range *ml {
		fmt.Println("Match ID:", match.Number)
		fmt.Println("Ranks:\nPlayer\t\tKills\tDeaths\n---")
		for _, rank := range match.Ranks {
			fmt.Println(rank.Player.Nick, "\t\t", rank.Kills, "\t", rank.Deaths)
		}
		fmt.Println("---------------------------------------------")
	}
}

func listCTFMatches(ml *match.CTFMatches) {
	fmt.Println("---------------------------------------------")
	fmt.Println("Game type:", match.CaptureTheFlag)
	if len(*ml) == 0 {
		fmt.Println("< No matches to display >")
		fmt.Println("---------------------------------------------")
		return
	}

	for _, match := range *ml {
		fmt.Println("Match ID:", match.Number)
		fmt.Println("Ranks:\nPlayer\t\tKills\tDeaths\tCaptures\n---")
		for _, rank := range match.Ranks {
			fmt.Println(rank.Player.Nick, "\t\t", rank.Kills, "\t", rank.Deaths, "\t", rank.Captures)
		}
		fmt.Println("---------------------------------------------")
	}
}

func listLMSMatches(ml *match.LMSMatches) {
	fmt.Println("---------------------------------------------")
	fmt.Println("Game type:", match.LastManStanding)
	if len(*ml) == 0 {
		fmt.Println("< No matches to display >")
		fmt.Println("---------------------------------------------")
		return
	}

	for _, match := range *ml {
		fmt.Println("Match ID:", match.Number)
		fmt.Println("Ranks:\nPlayer\t\tKills\tDeaths\n---")
		for _, rank := range match.Ranks {
			fmt.Println(rank.Player.Nick, "\t\t", rank.Kills, "\t", rank.Deaths)
		}
		fmt.Println("---------------------------------------------")
	}
}

func listDuelMatches(ml *match.DuelMatches) {
	fmt.Println("---------------------------------------------")
	fmt.Println("Game type:", match.Duel)
	if len(*ml) == 0 {
		fmt.Println("< No matches to display >")
		fmt.Println("---------------------------------------------")
		return
	}

	for _, match := range *ml {
		fmt.Println("Match ID:", match.Number)
		fmt.Println("Ranks:\nPlayer\t\tKills\tDeaths\n---")
		for _, rank := range match.Ranks {
			fmt.Println(rank.Player.Nick, "\t\t", rank.Kills, "\t", rank.Deaths)
		}
		fmt.Println("---------------------------------------------")
	}
}
