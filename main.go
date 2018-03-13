package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zacscodingclub/tic-tac-go/game"
)

var (
	numPlayers int
)

func main() {
	fmt.Println("Welcome to Tic Tac Go!")
	n := 10
	for n > 2 {
		n = askNumberOfPlayers()
	}

	fmt.Printf("now starting with %v players\n", n)
	g := game.NewGame(n)
	g.Board.Display()
	// g.Play()
}

func askNumberOfPlayers() int {
	fmt.Println("How many people will be playing? 1 or 2?")
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
