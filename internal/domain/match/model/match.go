package model

import (
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type Match struct {
	Id          Id
	PlayerOneId playerModel.Id
	PlayerTwoId playerModel.Id
	Datetime    time.Time
}
