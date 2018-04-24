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

// Insert Add a new player in db.
func (player *Player) Insert(dbConn *sql.DB) (int64, error) {

	var id int64

	err := dbConn.QueryRow(`
		INSERT INTO player (first_name, last_name, points)
		VALUES ($1,$2,$3) RETURNING id`,
			player.FirstName, player.LastName, player.Points).Scan(&id)

	if err != nil {
		log.Printf("Could not create a new player %v. Error: %v\n", player, err)
		return 0, err
	}

	log.Printf("Inserted player %v\n", player)

	return id, nil
}

// GetByID Fetch player by ID.
func (player *Player) GetByID(dbConn *sql.DB) error {

	err := dbConn.QueryRow(`
		SELECT id, first_name, last_name, points 
		FROM player
		WHERE id = $1`,
		&player.ID).Scan(&player.ID, &player.FirstName,
			&player.LastName, &player.Points)

	if err != nil {
		log.Printf("Could not get player %v, err: %v\n", player, err)
		return err
	}

	log.Printf("Successfully fetched player %v\n", player)

	return nil
}

// GetAll Fetch all players in DB.
func (player *Player) GetAll(dbConn *sql.DB) ([]Player, error) {

	players := make([]Player, 0, 32)
	rows, err := dbConn.Query(`
		SELECT id, first_name, last_name, points
		FROM player`)

	if err != nil {
		log.Printf("Could not fetch all players in DB. Error: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var player Player
		rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Points)
		players = append(players, player)
	}

	log.Print("Successfully fetched all players\n")

	return players, nil
}

// Update Update the given player in database.
func (player *Player) Update(dbConn *sql.DB) (int64, error) {

	result, err := dbConn.Exec(`
		UPDATE player 
		SET first_name = $1, last_name = $2, points = $3
		WHERE id = $4`,
		player.FirstName, player.LastName, player.Points, player.ID)

	if err != nil {
		log.Printf("Could not update player in DB. Player: %v / Error: %v\n",
			player, err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Could not get the amount of affected rows. Error: %v\n",
			err)
		return 0, err
	}

	log.Printf("Successfully updated player %v\n", player)

	return rowsAffected, nil
}
