# PPio RESTful API
RESTful API written in Golang with Gorilla MUX package. 

## Endpoints
All the endpoints are prefixed by /ppio. 

|      URI       |  Method  |                  Description                  |
| :------------- | :------- | :-------------------------------------------- |
| `/players`     | `GET`    | Get the full list of players.                 |
| `/players`     | `POST`   | Create a new player and persist it            |
| `/players/:id` | `GET`    | Get the player indexed by `{id}`              |
| `/players/:id` | `PUT`    | Update the player indexed by `{id}`           |
| `/players/:id` | `DELETE` | Delete the player indexed by `{id}`. __TODO__ |
| `/games`       | `GET`    | Get the full list of games.                   |
| `/games`       | `POST`   | Create a new game and persist it.             |
| `/games/:id`   | `GET`    | Get the game indexed by `{id}`.               |
| `/games/:id`   | `PUT`    | Update the game indexed by `{id}`.            |
| `/games/:id`   | `DELETE` | Delete the game indexed by `{id}`.            |
