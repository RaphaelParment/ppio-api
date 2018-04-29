package models

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const gameTable = "game g"

// Game structure
type Game struct {
	ID              int64     `json:"id,omitempty"`
	DateTime        time.Time `json:"datetime,omitempty"`
	Player1ID       int64     `json:"player1Id"`
	Player2ID       int64     `json:"player2Id"`
	WinnerID        int64     `json:"winnerId"`
	ValidationState int       `json:"validationState"`
	EditedByID      int64     `json:"editedById"`
	Sets            []Set     `json:"sets"`
}

func prepareGameCountQuery() bytes.Buffer {
	var queryBlder bytes.Buffer
	queryBlder.WriteString(startQueryCount)
	queryBlder.WriteString(gameTable)
	return queryBlder
}

func prepareGameQuery() bytes.Buffer {
	var queryBlder bytes.Buffer
	queryBlder.WriteString(startGameQuery)
	queryBlder.WriteString(gameTable)
	return queryBlder
}

func prepareGameWhereClause(filter map[string]interface{}, queryBld *bytes.Buffer, params *[]interface{}, isCount bool) error {
	placeHolderCnt := 1
	playerFilter, ok := filter["playerFirstName"]
	playerClause := ok
	if ok {
		queryBld.WriteString(" JOIN player p1 ON g.player1_id = p1.id JOIN player p2 ON g.player2_id = p2.id")
	}
	queryBld.WriteString(" WHERE 1=1 ")
	validatedFilter, ok := filter["validated"]
	if ok {
		validatedValue := 0
		if validatedFilter.(bool) {
			validatedValue = 1
		}
		queryBld.WriteString(fmt.Sprintf(" AND validation_state = $%d", placeHolderCnt))
		placeHolderCnt = placeHolderCnt + 1
		*params = append(*params, validatedValue)
	}

	if playerClause {
		queryBld.WriteString(fmt.Sprintf(" AND (p2.first_name = $%d OR p1.first_name = $%d)", placeHolderCnt, placeHolderCnt+1))
		placeHolderCnt = placeHolderCnt + 2
		*params = append(*params, playerFilter)
		*params = append(*params, playerFilter)
	}

	if !isCount {
		queryBld.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", placeHolderCnt, placeHolderCnt+1))
		*params = append(*params, filter["limit"].(int))
		*params = append(*params, filter["offset"].(int))
	}

	return nil
}

// Insert game
func (game *Game) Insert(dbConn *sql.DB) (int64, error) {

	var id int64
	err := dbConn.QueryRow(`
		INSERT INTO game (player1_id, player2_id, winner_id, validation_state,
		edited_by_id, datetime)
		VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		game.Player1ID, game.Player2ID, game.WinnerID,
		game.ValidationState, game.EditedByID, game.DateTime).Scan(&id)

	if err != nil {
		log.Printf("Could not insert game %v. Error: %v\n", game, err)
		return 0, err
	}

	for _, set := range game.Sets {

		_, err := dbConn.Exec(`
		INSERT INTO set (game_id, score1, score2)
		VALUES ($1, $2, $3)`, id, set.Score1, set.Score2)

		if err != nil {
			log.Printf("Could not insert set %v for game %v. Error: %v\n",
				set, game, err)
			return 0, err
		}
	}

	log.Printf("Inserted game with ID %v\n", game)

	return id, nil
}

// GetByID Returns a game with its corresponding sets.
func (game *Game) GetByID(dbConn *sql.DB) error {

	// Get all sets for a given game
	var sets []Set

	err := dbConn.QueryRow(`
		SELECT id, player1_id, player2_id, winner_id, validation_state,
		edited_by_id, datetime
		FROM game
		WHERE id = $1`,
		&game.ID).
		Scan(&game.ID, &game.Player1ID, &game.Player2ID,
			&game.WinnerID, &game.ValidationState, &game.EditedByID,
			&game.DateTime)

	if err != nil {
		log.Printf("Could not get game %v, err: %v\n", game, err)
		return err
	}

	sets, err = game.GetSets(dbConn)

	if err != nil {
		return err
	}

	game.Sets = sets

	log.Printf("Successfully fetched game %v\n", game)

	return nil
}

// GetAll Returns all games and the total count.
func (game *Game) GetAll(dbConn *sql.DB, filters map[string]interface{}) ([]Game, int64, error) {

	var countRows int64
	games := make([]Game, 0, 512)
	params := make([]interface{}, 0, 8)

	queryBlder := prepareGameCountQuery()
	err := prepareGameWhereClause(filters, &queryBlder, &params, true)
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

	queryBlder = prepareGameQuery()
	params = make([]interface{}, 0, 8)
	err = prepareGameWhereClause(filters, &queryBlder, &params, false)
	if err != nil {
		return nil, 0, err
	}

	rows, err := dbConn.Query(queryBlder.String(), params...)
	if err != nil {
		log.Printf("Could not fetch all games in DB. Error: %v", err)
		return nil, 0, err
	}
	for rows.Next() {
		var game Game
		rows.Scan(&game.ID, &game.Player1ID, &game.Player2ID,
			&game.WinnerID, &game.ValidationState, &game.EditedByID,
			&game.DateTime)
		sets, err := game.GetSets(dbConn)

		if err != nil {
			return nil, 0, err
		}

		game.Sets = sets
		games = append(games, game)
	}

	return games, countRows, nil
}

// Update updates a given game
func (game *Game) Update(dbConn *sql.DB) (int64, error) {

	result, err := dbConn.Exec(`
		UPDATE game SET player1_id = $1, player2_id = $2, winner_id = $3,
		validation_state = $4, edited_by_id = $5, datetime = $6
		WHERE id = $7`,
		game.Player1ID, game.Player2ID, game.WinnerID, game.ValidationState,
		game.EditedByID, game.DateTime, game.ID)

	if err != nil {
		log.Printf("Could not update game: %v, err: %v\n", game, err)
		return 0, err
	}

	for _, set := range game.Sets {
		result, err = dbConn.Exec(`
			UPDATE set SET score1 = $1, score2 = $2
			WHERE id = $3`,
			set.Score1, set.Score2, set.ID)

		if err != nil {
			log.Printf("Could not update set: %v, err: %v\n", set, err)
			return 0, err
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Could not get the amount of affected rows. Error: %v\n",
			err)
		return 0, err
	}

	return rowsAffected, nil

}

// Delete deletes a game
func (game *Game) Delete(dbConn *sql.DB) (int64, error) {

	result, err := dbConn.Exec(`
		DELETE FROM game
		WHERE id = $1`, game.ID)

	if err != nil {
		log.Printf("Could not delete game: %v, err: %v\n", game, err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf(`
			Could not get the number of affected rows while
			deleting game: %v. Error: %v\n`, game, err)
		return 0, err
	}

	if rowsAffected != 1 {
		log.Printf("Delete more than 1 game... Error\n")
		err = errors.New("deleting more than 1 item")
		return 0, err
	}

	log.Printf("Successfully deleted game %v\n", game)

	return game.ID, nil
}

// GetSets returns all sets for a game
func (game *Game) GetSets(dbConn *sql.DB) ([]Set, error) {

	var sets []Set
	var set Set
	rows, err := dbConn.Query(`
		SELECT id, score1, score2 FROM set
		WHERE game_id = $1`, game.ID)

	if err != nil {
		log.Printf("Could not fetch all sets for game %v. Error: %v\n",
			game, err)
		return nil, err
	}

	for rows.Next() {
		set.GameID = game.ID
		rows.Scan(&set.ID, &set.Score1, &set.Score2)
		sets = append(sets, set)
	}

	return sets, nil
}
