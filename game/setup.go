package game

import (
	"fmt"
	"math/rand"

	"github.com/tommyblue/minesweeper"
)

func (g *Game) initBoard() {
	g.Board.Tiles = make([][]minesweeper.Tile, g.Board.Cols)
	var i int32
	for i = 0; i < g.Board.Cols; i++ {
		g.Board.Tiles[i] = make([]minesweeper.Tile, g.Board.Rows)
	}
}

// put mines randomly
func (g *Game) setMines() {
	g.initBoard()
	minesToBePlaced := g.Board.Mines
	for minesToBePlaced > 0 {
		col := rand.Int31n(g.Board.Cols)
		row := rand.Int31n(g.Board.Rows)
		if g.Board.Tiles[col][row] != minesweeper.Mine {
			g.Board.Tiles[col][row] = minesweeper.Mine
			g.updateNumbersAroundMine(col, row)
			minesToBePlaced--
		}
	}
}

// update number of mines near tiles
func (g *Game) updateNumbersAroundMine(col, row int32) {
	for x := col - 1; x <= col+1; x++ {
		if x >= 0 && x < g.Board.Cols {
			for y := row - 1; y <= row+1; y++ {
				if y >= 0 && y < g.Board.Rows {
					if g.Board.Tiles[x][y] < minesweeper.Mine {
						g.Board.Tiles[x][y]++
					}
				}
			}
		}
	}
}

func (g *Game) setInitialState() {
	g.State = &minesweeper.GameState{
		CurrentState:  minesweeper.InitialScreen,
		SelectedLevel: minesweeper.EasyLevel,
	}
}

func (g *Game) initLevel() {
	var cols, rows, mines int32
	switch g.State.SelectedLevel {
	case minesweeper.EasyLevel:
		cols = 20
		rows = 15
		mines = 30
		break
	case minesweeper.MediumLevel:
		cols = 25
		rows = 20
		mines = 100
		break
	case minesweeper.HardLevel:
		cols = 33
		rows = 25
		mines = 200
		break
	}
	g.Board = &minesweeper.Board{
		Cols:  cols,
		Rows:  rows,
		Mines: mines,
	}

	g.State.DiscoveredTiles = make([][]bool, g.Board.Cols)

	var x, y int32
	for x = 0; x < g.Board.Cols; x++ {
		g.State.DiscoveredTiles[x] = make([]bool, g.Board.Rows)
		for y = 0; y < g.Board.Rows; y++ {
			g.State.DiscoveredTiles[x][y] = false
		}
	}
	g.initMaskedBoard()
	g.updateMaskedBoard()

	g.setMines()
	if minesweeper.IsDebug() {
		printTiles(g.Board.Tiles)
	}
}

func (g *Game) updateMaskedBoard() {
	var x, y int32
	var unmaskedTiles int32 = 0
	for x = 0; x < g.MaskedBoard.Cols; x++ {
		for y = 0; y < g.MaskedBoard.Rows; y++ {
			if g.State.DiscoveredTiles[x][y] == true {
				g.MaskedBoard.Tiles[x][y] = g.Board.Tiles[x][y]
			} else {
				unmaskedTiles++
			}
		}
	}
	if unmaskedTiles == g.Board.Mines {
		// Win!
		fmt.Println("You win!")
	}

}

// To be called on initial setup to init the masked board
func (g *Game) initMaskedBoard() {
	maskedBoard := minesweeper.Board{
		Cols:  g.Board.Cols,
		Rows:  g.Board.Rows,
		Mines: g.Board.Mines,
	}
	g.MaskedBoard = &maskedBoard

	var x, y int32
	tiles := make([][]minesweeper.Tile, g.Board.Cols)
	for x = 0; x < g.MaskedBoard.Cols; x++ {
		tiles[x] = make([]minesweeper.Tile, g.MaskedBoard.Rows)
		for y = 0; y < g.MaskedBoard.Rows; y++ {
			tiles[x][y] = minesweeper.Unknown
		}
	}
	g.MaskedBoard.Tiles = tiles
}
