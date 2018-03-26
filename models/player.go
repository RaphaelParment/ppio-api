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

	log.Printf("Inserted player with ID '%d'", id)

	return id
}

// GetByID Fetch player by ID.
func (player *Player) GetByID(dbConn *sql.DB) error {

	err := dbConn.QueryRow("SELECT id, first_name, last_name, points FROM player WHERE id = $1",
		&player.ID).Scan(&player.ID, &player.FirstName, &player.LastName, &player.Points)

	if err != nil {
		log.Fatalf("Could not get game %v, err: %v", player, err)
	}

	return nil
}

// GetAll Fetch all players in DB.
func (player *Player) GetAll(dbConn *sql.DB) ([]Player, error) {

	players := make([]Player, 0, 32)
	rows, err := dbConn.Query("SELECT id, first_name, last_name, points FROM player")
	if err != nil {
		log.Printf("Could not fetch all players in DB. Error: %v", err)
		return nil, err
	}
	for rows.Next() {
		var player Player
		rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Points)
		players = append(players, player)
	}

	return players, nil
}
