package entity

type Score []Set

type Set struct {
	PlayerOneScore int `json:"player_one_score"`
	PlayerTwoScore int `json:"player_two_score"`
}
