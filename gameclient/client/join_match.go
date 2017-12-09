package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/mrclayman/rest-and-go/gameclient/client/match"
	"github.com/mrclayman/rest-and-go/gameclient/client/net"
)

func getMatchID() match.Number {
	var matchID match.Number
	for {
		fmt.Print("Enter the ID of the match you wish to join: ")
		_, err := fmt.Scanf("%v", &matchID)
		if err != nil {
			fmt.Println("I need an integer")
			FlushStdin()
		} else {
			break
		}
	}
	return matchID
}

// JoinMatch queries the server for active matches and
// lets the player choose a match to join
func JoinMatch(c *http.Client, ps net.PlayerSession) error {

	matchList, err := ListMatches(c, ps)
	if err != nil {
		fmt.Println("Cannot join due to previous errors: " + err.Error())
		return err
	}

	mIDsToGTypes := collectMatchIDsAndGameTypes(matchList)

	matchID := getMatchID()
	matchGT, ok := mIDsToGTypes[matchID]
	if !ok {
		return errors.New("Cannot proceed, match ID specified does not correspond to any existing match")
	}

	joinData := map[string]interface{}{
		"player_id": ps.ID,
		"token":     ps.Token,
		"match":     match.ID{Number: matchID, Type: matchGT},
	}

	var postData []byte
	postData, err = json.Marshal(joinData)
	if err != nil {
		return err
	}

	var ms net.MatchSession
	err = net.Post(c, "/match/join", postData, &ms)
	if err != nil {
		return err
	}

	return runMatchLoop(c, ps, ms)
}

// collectMatchIDsAndGameTypes iterates through all
// ongoing matches of all types and creates a mapping
// between their ID's and their game types to make it
// easier for the user to pick one. Since game type is
// required on server's input, it is preferable to let
// the player pick just the ID and add the game type
// automatically
func collectMatchIDsAndGameTypes(ml *match.Matchlist) map[match.Number]string {
	retval := make(map[match.Number]string)
	for _, m := range ml.DM {
		retval[m.Number] = match.DeathMatch
	}

	for _, m := range ml.CTF {
		retval[m.Number] = match.CaptureTheFlag
	}

	for _, m := range ml.LMS {
		retval[m.Number] = match.LastManStanding
	}

	for _, m := range ml.Duel {
		retval[m.Number] = match.Duel
	}

	return retval
}
