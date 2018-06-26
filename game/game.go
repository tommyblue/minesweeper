package game

import (
	"fmt"

	"github.com/tommyblue/minesweeper"
)

type Game struct {
	Board *minesweeper.Board
	State *minesweeper.GameState
	UI    minesweeper.UI
}

func Setup(board *minesweeper.Board, ui minesweeper.UI) minesweeper.Game {
	// All tiles are still to be discovered
	initialState := &minesweeper.GameState{
		DiscoveredTiles: [][]bool{},
	}
	initialState.DiscoveredTiles = make([][]bool, board.Cols)
	var x, y int32
	for x = 0; x < board.Cols; x++ {
		initialState.DiscoveredTiles[x] = make([]bool, board.Rows)
		for y = 0; y < board.Rows; y++ {
			initialState.DiscoveredTiles[x][y] = false
		}
	}

	game := &Game{
		Board: board,
		UI:    ui,
		State: initialState,
	}

	game.setMines()
	printTiles(game.Board.Tiles)

	return game
}

func (g *Game) Start() {
	fmt.Println("Starting game...")
	// loop
	g.UI.StartRunning()
	for g.UI.ShouldRun() {
		g.UI.ManageInput()
		g.UI.Draw(g.Board)
	}
}

func (g *Game) Exit() {
	fmt.Println("Closing game...")
}

func (g *Game) Quit() {
	g.UI.StopRunning()
}

func printTiles(tiles [][]minesweeper.Tile) {
	for _, x := range tiles {
		fmt.Printf("%v\n", x)
	}
}
