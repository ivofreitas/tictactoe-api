package mock

import (
	"github.com/chrisfregly/tictactoe"
	"tictactoe-api/internal/domain"
)

var (
	playerX = tictactoe.Player("X")
	Move    = domain.Move{
		Player: playerX,
		Row:    0,
		Column: 0,
	}
)
