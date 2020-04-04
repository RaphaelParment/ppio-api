package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator"
)

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// Player contains the information of a player
type Player struct {
	// The id of the player
	// required: true
	// min: 1
	ID int `json:"id"`
	// The first name of the player
	// required: true
	FirstName string `json:"firstName" validate:"required"`
	// The last name of the player
	// required: true
	LastName string `json:"lastName" validate:"required"`
	// The email of the player
	// required: true
	Email string `json:"email" validate:"email"`
	// The number of points
	// min: 1
	// max: 100
	Points int `json:"points" validate:"lt=101"`
}

func (p *Player) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

type Players []*Player

func (p *Player) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Players) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Player) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
