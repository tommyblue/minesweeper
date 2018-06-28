package game

import (
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
	// All tiles are still to be discovered
	g.State = &minesweeper.GameState{
		DiscoveredTiles: [][]bool{},
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
}

func (g *Game) updateMaskedBoard() {
	g.MaskedBoard.Tiles = make([][]minesweeper.Tile, g.Board.Cols)

	var x, y int32
	for x = 0; x < g.MaskedBoard.Cols; x++ {
		g.MaskedBoard.Tiles[x] = make([]minesweeper.Tile, g.MaskedBoard.Rows)
		for y = 0; y < g.MaskedBoard.Rows; y++ {
			if g.State.DiscoveredTiles[x][y] == false {
				g.MaskedBoard.Tiles[x][y] = minesweeper.Unknown
			} else {
				g.MaskedBoard.Tiles[x][y] = g.Board.Tiles[x][y]
			}
		}
	}
}

// To be called on initial setup to init the masked board
func (g *Game) initMaskedBoard() {
	if g.MaskedBoard == nil {
		maskedBoard := minesweeper.Board{
			Cols:  g.Board.Cols,
			Rows:  g.Board.Rows,
			Mines: g.Board.Mines,
			Tiles: [][]minesweeper.Tile{},
		}
		g.MaskedBoard = &maskedBoard
	}
}
