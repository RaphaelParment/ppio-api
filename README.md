# PPio RESTful API

RESTful API written in Golang with Gorilla MUX package.

## Endpoints

| URI            | Method | Description                   |
|:---------------|:-------|:------------------------------|
| `/matches`     | `GET`  | Get the full list of matches. |
| `/matches/:id` | `GET`  | Get a specific match.         |
| `/matches`     | `POST` | Create a new match.           |

## Infrastructure

```shell
# Start infrastructure containers
$ make infra-up

# Stop infrastructure containers
$ make infra-down
```

## Run app from lab

```shell
# Start lab
$ make lab-up

# Stop lan
$ make lab-down
```

With running infrastructure and lab containers

```shell
# Login lab
$ docker exec -it ppio-lab sh
$ go run main.go
```

## Connect to db from lab

```shell
$ docker exec -it ppio-lab sh
$ PGPASSWORD=dummy psql -h db -p 5432 -d ppio -U ppio
```

## Query server

### Get one match

```shell
# Getting match with id 1
$ curl http://localhost:9001/matches/1
```

### Get all matches

```shell
$ curl http://localhost:9001/matches
```

### Persist match
```shell
$ curl -d '{"player_one_id":1,"player_two_id":5,"result":{"winner_id":5,"loser_retired":false},"score":[{"player_one_score":11,"player_two_score":3},{"player_one_score":0,"player_two_score":11},{"player_one_score":8,"player_two_score":11}],"datetime":"2023-03-05 22:50:00"}' http://localhost:9001/matches
```