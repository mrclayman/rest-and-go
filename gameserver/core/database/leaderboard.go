package database

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/leaderboard"
	"github.com/mrclayman/rest-and-go/gameserver/core/match"
	"gopkg.in/mgo.v2/bson"
)

var (
	dmLeaderboardSortingCriterion   = []string{"-kills", "deaths"}
	ctfLeaderboardSortingCriterion  = []string{"-captures", "-kills", "deaths"}
	lmsLeaderboardSortingCriterion  = []string{"-wins", "-kills", "deaths"}
	duelLeaderboardSortingCriterion = lmsLeaderboardSortingCriterion
)

// GetDMLeaderboard returns the leaderboard
// associated with the DeathMatch game type
func (db *Database) GetDMLeaderboard() ([]leaderboard.DMLeaderboardRecord, error) {
	lbName := db.leaderboardCollPrefix + match.GameTypeToString(match.DeathMatch)

	c := db.session.DB(db.dbName).C(lbName)
	q := c.Find(bson.M{}).Sort(dmLeaderboardSortingCriterion...)

	r := make(leaderboard.DMLeaderboard, 10)
	if err := q.All(&r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetCTFLeaderboard returns the leaderboard
// associated with the CTF game type
func (db *Database) GetCTFLeaderboard() (*leaderboard.CTFLeaderboard, error) {
	lbName := db.leaderboardCollPrefix + match.GameTypeToString(match.CaptureTheFlag)

	c := db.session.DB(db.dbName).C(lbName)
	q := c.Find(bson.M{}).Sort(ctfLeaderboardSortingCriterion...)

	r := make(leaderboard.CTFLeaderboard, 10)
	if err := q.All(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

// GetLMSLeaderboard returns the leaderboard
// associated with the LMS game type
func (db *Database) GetLMSLeaderboard() (*leaderboard.LMSLeaderboard, error) {
	lbName := db.leaderboardCollPrefix + match.GameTypeToString(match.LastManStanding)

	c := db.session.DB(db.dbName).C(lbName)
	q := c.Find(bson.M{}).Sort(lmsLeaderboardSortingCriterion...)

	r := make(leaderboard.LMSLeaderboard, 10)
	if err := q.All(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

// GetDuelLeaderboard returns the leaderboard
// associated with the Duel game type
func (db *Database) GetDuelLeaderboard() (*leaderboard.DuelLeaderboard, error) {
	lbName := db.leaderboardCollPrefix + match.GameTypeToString(match.Duel)

	c := db.session.DB(db.dbName).C(lbName)
	q := c.Find(bson.M{}).Sort(duelLeaderboardSortingCriterion...)

	r := make(leaderboard.DuelLeaderboard, 10)
	if err := q.All(&r); err != nil {
		return nil, err
	}
	return &r, nil
}
