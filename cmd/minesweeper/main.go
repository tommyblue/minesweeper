package main

import (
	"github.com/tommyblue/minesweeper"
	"github.com/tommyblue/minesweeper/game"
	"github.com/tommyblue/minesweeper/ui"
)

func main() {
	board := &minesweeper.Board{
		Height: 20,
		Width:  20,
		Mines:  5,
	}

	ui := ui.Initialize()
	game := game.Setup(board, ui)

	game.Start()
	defer game.Exit()
}
