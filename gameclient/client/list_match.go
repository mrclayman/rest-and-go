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
	var m match.Matchlist

	if err := net.Get(c, "/matches", ps, &m); err != nil {
		return nil, err
	}

	listDMMatches(&m.DM)
	listCTFMatches(&m.CTF)
	listLMSMatches(&m.LMS)
	listDuelMatches(&m.Duel)
	return &m, nil
}

func listDMMatches(ml *match.DMMatches) {
	fmt.Println("---------------------------------------------")
	fmt.Println("Game type:", match.DeathMatch)
	if len(*ml) == 0 {
		fmt.Println("< No matches to display >")
		fmt.Println("---------------------------------------------")
		return
	}

	for _, m := range *ml {
		printDMMatch(m)
		fmt.Println("---------------------------------------------")
	}
}

func printDMMatch(m *match.DMMatch) {
	fmt.Println("Match ID:", m.Number)
	fmt.Println("Ranks:\nPlayer\t\tKills\tDeaths\n---")
	for _, rank := range m.Ranks {
		fmt.Println(rank.Player.Nick, "\t\t", rank.Kills, "\t", rank.Deaths)
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

	for _, m := range *ml {
		printCTFMatch(m)
		fmt.Println("---------------------------------------------")
	}
}

func printCTFMatch(m *match.CTFMatch) {
	fmt.Println("Match ID:", m.Number)
	fmt.Println("Ranks:\nPlayer\t\tKills\tDeaths\tCaptures\n---")
	for _, rank := range m.Ranks {
		fmt.Println(rank.Player.Nick, "\t\t", rank.Kills, "\t", rank.Deaths, "\t", rank.Captures)
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

	for _, m := range *ml {
		printLMSMatch(m)
		fmt.Println("---------------------------------------------")
	}
}

func printLMSMatch(m *match.LMSMatch) {
	fmt.Println("Match ID:", m.Number)
	fmt.Println("Ranks:\nPlayer\t\tKills\tDeaths\n---")
	for _, rank := range m.Ranks {
		fmt.Println(rank.Player.Nick, "\t\t", rank.Kills, "\t", rank.Deaths)
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

	for _, m := range *ml {
		printDuelMatch(m)
		fmt.Println("---------------------------------------------")
	}
}

func printDuelMatch(m *match.DuelMatch) {
	fmt.Println("Match ID:", m.Number)
	fmt.Println("Ranks:\nPlayer\t\tKills\tDeaths\n---")
	for _, rank := range m.Ranks {
		fmt.Println(rank.Player.Nick, "\t\t", rank.Kills, "\t", rank.Deaths)
	}
}
