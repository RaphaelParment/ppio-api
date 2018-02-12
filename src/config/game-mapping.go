package config

const (
	GameMapping = `
{
	"settings": {
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings": {
		"doc": {
			"properties": {
				"date": {
					"type": "date",
					"format" : "text"
				},
				"player1": {
					"type": "Player"
				},
				"player2": {
					"type": "Player"
				},
				"score1": {
					"type": "integer"
				},
				"score2": {
					"type": "integer"
				}
			}
		}
	}
}
`
)