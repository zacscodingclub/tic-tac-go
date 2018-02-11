package game

import (
	"github.com/zacscodingclub/tic-tac-go/board"
)

var winningCombinations = [8][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{6, 4, 2},
}

type Game struct {
	Board   board.Board
	Players []player
}

func NewGame() *Game {
	b := board.NewBoard()
	return &Game{Board: *b}
}
