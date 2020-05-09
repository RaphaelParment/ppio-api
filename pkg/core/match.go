package core

import (
	"encoding/json"
	"io"
	"time"

	"github.com/go-playground/validator"
)

// KeyMatch is used as request context key
type KeyMatch struct{}

// Match contains information about a match
type Match struct {
	ID             int       `json:"id"`
	FirstPlayerID  int       `json:"first_player_id" validate:"required"`
	SecondPlayerID int       `json:"second_player_id" validate:"required"`
	Datetime       time.Time `json:"date_time" validate:"required"`
}

// Matches represents multiple matches
type Matches []*Match

// ToJSON returns the json representation of the match
func (m *Match) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

// FromJSON returns match representation of <r>
func (m *Match) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(m)
}

// Validate checks the match contains correct field values
func (m *Match) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

// GetKey returns the request context key
func (m *Match) GetKey() interface{} {
	return KeyMatch{}
}

// ToJSON returns the json representation of multiple matches
func (m *Matches) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}
