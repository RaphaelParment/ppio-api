package entity

type MatchScore struct {
	MatchID        int `json:"match_id" `
	GameNbr        int `json:"game_number"`
	PlayerOneScore int `json:"player_one_score"`
	PlayerTwoScore int `json:"player_two_score"`
}
