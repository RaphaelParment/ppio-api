package entity

import (
	"time"
)

type Match struct {
	Id          int       `db:"id"`
	PlayerOneId int       `db:"player_one_id"`
	PlayerTwoId int       `db:"player_two_id"`
	Datetime    time.Time `db:"date_time"`
}
