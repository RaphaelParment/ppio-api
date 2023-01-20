package model

import (
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type Match struct {
	id          Id
	playerOneId playerModel.Id
	playerTwoId playerModel.Id
	result      Result
	score       Score
	datetime    time.Time
}

func NewMatch(
	id Id,
	playerOneId playerModel.Id,
	playerTwoId playerModel.Id,
	result Result,
	score Score,
	datetime time.Time,
) Match {
	return Match{
		id:          id,
		playerOneId: playerOneId,
		playerTwoId: playerTwoId,
		result:      result,
		score:       score,
		datetime:    datetime,
	}
}

func (m Match) Id() Id {
	return m.id
}

func (m Match) PlayerOneId() playerModel.Id {
	return m.playerOneId
}

func (m Match) PlayerTwoId() playerModel.Id {
	return m.playerTwoId
}

func (m Match) Result() Result {
	return m.result
}

func (m Match) Score() Score {
	return m.score
}

func (m Match) Datetime() time.Time {
	return m.datetime
}
