package main

import (
	"log"
	"net/http"
	"os"
	"tictactoe-api/internal/app"
	"tictactoe-api/internal/app/handler"
	"tictactoe-api/internal/app/service"
)

// main is the entry point of the application
func main() {
	gameHandler := handler.NewGame(service.NewGame())
	router := app.NewRouter(gameHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8088"
	}
	log.Println("server is starting...")
	log.Fatal(http.ListenAndServe(port, router))
}
