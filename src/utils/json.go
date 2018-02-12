package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"ppio-web/src/models"
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