package models

type Player struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Points    int    `json:"points"`
}