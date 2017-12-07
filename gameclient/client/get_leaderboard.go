package client

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mrclayman/rest-and-go/gameclient/client/leaderboard"
	"github.com/mrclayman/rest-and-go/gameclient/client/match"
	"github.com/mrclayman/rest-and-go/gameclient/client/net"
)

func getDesiredLeaderboardType() string {
	var gtype string
	for {
		fmt.Println("Allowed game modes are (pick a mode and type in its abbreviation or the full name):")
		fmt.Println("* DM/Deathmatch")
		fmt.Println("* CTF/Capture the Flag")
		fmt.Println("* LMS/Last Man Standing")
		fmt.Println("* Duel")
		fmt.Print("Which game type are you interested in? ")
		gtype = ReadLine()

		if lwrGType, ok := match.IsValidGameType(gtype); !ok {
			fmt.Println("'" + gtype + "' is not a valid game type")
		} else {
			gtype = lwrGType
			break
		}
	}
	return gtype
}

// GetLeaderboard queries the server for leaderboard for the given
// game type and prints it out to stdout
func GetLeaderboard(c *http.Client, ps net.PlayerSession) error {
	gt := getDesiredLeaderboardType()

	var err error
	switch gt {
	case match.DeathMatch:
		err = handleDMLeaderboard(c, ps)
	case match.CaptureTheFlag:
		err = handleCTFLeaderboard(c, ps)
	case match.LastManStanding:
		err = handleLMSLeaderboard(c, ps)
	case match.Duel:
		err = handleDuelLeaderboard(c, ps)
	default:
		err = errors.New("Unhandled game type '" + gt + "' in GetLeaderboard()")
	}

	if err != nil {
		return err
	}

	return nil
}

// handleDMLeaderboard unmarshals the data in the
// input map into a list of DeathMatch-type leaderboard
// records and prints out their contents
func handleDMLeaderboard(c *http.Client, ps net.PlayerSession) error {
	lboard := make(leaderboard.DMLeaderboard, 0, 10)
	if err := net.Get(c, "/leaderboards/"+match.DeathMatch, ps, &lboard); err != nil {
		return err
	}

	fmt.Println("-----------------------\nDeathMatch game type leaderboards")
	fmt.Println("Player\tKills\tDeaths\n-----------------------")
	for _, r := range lboard {
		fmt.Println(r.Player.Nick, "\t", r.Kills, "\t", r.Deaths)
	}
	fmt.Println("-----------------------")
	return nil
}

// handleCTFLeaderboard unmarshals the data in the
// input map into a list of DeathMatch-type leaderboard
// records and prints out their contents
func handleCTFLeaderboard(c *http.Client, ps net.PlayerSession) error {
	lboard := make(leaderboard.CTFLeaderboard, 0, 10)
	if err := net.Get(c, "/leaderboards/"+match.CaptureTheFlag, ps, &lboard); err != nil {
		return err
	}

	fmt.Println("-----------------------\nCapture the Flag game type leaderboards")
	fmt.Println("Player\tKills\tDeaths\tCaptures\n-----------------------")
	for _, r := range lboard {
		fmt.Println(r.Player.Nick, "\t", r.Kills, "\t", r.Deaths, "\t", r.Captures)
	}
	fmt.Println("-----------------------")
	return nil
}

// handleLMSLeaderboard unmarshals the data in the
// input map into a list of DeathMatch-type leaderboard
// records and prints out their contents
func handleLMSLeaderboard(c *http.Client, ps net.PlayerSession) error {
	lboard := make(leaderboard.LMSLeaderboard, 0, 10)
	if err := net.Get(c, "/leaderboards/"+match.LastManStanding, ps, &lboard); err != nil {
		return err
	}

	fmt.Println("-----------------------\nLast Man Standing game type leaderboards")
	fmt.Println("Player\tKills\tDeaths\tWins\n-----------------------")
	for _, r := range lboard {
		fmt.Println(r.Player.Nick, "\t", r.Kills, "\t", r.Deaths, "\t", r.Wins)
	}
	fmt.Println("-----------------------")
	return nil
}

// handleDuelLeaderboard unmarshals the data in the
// input map into a list of DeathMatch-type leaderboard
// records and prints out their contents
func handleDuelLeaderboard(c *http.Client, ps net.PlayerSession) error {
	lboard := make(leaderboard.DuelLeaderboard, 0, 10)
	if err := net.Get(c, "/leaderboards/"+match.Duel, ps, &lboard); err != nil {
		return err
	}

	fmt.Println("-----------------------\nDuel game type leaderboards")
	fmt.Println("Player\tKills\tDeaths\tWins\n-----------------------")
	for _, r := range lboard {
		fmt.Println(r.Player.Nick, "\t", r.Kills, "\t", r.Deaths, "\t", r.Wins)
	}
	fmt.Println("-----------------------")
	return nil
}
