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

	var lboard []map[string]interface{}

	if err := net.Get(c, "/leaderboards/"+gt, ps, &lboard); err != nil {
		return err
	}

	var err error
	switch gt {
	case match.DeathMatch:
		err = handleDMLeaderboard(lboard)
	case match.CaptureTheFlag:
		err = handleCTFLeaderboard(lboard)
	case match.LastManStanding:
		err = handleLMSLeaderboard(lboard)
	case match.Duel:
		err = handleDuelLeaderboard(lboard)
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
func handleDMLeaderboard(lbMap []map[string]interface{}) error {
	lb, err := leaderboard.UnmarshalDMLeaderboard(lbMap)
	if err != nil {
		return err
	}

	fmt.Println("-----------------------\nDeathMatch game type leaderboards")
	fmt.Println("Player\tKills\tDeaths\n-----------------------")
	for _, r := range *lb {
		fmt.Println(r.Player.Nick, "\t", r.Kills, "\t", r.Deaths)
	}
	fmt.Println("-----------------------")
	return nil
}

// handleCTFLeaderboard unmarshals the data in the
// input map into a list of DeathMatch-type leaderboard
// records and prints out their contents
func handleCTFLeaderboard(lbMap []map[string]interface{}) error {
	lb, err := leaderboard.UnmarshalCTFLeaderboard(lbMap)
	if err != nil {
		return err
	}

	fmt.Println("-----------------------\nCapture the Flag game type leaderboards")
	fmt.Println("Player\tKills\tDeaths\tCaptures\n-----------------------")
	for _, r := range *lb {
		fmt.Println(r.Player.Nick, "\t", r.Kills, "\t", r.Deaths, "\t", r.Captures)
	}
	fmt.Println("-----------------------")
	return nil
}

// handleLMSLeaderboard unmarshals the data in the
// input map into a list of DeathMatch-type leaderboard
// records and prints out their contents
func handleLMSLeaderboard(lbMap []map[string]interface{}) error {
	lb, err := leaderboard.UnmarshalLMSLeaderboard(lbMap)
	if err != nil {
		return err
	}

	fmt.Println("-----------------------\nLast Man Standing game type leaderboards")
	fmt.Println("Player\tKills\tDeaths\tWins\n-----------------------")
	for _, r := range *lb {
		fmt.Println(r.Player.Nick, "\t", r.Kills, "\t", r.Deaths, "\t", r.Wins)
	}
	fmt.Println("-----------------------")
	return nil
}

// handleDuelLeaderboard unmarshals the data in the
// input map into a list of DeathMatch-type leaderboard
// records and prints out their contents
func handleDuelLeaderboard(lbMap []map[string]interface{}) error {
	lb, err := leaderboard.UnmarshalDuelLeaderboard(lbMap)
	if err != nil {
		return err
	}

	fmt.Println("-----------------------\nDuel game type leaderboards")
	fmt.Println("Player\tKills\tDeaths\tWins\n-----------------------")
	for _, r := range *lb {
		fmt.Println(r.Player.Nick, "\t", r.Kills, "\t", r.Deaths, "\t", r.Wins)
	}
	fmt.Println("-----------------------")
	return nil
}
