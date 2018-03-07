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

func GenerateGames(players []*models.Player) []models.Game {

	var games []models.Game

	for _, homePlayer := range players {
		for _, awayPlayer := range players {

			// generate a game
			if homePlayer.FirstName != awayPlayer.FirstName {

				// generate scores
				homeScore := rand.Intn(20)
				awayScore := rand.Intn(20)

				if homeScore >= awayScore {
					if awayScore > 15 {
						awayScore = homeScore - 2
					} else {
						homeScore = 11
						awayScore = rand.Intn(9) + 1
					}
				} else {
					if homeScore > 15 {
						homeScore = awayScore - 2
					} else {
						awayScore = 11
						homeScore = rand.Intn(9) + 1
					}
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
					DateTime: objDatetime,
					Player1:  *homePlayer,
					Player2:  *awayPlayer,
					Score1:   homeScore,
					Score2:   awayScore,
				}

				games = append(games, game)
			}
		}
	}

	return games
}
