package board

import (
	"testing"
)

func TestBoardFullWithEmptyBoard(t *testing.T) {
	b := NewBoard()

	full := b.IsFull()
	if full != false {
		t.Errorf("Expected empty board to not be full: %t to equal: false", full)
	}
}

func TestBoardFullWithSingleItem(t *testing.T) {
	b := NewBoard()
	b.Cells[0] = "X"
	b.Cells[3] = "O"
	full := b.IsFull()
	if full != false {
		t.Errorf("Expected sparse board to not be full: %t to equal: false", full)
	}
}

func TestBoardFullWithFullBoard(t *testing.T) {
	b := NewBoard()
	for i := 0; i < 9; i++ {
		b.Cells[i] = "X"
	}
	full := b.IsFull()
	if full != true {
		t.Errorf("Expected full board to be full: %t to equal: true", full)
	}
}
