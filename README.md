# PPio RESTful API

RESTful API written in Golang with Gorilla MUX package.

## Endpoints

| URI            | Method | Description                   |
|:---------------|:-------|:------------------------------|
| `/matches`     | `GET`  | Get the full list of matches. |
| `/matches/:id` | `GET`  | Get a specific match.         |
| `/matches`     | `POST` | Create a new match.           |

# Infrastructure

Build network:

```shell
docker network create ppio
```

```shell
cd infrastructure
docker-compose up -d
```

## Using lab

```shell
docker-compose up -d lab
docker exec -it lab sh
```

Connect to db

```shell
# in lab container
psql -h db -p 5432 -d ppio -U ppio
```