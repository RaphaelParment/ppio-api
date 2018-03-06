package models

import (
	"database/sql"
	"log"
)

type Player struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Points    int    `json:"points"`
}

func (player *Player) Insert(dbConn *sql.DB) int64 {

	var id int64

	err := dbConn.QueryRow("INSERT INTO player (first_name, last_name, points) VALUES ($1,$2,$3) RETURNING id", player.FirstName, player.LastName, player.Points).Scan(&id)

	if err != nil {
		log.Fatalf("Could not create a new player %v. Error: %v\n", player, err)
	}

	return id
}
