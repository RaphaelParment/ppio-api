package core

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

// KeyMatchGamesScores is used as request context key
type KeyMatchGamesScores struct{}

// GameScore contains information about a single match's gane
type GameScore struct {
	MatchID           int `json:"match_id" validate:"required"`
	GameNbr           int `json:"game_number" validate:"required"`
	FirstPlayerScore  int `json:"first_player_score" validate:"required"`
	SecondPlayerScore int `json:"second_player_score" validate:"required"`
}

// GameScores represents multiple game score
type GameScores []GameScore

// Validate checks the match contains correct field values
func (g *GameScore) Validate() error {
	validate := validator.New()
	return validate.Struct(g)
}

// FromJSON returns match representation of <r>
func (gs *GameScores) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(gs)
}

// ToJSON returns match representation of <r>
func (gs *GameScores) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(gs)
}

// Validate checks the match contains correct field values
func (gs GameScores) Validate() error {
	for _, score := range gs {
		validate := validator.New()
		return validate.Struct(score)
	}
	return nil
}

// GetKey returns the context key
func (gs GameScores) GetKey() interface{} {
	return KeyMatchGamesScores{}
}
