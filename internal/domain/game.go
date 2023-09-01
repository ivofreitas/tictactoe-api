package domain

import "github.com/chrisfregly/tictactoe"

// Game current tictactoe game status
type Game struct {
	Turn   tictactoe.Player                                            `json:"turn"`
	Winner *tictactoe.Player                                           `json:"winner,omitempty"`
	Board  [tictactoe.BoardSize][tictactoe.BoardSize]*tictactoe.Player `json:"board"`
}

// Move request on the tictactoe game
type Move struct {
	Player tictactoe.Player `json:"player"`
	Row    int              `json:"row"`
	Column int              `json:"column"`
}
