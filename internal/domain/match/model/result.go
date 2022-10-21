package model

type Result struct {
	MatchID      int  `json:"match_id" validate:"required"`
	WinnerID     int  `json:"winner_id" validate:"required"`
	LoserRetired bool `json:"loser_retired"`
}
