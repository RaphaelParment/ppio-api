package entity

type Set struct {
	Id             int `db:"id"`
	PlayerOneScore int `db:"player_one_score"`
	PlayerTwoScore int `db:"player_two_score"`
}
