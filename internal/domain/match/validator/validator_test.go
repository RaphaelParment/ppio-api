package validator

import (
	"github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCanValidateMatch(t *testing.T) {
	datetime, _ := time.Parse(time.DateTime, "2006-01-02 15:04:05")
	tt := []struct {
		name             string
		match            model.Match
		expectedProblems map[string]string
	}{
		{
			name: "same two players",
			match: model.NewMatch(
				model.NewUndefinedId(),
				playerModel.Id(1),
				playerModel.Id(1),
				model.NewResult(playerModel.Id(1), false),
				model.Score{
					model.NewSet(11, 9),
				},
				datetime,
			),
			expectedProblems: map[string]string{
				"player-one-id": "player one cannot be the same as player two",
				"player-two-id": "player two cannot be the same as player one",
			},
		},
		{
			name: "players have same amounts of points in one set",
			match: model.NewMatch(
				model.NewUndefinedId(),
				playerModel.Id(1),
				playerModel.Id(2),
				model.NewResult(playerModel.Id(1), false),
				model.Score{
					model.NewSet(11, 9),
					model.NewSet(11, 11),
				},
				datetime,
			),
			expectedProblems: map[string]string{
				"set-score-1": "player one score cannot equal player two score",
			},
		},
		{
			name: "players have won same number of sets",
			match: model.NewMatch(
				model.NewUndefinedId(),
				playerModel.Id(1),
				playerModel.Id(2),
				model.NewResult(playerModel.NewUndefinedId(), false),
				model.Score{
					model.NewSet(11, 9),
					model.NewSet(8, 11),
				},
				datetime,
			),
			expectedProblems: map[string]string{
				"match-score": "player one cannot have won the same number of sets as player two",
			},
		},
		{
			name: "player one marked as winner but has not won most sets",
			match: model.NewMatch(
				model.NewUndefinedId(),
				playerModel.Id(1),
				playerModel.Id(2),
				model.NewResult(playerModel.Id(1), false),
				model.Score{
					model.NewSet(11, 9),
					model.NewSet(8, 11),
					model.NewSet(6, 11),
				},
				datetime,
			),
			expectedProblems: map[string]string{
				"match-winner": "player two has won most sets but it not marked as winner",
			},
		},
		{
			name: "player two marked as winner but has not won most sets",
			match: model.NewMatch(
				model.NewUndefinedId(),
				playerModel.Id(1),
				playerModel.Id(2),
				model.NewResult(playerModel.Id(2), false),
				model.Score{
					model.NewSet(11, 9),
					model.NewSet(8, 11),
					model.NewSet(11, 4),
				},
				datetime,
			),
			expectedProblems: map[string]string{
				"match-winner": "player one has won most sets but it not marked as winner",
			},
		},
	}

	validator := NewMatchValidator()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			problems := validator.ValidateMatch(tc.match)
			assert.ObjectsAreEqual(tc.expectedProblems, problems)
			assert.Equal(t, tc.expectedProblems, problems)
		})
	}
}
