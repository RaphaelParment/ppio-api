package models

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const playerTable = "player p"

// Player Model of the player as stored in the database.
type Player struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Points    int    `json:"points"`
}

func preparePlayerCountQuery() bytes.Buffer {
	var queryBlder bytes.Buffer
	queryBlder.WriteString(startQueryCount)
	queryBlder.WriteString(playerTable)
	return queryBlder
}

func preparePlayerQuery() bytes.Buffer {
	var queryBlder bytes.Buffer
	queryBlder.WriteString(startPlayerQuery)
	queryBlder.WriteString(playerTable)
	return queryBlder
}

func preparePlayerWhereClause(filter map[string]interface{}, queryBld *bytes.Buffer, params *[]interface{}, isCount bool) error {
	placeHolderCnt := 1

	queryBld.WriteString(" WHERE 1=1 ")
	firstName, ok := filter["firstName"]
	if ok {
		queryBld.WriteString(fmt.Sprintf(" AND first_name = $%d", placeHolderCnt))
		placeHolderCnt = placeHolderCnt + 1
		*params = append(*params, firstName)
	}
	lastName, ok := filter["lastName"]
	if ok {
		queryBld.WriteString(fmt.Sprintf(" AND last_name = $%d", placeHolderCnt))
		placeHolderCnt = placeHolderCnt + 1
		*params = append(*params, lastName)
	}

	points, ok := filter["points"]
	if ok {
		queryBld.WriteString(fmt.Sprintf(" AND points >= $%d", placeHolderCnt))
		placeHolderCnt = placeHolderCnt + 1
		*params = append(*params, points.(int))
	}

	if !isCount {
		queryBld.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", placeHolderCnt, placeHolderCnt+1))
		*params = append(*params, filter["limit"].(int))
		*params = append(*params, filter["offset"].(int))
	}

	return nil
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
func (player *Player) GetAll(dbConn *sql.DB, filters map[string]interface{}) ([]Player, int64, error) {

	var countRows int64
	players := make([]Player, 0, 64)
	params := make([]interface{}, 0, 8)

	queryBlder := preparePlayerCountQuery()
	err := preparePlayerWhereClause(filters, &queryBlder, &params, true)
	if err != nil {
		return nil, 0, err
	}

	row := dbConn.QueryRow(queryBlder.String(), params...)
	if row == nil {
		log.Printf("Could not fetch the count of games")
		return nil, 0, errors.New("Count of games not be fetched")
	}
	if err := row.Scan(&countRows); err != nil {
		log.Printf("Could not scan the count of games from database. Error: %v", err)
		return nil, 0, err
	}

	queryBlder = preparePlayerQuery()
	params = make([]interface{}, 0, 8)
	err = preparePlayerWhereClause(filters, &queryBlder, &params, false)
	if err != nil {
		return nil, 0, err
	}

	rows, err := dbConn.Query(queryBlder.String(), params...)
	if err != nil {
		log.Printf("Could not fetch all games in DB. Error: %v", err)
		return nil, 0, err
	}
	for rows.Next() {
		var player Player
		rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Points)
		players = append(players, player)

	}

	log.Print("Successfully fetched all players\n")

	return players, countRows, nil
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
