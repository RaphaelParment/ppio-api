# PPio RESTful API
RESTful API written in Golang with Gorilla MUX package. 

## Endpoints

| URI            | Method   | Description                                          |
| :------------- | :------- | :--------------------------------------------------- |
| `/players`     | `GET`    | Get the full list of players.                        |
| `/players`     | `POST`   | Create a new player and persist it                   |
| `/players/:id` | `GET`    | Get the player indexed by `{id}`                     |
| `/players/:id` | `PUT`    | Update the player indexed by `{id}`                  |
| `/players/:id` | `DELETE` | Delete the player indexed by `{id}`.                 |
| `/matches`     | `GET`    | Get the full list of games.                          |
| `/matches`     | `POST`   | Create a new match.                                  |
| `/results/:id` | `GET`    | Get the match result for match indexed by `{id}`.    |
| `/results`     | `POST`   | Create the match result.                             |
| `/scores/:id`  | `GET`    | Get the match games scores for match indexed `{id}`. |
| `/scores`      | `POST`   | Create the match games scores.                       |