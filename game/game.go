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
	Players []string
}

func NewGame(n int) *Game {
	b := board.NewBoard()
	p := make([]string, 2)
	if n == 1 {
		p[0] = "human"
		p[1] = "computer"
	} else {
		p[0] = "human"
		p[1] = "human"
	}
	return &Game{Board: *b, Players: p}
}
