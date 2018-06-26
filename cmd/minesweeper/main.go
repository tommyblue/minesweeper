package main

import (
	"math/rand"
	"time"

	"github.com/tommyblue/minesweeper"
	"github.com/tommyblue/minesweeper/game"
	"github.com/tommyblue/minesweeper/ui"
)

func main() {
	rand.Seed(time.Now().Unix())

	board := &minesweeper.Board{
		Cols:  20,
		Rows:  20,
		Mines: 40,
	}

	ui := ui.Initialize()
	game := game.Setup(board, ui)

	game.Start()
	defer game.Exit()
}
