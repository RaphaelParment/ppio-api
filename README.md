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
$ docker exec -it lab sh
$ go run main.go
```

## Connect to db from lab

```shell
$ docker exec -it lab sh
$ PGPASSWORD=dummy psql -h db -p 5432 -d ppio -U ppio
```