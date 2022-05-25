package entity

import (
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	"time"
)

type Match struct {
	Id             int       `json:"id"`
	FirstPlayerId  int       `json:"first_player_id"`
	SecondPlayerId int       `json:"second_player_id"`
	Datetime       time.Time `json:"date_time"`
}

func MatchToJSON(match matchModel.Match) Match {
	return Match{
		Id:             int(match.Id),
		FirstPlayerId:  int(match.FirstPlayerId),
		SecondPlayerId: int(match.SecondPlayerId),
		Datetime:       match.Datetime,
	}
}
