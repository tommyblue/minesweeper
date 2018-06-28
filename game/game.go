package game

import (
	"fmt"

	"github.com/tommyblue/minesweeper"
)

type Game struct {
	Board          *minesweeper.Board
	State          *minesweeper.GameState
	MaskedBoard    *minesweeper.Board
	UI             minesweeper.UI
	EventCallbacks *eventCallbacks
}

func Setup(board *minesweeper.Board, ui minesweeper.UI) minesweeper.Game {

	game := &Game{
		Board: board,
		UI:    ui,
	}
	game.EventCallbacks = &eventCallbacks{
		leftClick:  game.leftClickOnTile,
		rightClick: game.rightClickOnTile,
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
		g.updateState()
		g.UI.Draw(g.MaskedBoard)
	}
}

func (g *Game) Exit() {
	fmt.Println("Closing game...")
}

func (g *Game) updateState() {
	g.UI.UpdateState(g.EventCallbacks)
}

func (g *Game) leftClickOnTile(x, y int32) {
	if g.State.DiscoveredTiles[x][y] != true && g.MaskedBoard.Tiles[x][y] != minesweeper.Flag {
		g.State.DiscoveredTiles[x][y] = true
		// TODO:
		// - if empty tile, expand discovered tiles
		// - if bomb, boom!
		g.updateMaskedBoard()
	}
}

func (g *Game) rightClickOnTile(x, y int32) {
	if g.State.DiscoveredTiles[x][y] != true {
		g.MaskedBoard.Tiles[x][y] = minesweeper.Flag
		g.updateMaskedBoard()
	}
}
