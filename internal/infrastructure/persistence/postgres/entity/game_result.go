package entity

type MatchResult struct {
	MatchID      int  `json:"match_id"`
	WinnerID     int  `json:"winner_id"`
	LoserRetired bool `json:"loser_retired"`
}
