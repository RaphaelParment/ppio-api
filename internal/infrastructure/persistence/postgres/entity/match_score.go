package entity

type MatchScore struct {
	MatchID           int `json:"match_id" `
	GameNbr           int `json:"game_number"`
	FirstPlayerScore  int `json:"first_player_score"`
	SecondPlayerScore int `json:"second_player_score"`
}
