package model

import (
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type MatchPatch struct {
	playerOneId *playerModel.Id
	playerTwoId *playerModel.Id
	result      *Result
	score       *Score
	datetime    *time.Time
}

func NewMatchPatch(
	playerOneId *playerModel.Id,
	playerTwoId *playerModel.Id,
	result *Result,
	score *Score,
	datetime *time.Time,
) MatchPatch {
	return MatchPatch{playerOneId: playerOneId, playerTwoId: playerTwoId, result: result, score: score, datetime: datetime}
}

func (m MatchPatch) PlayerOneId() *playerModel.Id {
	return m.playerOneId
}

func (m MatchPatch) PlayerTwoId() *playerModel.Id {
	return m.playerTwoId
}

func (m MatchPatch) Result() *Result {
	return m.result
}

func (m MatchPatch) Score() *Score {
	return m.score
}

func (m MatchPatch) Datetime() *time.Time {
	return m.datetime
}
