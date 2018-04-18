package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"ppio/models"
	"time"
)

func GetPlayers() []*models.Player {
	raw, err := ioutil.ReadFile("/var/run/ppio/data/dummy-players.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []*models.Player
	err = json.Unmarshal(raw, &c)
	if err != nil {
		log.Fatalf("Error occured when unmarshalling the dummy players. Error :%v", err)
	}
	return c
}

// winner 1 -> player1 wins
// winner 2 -> player2 wins
func generateSet(winner int) models.Set {

	var player1Score int
	var player2Score int

	if winner == 1 {
		player1Score = 11
		player2Score = rand.Intn(10)
	} else {
		player2Score = 11
		player1Score = rand.Intn(10)
	}

	set := models.Set{
		Score1: player1Score,
		Score2: player2Score,
	}

	return set
}

func GenerateGames(players []*models.Player) []models.Game {

	var games []models.Game
	validated := false

	for _, homePlayer := range players {
		for _, awayPlayer := range players {

			// generate a game
			if homePlayer.FirstName != awayPlayer.FirstName {

				var winner int
				var winnerID int64
				var editedByID int64
				var sets []models.Set
				var numberOfSets int

				if rand.Float32() > 0.5 {
					numberOfSets = 2
					editedByID = homePlayer.ID
				} else {
					numberOfSets = 3
					editedByID = awayPlayer.ID
				}

				if rand.Float32() > 0.5 {
					winner = 1
					winnerID = homePlayer.ID
				} else {
					winner = 2
					winnerID = awayPlayer.ID
				}

				for i := 0; i < numberOfSets; i++ {

					var set models.Set

					if numberOfSets == 2 {
						set = generateSet(winner)
					} else {
						set = generateSet((winner % 2) + 1) // winner -> 1 then 2, winner 2 -> then 1
						winner++
					}

					sets = append(sets, set)
				}

				// generate random datetime (1 month span)
				day := rand.Intn(28) + 1
				hour := rand.Intn(24)
				minute := rand.Intn(60)

				datetime := fmt.Sprintf("2018-02-%02d %02d:%02d:00",
					day, hour, minute)
				objDatetime, err := time.Parse("2006-01-02 15:04:05", datetime)
				if err != nil {
					log.Fatalf("Could not parse date time %s. Error: %v\n", datetime, err)
				}
				game := models.Game{
					DateTime:   objDatetime,
					Player1ID:  homePlayer.ID,
					Player2ID:  awayPlayer.ID,
					WinnerID:   winnerID,
					Validated:  validated,
					EditedByID: editedByID,
					Sets:       sets,
				}

				if !validated {
					validated = true
				} else {
					validated = false
				}

				games = append(games, game)
			}
		}
	}

	return games
}
