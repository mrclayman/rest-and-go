package match

import (
	"errors"

	"github.com/mrclayman/rest-and-go/gameclient/client/player"
	"github.com/mrclayman/rest-and-go/gameclient/client/shared"
)

// CTFPlayerRank defines the structure of a player's
// rank in the CTF game type
type CTFPlayerRank struct {
	Player   player.Player `json:"player"`
	Kills    uint          `json:"kills"`
	Deaths   uint          `json:"deaths"`
	Captures uint          `json:"captures"`
}

// CTFPlayerRanks is a slice of CTF game type
// player rank objects
type CTFPlayerRanks []CTFPlayerRank

// CTFMatch contains information on
// a CTF type match
type CTFMatch struct {
	Number Number         `json:"match_id"`
	Ranks  CTFPlayerRanks `json:"ranks"`
}

// CTFMatches defines a slice of CTF-type matches
type CTFMatches []*CTFMatch

// unmarshalCTFMatch unmarshals the contents of the
// input map into an instance of CTFMatch
func unmarshalCTFMatch(in map[string]interface{}) (*CTFMatch, error) {
	mID, err := IDFromMap(in)
	if err != nil {
		return nil, err
	}

	var playerRanks map[string]interface{}
	playerRanks, err = shared.RanksFromMap(in)
	if err != nil {
		return nil, err
	}

	var ranks CTFPlayerRanks
	for _, r := range playerRanks {
		rm, ok := r.(map[string]interface{})
		if !ok {
			return nil, errors.New("Rank record not a map of values")
		}

		rank, err := unmarshalCTFRankRecord(rm)
		if err != nil {
			return nil, err
		}

		ranks = append(ranks, rank)
	}

	return &CTFMatch{
		Number: Number(mID),
		Ranks:  ranks,
	}, nil
}

// unmarshalCTFRankRecord unmarshals the contents
// of the input map into an instance of
// the CTFPlayerRank structure
func unmarshalCTFRankRecord(in map[string]interface{}) (CTFPlayerRank, error) {
	player, err := player.FromMap(in)
	if err != nil {
		return CTFPlayerRank{}, err
	}

	var kills uint
	kills, err = shared.KillCountFromMap(in)
	if err != nil {
		return CTFPlayerRank{}, err
	}

	var deaths uint
	deaths, err = shared.DeathCountFromMap(in)
	if err != nil {
		return CTFPlayerRank{}, err
	}

	var captures uint
	captures, err = shared.CaptureCountFromMap(in)
	if err != nil {
		return CTFPlayerRank{}, err
	}

	return CTFPlayerRank{
		Player:   player,
		Kills:    kills,
		Deaths:   deaths,
		Captures: captures,
	}, nil
}
