package entity

type MatchResult struct {
	MatchID      int  `json:"match_id"`
	WinnerID     int  `json:"winner_id"`
	GamesPlayed  int  `json:"games_played"`
	LoserRetired bool `json:"loser_retired"`
}
