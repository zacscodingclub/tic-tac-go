package board

import "fmt"

type Board struct {
	Cells [9]string
}

func NewBoard() *Board {
	var c [9]string
	for i, _ := range c {
		c[i] = " "
	}
	return &Board{Cells: c}
}

func (b *Board) Display() {
	gameLine(b.Cells[0], b.Cells[1], b.Cells[2])
	divLine()
	gameLine(b.Cells[3], b.Cells[4], b.Cells[5])
	divLine()
	gameLine(b.Cells[6], b.Cells[7], b.Cells[8])
}

func (b *Board) IsFull() bool {
	for _, cell := range b.Cells {
		if cell != "X" && cell != "O" {
			return false
		}
	}
	return true
}

func (b *Board) IsTaken(n int) bool {
	return b.Cells[n] == "X" || b.Cells[n] == "O"
}

func (b *Board) IsValid(n int) bool {
	l := n >= 0 && n <= 8
	return l && !b.IsTaken(n)
}

func (b *Board) Move(n int, p string) {
	b.Cells[n] = p
}

func gameLine(pos1, pos2, pos3 string) {
	fmt.Printf(" %v | %v | %v \n", pos1, pos2, pos3)
}

func divLine() {
	fmt.Println("-----------")
}
