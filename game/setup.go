package game

import (
	"math/rand"

	"github.com/tommyblue/minesweeper"
)

func (g *Game) initBoardTiles() {
	g.Board.Tiles = make([][]minesweeper.Tile, g.Board.Cols)
	var i int32
	for i = 0; i < g.Board.Cols; i++ {
		g.Board.Tiles[i] = make([]minesweeper.Tile, g.Board.Rows)
	}
}

// put mines randomly
func (g *Game) setMines() {
	g.initBoardTiles()
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

func (g *Game) initBoard() {
	var cols, rows, mines int32
	switch g.State.SelectedLevel {
	case minesweeper.EasyLevel:
		cols = 20
		rows = 16
		mines = 30
		break
	case minesweeper.MediumLevel:
		cols = 26
		rows = 20
		mines = 60
		break
	case minesweeper.HardLevel:
		cols = 32
		rows = 24
		mines = 150
		break
	}
	g.Board = &minesweeper.Board{
		Cols:  cols,
		Rows:  rows,
		Mines: mines,
	}
	g.UI.SetTileSizes(cols, rows)
}

func (g *Game) initDiscoveredTiles() {
	g.State.DiscoveredTiles = make([][]bool, g.Board.Cols)

	var x, y int32
	for x = 0; x < g.Board.Cols; x++ {
		g.State.DiscoveredTiles[x] = make([]bool, g.Board.Rows)
		for y = 0; y < g.Board.Rows; y++ {
			g.State.DiscoveredTiles[x][y] = false
		}
	}
}

func (g *Game) initLevel() {
	g.initBoard()
	g.initDiscoveredTiles()
	g.initMaskedBoard()
	g.updateMaskedBoard()

	g.setMines()
	if minesweeper.IsDebug() {
		printTiles(g.Board.Tiles)
	}
}

/*
Updates the MaskedBoard "discovering" tiles, i.e. copying
the tile value from Board.Tiles if that tile has value
true in State.DiscoveredTiles.
After this updated, check the winning condition
*/
func (g *Game) updateMaskedBoard() {
	var x, y int32
	var unmaskedTiles int32
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
		g.setState(minesweeper.Win)
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
