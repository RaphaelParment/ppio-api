package core

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

// KeyPlayer is used as request context key
type KeyPlayer struct{}

// Player contains the information of a player
type Player struct {
	// The id of the player
	// required: true
	// min: 1
	ID int `json:"id"`
	// The first name of the player
	// required: true
	FirstName string `json:"first_name" validate:"required"`
	// The last name of the player
	// required: true
	LastName string `json:"last_name" validate:"required"`
	// The email of the player
	// required: true
	Email string `json:"email" validate:"email"`
	// The number of points
	// min: 1
	// max: 100
	Points int `json:"points" validate:"lt=101"`
}

// Players represents multiple players
type Players []*Player

// Validate checks the player contains the correct fields
func (p *Player) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

// FromJSON returns player representation of <r>
func (p *Player) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// ToJSON return the json representation of <p>
func (p *Player) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetKey returns the request context key
func (p *Player) GetKey() interface{} {
	return KeyPlayer{}
}

// ToJSON return the json representation of <ps>
func (ps *Players) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ps)
}
