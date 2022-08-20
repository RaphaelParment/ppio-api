package entity

type MatchResult struct {
	GameID       int  `json:"game_id"`
	WinnerID     int  `json:"winner_id"`
	LoserRetired bool `json:"loser_retired"`
}
