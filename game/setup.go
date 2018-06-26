package game

import (
	"math/rand"

	"github.com/tommyblue/minesweeper"
)

// put mines randomly
func (g *Game) setMines() {
	board := g.Board
	g.Board.Tiles = make([][]minesweeper.Tile, board.Cols)
	var i int32
	for i = 0; i < board.Cols; i++ {
		g.Board.Tiles[i] = make([]minesweeper.Tile, board.Rows)
	}

	for i = 0; i < board.Mines; i++ {
		col := rand.Int31n(board.Cols)
		row := rand.Int31n(board.Rows)
		g.Board.Tiles[col][row] = Mine
		g.updateNumbersAroundMine(col, row)
	}
}

func (g *Game) updateNumbersAroundMine(col, row int32) {
	for x := col - 1; x <= col+1; x++ {
		if x >= 0 && x < g.Board.Cols {
			for y := row - 1; y <= row+1; y++ {
				if y >= 0 && y < g.Board.Rows {
					if g.Board.Tiles[x][y] < Mine {
						g.Board.Tiles[x][y]++
					}
				}
			}
		}
	}
}
