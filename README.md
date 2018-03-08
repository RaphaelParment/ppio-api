# PPio RESTful API
RESTful API written in Golang with Gorilla MUX package. 

## Endpoints
All the endpoints are prefixed by /ppio. 

| URI            | Method | Description                         |
|:---------------|:-------|:------------------------------------|
| `/players`     | `GET`  | Get the full list of players.       |
| `/players/:id` | `GET`  | Get the player indexed by `{id}`    |
| `/players/:id` | `PUT`  | Update the player indexed by `{id}` |
| `/players`     | `POST` | Create a new player and persist it  |
