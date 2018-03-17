package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func Run() {
	g := newGame()
	g.Board.Display()
	g.Play()
}

func newGame() *Game {
	fmt.Println("Welcome to Tic Tac Go!")
	n := 10
	for n > 2 {
		n = askNumberOfPlayers()
	}
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

func (g *Game) Play() {
	for !g.isOver() {
		g.turn()
	}
	if g.isWon() {
		fmt.Printf("Congratulations! Player %s won!\n", g.Board.Cells[g.checkWin()])
	} else {
		fmt.Println("Tie Game!")
	}
}

func (g *Game) isOver() bool {
	return g.isWon() || g.Board.IsFull()
}

func (g *Game) isWon() bool {
	return g.checkWin() >= 0
}

func (g *Game) checkWin() int {
	for _, combo := range winningCombinations {
		if g.Board.Cells[combo[0]] != " " &&
			g.Board.Cells[combo[0]] == g.Board.Cells[combo[1]] &&
			g.Board.Cells[combo[1]] == g.Board.Cells[combo[2]] {
			return combo[0]
		}
	}
	return -1
}

func (g *Game) currentPlayer() string {
	if g.turnCount()%2 == 0 {
		return "X"
	}

	return "O"
}

func (g *Game) turn() {
	p := g.currentPlayer()
	fmt.Printf("Where would you like to move Player %s? (1 - 9)\n", p)
	move := intToIndex(inputToInt())

	if g.Board.IsValid(move) {
		g.Board.Move(move, p)
		g.Board.Display()
	} else {
		fmt.Println("That move is not valid.  Please try again.")
	}
}

func (g *Game) turnCount() int {
	count := 0
	for _, cell := range g.Board.Cells {
		if cell != " " {
			count += 1
		}
	}
	return count
}

func askNumberOfPlayers() int {
	fmt.Println("How many people will be playing? 1 or 2?")
	num := inputToInt()

	return num
}

func inputToInt() int {
	reader := bufio.NewReader(os.Stdin)
	const delimiter = '\n'
	input, err := reader.ReadString(delimiter)
	input = strings.Replace(input, "\n", "", -1)
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return num
}

func intToIndex(n int) int {
	return n - 1
}
