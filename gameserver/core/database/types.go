package database

var (
	dmLeaderboardSortingCriterion   = []string{"-kills", "deaths"}
	ctfLeaderboardSortingCriterion  = []string{"-captures", "-kills", "deaths"}
	lmsLeaderboardSortingCriterion  = []string{"-wins", "-kills", "deaths"}
	duelLeaderboardSortingCriterion = lmsLeaderboardSortingCriterion
)
