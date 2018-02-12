package config

const PlayerMapping = `
{
	"settings": {
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings": {
		"doc": {
			"properties": {
				"firstName": {
					"type": "text"
				},
				"lastName": {
					"type": "text"
				},
				"points": {
					"type": "integer"
				}
			}
		}
	}
}`
