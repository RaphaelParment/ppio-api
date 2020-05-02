package core

import (
	"encoding/json"
	"io"
	"time"
)

// Match contains information about a match between
// a home player and an away player
type Match struct {
	ID              int
	HomePlayerID    int
	AwayPlayerID    int
	HomePlayerScore int
	AwayPlayerScore int
	Datetime        time.Time
}

type Matches []*Match

// ToJSON returns the json representation of the match
func (m *Match) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *Matches) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}
