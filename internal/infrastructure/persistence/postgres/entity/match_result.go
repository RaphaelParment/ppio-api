package entity

type MatchResult struct {
	MatchID      int  `db:"match_id"`
	WinnerID     int  `db:"winner_id"`
	LoserRetired bool `db:"loser_retired"`
}
