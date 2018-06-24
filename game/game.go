package game

import (
	"fmt"

	"github.com/tommyblue/minesweeper"
)

type Game struct {
	Board *minesweeper.Board
	UI    minesweeper.UI
}

func Setup(board *minesweeper.Board, ui minesweeper.UI) minesweeper.Game {
	game := &Game{
		Board: board,
		UI:    ui,
	}

	return game
}

func (g *Game) Start() {
	fmt.Println("Starting game...")
	ui := g.UI
	// loop
	ui.StartRunning()
	for ui.ShouldRun() {
		ui.ManageInput()
		ui.Draw()
	}
}

func (g *Game) Exit() {
	fmt.Println("Closing game...")
}

func (g *Game) Quit() {
	g.UI.StopRunning()
}
