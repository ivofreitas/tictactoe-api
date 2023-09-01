package service

import (
	"github.com/chrisfregly/tictactoe"
	"tictactoe-api/internal/domain"
)

// Game wrapper around tictactoe implementation
type Game interface {
	Get() *domain.Game
	Move(move *domain.Move) (*domain.Game, error)
	Delete()
}

type game struct {
	tictactoe.TicTacToe
}

func NewGame() Game {
	return &game{tictactoe.NewTicTacToe()}
}

// Get current game board status
func (g *game) Get() *domain.Game {
	entity := new(domain.Game)
	if winner := g.TicTacToe.GetWinner(); winner != nil {
		entity.Winner = winner
	}
	entity.Turn = g.TicTacToe.GetTurn()
	entity.Board = g.TicTacToe.GetBoard()

	return entity
}

// Move make a move on the game board
func (g *game) Move(move *domain.Move) (*domain.Game, error) {
	if err := g.TicTacToe.Move(move.Player, move.Row, move.Column); err != nil {
		return nil, err
	}

	return g.Get(), nil
}

// Delete reset the game board
func (g *game) Delete() {
	g.TicTacToe = tictactoe.NewTicTacToe()
}
