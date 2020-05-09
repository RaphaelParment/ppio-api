package core

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

// KeyMatchResult is used as request context key
type KeyMatchResult struct{}

// MatchResult contains the result of a match
type MatchResult struct {
	MatchID      int  `json:"match_id" validate:"required"`
	WinnerID     int  `json:"winner_id" validate:"required"`
	GamesPlayed  int  `json:"games_played" validate:"required"`
	LoserRetired bool `json:"loser_retired"`
}

// ToJSON returns the json representation of the match
func (m *MatchResult) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

// FromJSON returns match representation of <r>
func (m *MatchResult) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(m)
}

// Validate checks the match contains correct field values
func (m *MatchResult) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

// GetKey returns the request context key
func (m *MatchResult) GetKey() interface{} {
	return KeyMatchResult{}
}
