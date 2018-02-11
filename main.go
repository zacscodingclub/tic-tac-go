package main

import (
	"flag"
	"fmt"

	"github.com/zacscodingclub/tic-tac-go/game"
)

var (
	numPlayers int
)

func init() {
	flag.IntVar(&numPlayers, "players", 1, "Number of human players")
	flag.IntVar(&numPlayers, "p", 1, "Number of human players")

	flag.Parse()
}

func main() {
	fmt.Println("Welcome to Tic Tac Go!")
	fmt.Println("Number of Players", numPlayers)
	g := game.NewGame()
	g.Board.Display()
	g.Play()
}
