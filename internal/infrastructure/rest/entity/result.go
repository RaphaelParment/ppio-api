package entity

type Result struct {
	WinnerID     int  `json:"winner_id"`
	LoserRetired bool `json:"loser_retired"`
}
