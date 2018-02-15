package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"ppio/src/models"
	"math/rand"
)

func GetPlayers() []models.Player {
	raw, err := ioutil.ReadFile("data/dummy-players.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []models.Player
	json.Unmarshal(raw, &c)
	return c
}

func GenerateGames(players []models.Player) []models.Game {

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
				day := rand.Intn(31) + 1
				hour := rand.Intn(24)
				minute := rand.Intn(60)

				datetime := fmt.Sprintf("2018-02-%d %d:%d:00",
					day, hour, minute)

				game := models.Game{
					DateTime: datetime,
					Player1: homePlayer,
					Player2: awayPlayer,
					Score1: homeScore,
					Score2: awayScore,
				}

				games = append(games, game)
			}
		}
	}

	return games
}