package main

import (
	"testing"

	"github.com/zacscodingclub/tic-tac-go/board"
)

func TestCreateNewBoard(t *testing.T) {
	b := board.NewBoard()

	if len(b.Cells) != 9 {
		t.Errorf("Board returned from NewBoard is the wrong size.")
	}

	for _, cell := range b.Cells {
		if cell != " " {
			t.Errorf("Expected cell: %s to equal: ' ' ", cell)
		}
	}
}

func TestBoardDisplay(t *testing.T) {
	t.Errorf("board.Display test not implemented")
}

func TestCreateNewGame(t *testing.T) {
	t.Errorf("game.NewGame(int) test not implemented")
}
