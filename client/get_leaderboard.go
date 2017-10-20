package client

import (
	"fmt"
	"net/http"
	"strings"
)

func isValidGameType(gt string) (string, bool) {
	gt = strings.ToLower(gt)

	switch gt {
	case "deathmatch":
		return "dm", true
	case "capture_the_flag":
		fallthrough
	case "capture the flag":
		return "ctf", true
	case "last_man_standing":
		fallthrough
	case "last man standing":
		return "lms", true
	case "dm":
		fallthrough
	case "ctf":
		fallthrough
	case "lms":
		fallthrough
	case "duel":
		return gt, true
	default:
		return "", false
	}
}

// GetLeaderboard queries the server for leaderboard for the given
// game type and prints it out to stdout
func GetLeaderboard(c *http.Client, auth PlayerAuthData) error {
	var gtype string
	for {
		fmt.Println("Allowed game modes are (pick a mode and type in its abbreviation or the full name):")
		fmt.Println("* DM/Deathmatch")
		fmt.Println("* CTF/Capture the Flag")
		fmt.Println("* LMS/Last Man Standing")
		fmt.Println("* Duel")
		fmt.Print("Which game type are you interested in? ")
		fmt.Scanf("%v", &gtype)

		if lwrGType, ok := isValidGameType(gtype); !ok {
			fmt.Println(gtype, "is not a valid game type")
		} else {
			gtype = lwrGType
			break
		}
	}

	var lboard Leaderboard

	if err := get(c, "/leaderboards/"+gtype, auth, &lboard); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("-----------------------\nLeaderboards for game type '" + gtype + "'")
	fmt.Println("Player\tKills\tDeaths\n-----------------------")
	for _, record := range lboard {
		fmt.Println(record.Nick, "\t", record.Kills, "\t", record.Deaths)
	}
	fmt.Println("-----------------------")
	return nil
}
