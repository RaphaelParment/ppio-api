package entity

import (
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type Match struct {
	Id          int       `json:"id"`
	PlayerOneId int       `json:"player_one_id"`
	PlayerTwoId int       `json:"player_two_id"`
	Datetime    time.Time `json:"date_time"`
}

func GameToJSON(match matchModel.Game) Match {
	return Match{
		Id:          int(match.Id),
		PlayerOneId: int(match.PlayerOneId),
		PlayerTwoId: int(match.PlayerTwoId),
		Datetime:    match.Datetime,
	}
}

func MatchFromJSON(match Match) matchModel.Game {
	return matchModel.Game{
		Id:          matchModel.Id(match.Id),
		PlayerOneId: playerModel.Id(match.PlayerOneId),
		PlayerTwoId: playerModel.Id(match.PlayerTwoId),
		Datetime:    match.Datetime,
	}
}
