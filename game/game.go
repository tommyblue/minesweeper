package game

import (
	"fmt"

	"github.com/tommyblue/minesweeper"
)

type Game struct {
	Board       *minesweeper.Board
	State       *minesweeper.GameState
	MaskedBoard *minesweeper.Board
	UI          minesweeper.UI
}

func Setup(board *minesweeper.Board, ui minesweeper.UI) minesweeper.Game {

	game := &Game{
		Board: board,
		UI:    ui,
	}

	game.setInitialState()

	game.setMines()
	if minesweeper.IsDebug() {
		printTiles(game.Board.Tiles)
	}
	return game
}

func (g *Game) Start() {
	fmt.Println("Starting game...")
	// loop
	g.UI.StartRunning()
	for g.UI.ShouldRun() {
		g.UI.ManageInput()
		g.UI.UpdateState(g.clickOnTile)
		g.UI.Draw(g.MaskedBoard)
	}
}

func (g *Game) Exit() {
	fmt.Println("Closing game...")
}

func (g *Game) clickOnTile(x, y int32) {
	g.State.DiscoveredTiles[x][y] = true
	g.updateMaskedBoard()
}
