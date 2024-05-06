package validator

import (
	"fmt"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
)

type matchValidator struct {
}

func NewMatchValidator() *matchValidator {
	return &matchValidator{}
}

func (v *matchValidator) ValidateMatch(match matchModel.Match) error {
	problems := make(map[string]string)

	if match.PlayerOneId() == match.PlayerTwoId() {
		problems["player-one-id"] = "player one cannot be the same as player two"
		problems["player-two-id"] = "player two cannot be the same as player one"
	}

	// Result should still make sense
	winnerId := match.Result().WinnerID()

	playerOneWonSets := 0
	playerTwoWonSets := 0
	for setIdx, set := range match.Score() {
		if set.PlayerOneScore() == set.PlayerTwoScore() {
			problems[fmt.Sprintf("set-score-%d", setIdx)] = "player one score cannot equal player two score"
			continue
		}

		if set.PlayerOneScore() > set.PlayerTwoScore() {
			playerOneWonSets++
		} else {
			playerTwoWonSets++
		}
	}

	if playerOneWonSets == playerTwoWonSets {
		problems["match-score"] = "player one cannot have won the same number of sets as player two"
	} else {
		if match.Result().LoserRetired() == false {
			if playerOneWonSets > playerTwoWonSets && winnerId != match.PlayerOneId() {
				problems["match-winner"] = "player one has won most sets but it not marked as winner"
			}

			if playerTwoWonSets > playerOneWonSets && winnerId != match.PlayerTwoId() {
				problems["match-winner"] = "player two has won most sets but it not marked as winner"
			}
		}
	}

	return NewValidatorError(problems)
}
