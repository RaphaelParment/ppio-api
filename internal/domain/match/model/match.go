package model

import (
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type Match struct {
	Id             Id
	FirstPlayerId  playerModel.Id
	SecondPlayerId playerModel.Id
	Datetime       time.Time
}
