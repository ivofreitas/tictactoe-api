# tictactoe-api

Simple tictactoe game server

## Requirements
*  Go 1.21 or higher. If you don't have Golang installed, you can download it from https://go.dev/doc/install)
*  Docker (https://docs.docker.com/engine/)

## Setup

1. Clone the repository: `git clone https://github.com/ivofreitas/tictactoe-api.git`
2. Install dependencies: `go mod tidy`
3. Run the server: `go run main.go`
4. Testing with coverage:  `go test -cover ./...`

## Usage

### Endpoints

The API has the following endpoints:

------------------

#### Game

GET /game Get the current game board status.
```
curl --location 'http://localhost:8088/game'
```
POST /game/move Place a move in the game board.
```
curl --location 'http://localhost:8088/game/move' \
--header 'Content-Type: application/json' \
--data '{
    "player": "X",
    "row": 2,
    "column": 2
}'
```
DELETE /game Reset the game board.
```
curl --location --request DELETE 'http://localhost:8088/game'
```

## Docker

Running it in a docker container:

1. Docker image: `docker build -t tictactoe-api .`

2. Run: `docker run -p 8088:8088 tictactoe-api`