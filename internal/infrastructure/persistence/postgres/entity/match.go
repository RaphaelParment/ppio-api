package entity

import (
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	"time"
)

type Match struct {
	Id          int       `json:"id"`
	PlayerOneId int       `json:"player_one_id"`
	PlayerTwoId int       `json:"player_two_id"`
	Datetime    time.Time `json:"date_time"`
}

func MatchToJSON(match matchModel.Match) Match {
	return Match{
		Id:          match.Id().Int(),
		PlayerOneId: match.PlayerOneId().Int(),
		PlayerTwoId: match.PlayerTwoId().Int(),
		Datetime:    match.Datetime(),
	}
}
