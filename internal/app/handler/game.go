package handler

import (
	"encoding/json"
	"net/http"
	"tictactoe-api/internal/app/service"
	"tictactoe-api/internal/domain"
)

type Game interface {
	Get(w http.ResponseWriter, r *http.Request)
	Move(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type Marshal func(v any) ([]byte, error)

type game struct {
	gameService service.Game
}

func NewGame(gameService service.Game) Game {
	return &game{gameService: gameService}
}

// Get current game board status
func (g *game) Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(g.gameService.Get())

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Move make a move on the game board
func (g *game) Move(w http.ResponseWriter, r *http.Request) {
	var m domain.Move
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)

	var (
		err    error
		result *domain.Game
	)
	defer func() {
		if err != nil {
			encoder.Encode(map[string]string{"error": err.Error()})
			return
		}
		encoder.Encode(result)
	}()

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err = g.gameService.Move(&m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete reset the game board
func (g *game) Delete(w http.ResponseWriter, r *http.Request) {
	g.gameService.Delete()
	w.WriteHeader(http.StatusNoContent)
}
