package game

import (
	"testing"
)

func TestGameAskNumberOfPlayers(t *testing.T) {
	t.Errorf("Game askNumberOfPlayers test not implemented")
}

func TestGameInputToInt(t *testing.T) {
	t.Errorf("Game inputToInt test not implemented")
}
func TestGameIntToIndex(t *testing.T) {
	a := intToIndex(1)
	if a != 0 {
		t.Errorf("Expected integer 1 to be converted to index: %d to equal: 0", a)
	}
	b := intToIndex(9)
	if b != 8 {
		t.Errorf("Expected integer 9 to be converted to index: %d to equal: 8", b)
	}
	c := intToIndex(5)
	if c != 4 {
		t.Errorf("Expected integer 5 to be converted to index: %d to equal: 4", c)
	}
}
