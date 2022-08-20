package model

import (
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type Game struct {
	Id          Id
	PlayerOneId playerModel.Id
	PlayerTwoId playerModel.Id
	Datetime    time.Time
}
