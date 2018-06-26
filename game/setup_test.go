package game

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tommyblue/minesweeper"
	"github.com/tommyblue/minesweeper/ui"
)

func TestGame_updateNumbersAroundMine(t *testing.T) {
	rand.Seed(time.Now().Unix())

	board := &minesweeper.Board{
		Cols:  4,
		Rows:  6,
		Mines: 5,
	}

	ui := ui.Initialize()

	g := &Game{
		Board: board,
		UI:    ui,
	}

	expected1 := [][]minesweeper.Tile{
		[]minesweeper.Tile{1, 9, 9, 2, 1, 0},
		[]minesweeper.Tile{1, 2, 3, 9, 2, 1},
		[]minesweeper.Tile{1, 1, 1, 2, 9, 1},
		[]minesweeper.Tile{9, 1, 0, 1, 1, 1},
	}
	tiles1 := [][]minesweeper.Tile{
		[]minesweeper.Tile{0, 9, 9, 0, 0, 0},
		[]minesweeper.Tile{0, 0, 0, 9, 0, 0},
		[]minesweeper.Tile{0, 0, 0, 0, 9, 0},
		[]minesweeper.Tile{9, 0, 0, 0, 0, 0},
	}

	expected2 := [][]minesweeper.Tile{
		[]minesweeper.Tile{9, 2, 1, 1, 1, 9},
		[]minesweeper.Tile{1, 2, 9, 1, 1, 1},
		[]minesweeper.Tile{1, 2, 1, 1, 1, 1},
		[]minesweeper.Tile{9, 1, 0, 0, 1, 9},
	}
	tiles2 := [][]minesweeper.Tile{
		[]minesweeper.Tile{9, 0, 0, 0, 0, 9},
		[]minesweeper.Tile{0, 0, 9, 0, 0, 0},
		[]minesweeper.Tile{0, 0, 0, 0, 0, 0},
		[]minesweeper.Tile{9, 0, 0, 0, 0, 9},
	}

	expected3 := [][]minesweeper.Tile{
		[]minesweeper.Tile{0, 1, 2, 2, 1, 0},
		[]minesweeper.Tile{1, 2, 9, 9, 2, 0},
		[]minesweeper.Tile{1, 9, 5, 9, 2, 0},
		[]minesweeper.Tile{1, 2, 9, 2, 1, 0},
	}
	tiles3 := [][]minesweeper.Tile{
		[]minesweeper.Tile{0, 0, 0, 0, 0, 0},
		[]minesweeper.Tile{0, 0, 9, 9, 0, 0},
		[]minesweeper.Tile{0, 9, 0, 9, 0, 0},
		[]minesweeper.Tile{0, 0, 9, 0, 0, 0},
	}

	tests := []struct {
		name     string
		expected [][]minesweeper.Tile
		tiles    [][]minesweeper.Tile
	}{
		{"Mines spread", expected1, tiles1},
		{"Mines on edges", expected2, tiles2},
		{"Mines very near", expected3, tiles3},
	}

	g.Board.Tiles = [][]minesweeper.Tile{
		[]minesweeper.Tile{
			0, 9, 9, 0, 0, 0,
		},
		[]minesweeper.Tile{
			0, 0, 0, 9, 0, 0,
		},
		[]minesweeper.Tile{
			0, 0, 0, 0, 9, 0,
		},
		[]minesweeper.Tile{
			9, 0, 0, 0, 0, 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g.Board.Tiles = tt.tiles
			var x, y int32
			for x = 0; x < g.Board.Cols; x++ {
				for y = 0; y < g.Board.Rows; y++ {
					if g.Board.Tiles[x][y] == minesweeper.Mine {
						g.updateNumbersAroundMine(x, y)
					}
				}
			}

			for x = 0; x < g.Board.Cols; x++ {
				for y = 0; y < g.Board.Rows; y++ {
					if g.Board.Tiles[x][y] != tt.expected[x][y] {
						t.Errorf("Board count is wrong: %v != %v", g.Board.Tiles[x][y], tt.expected[x][y])
					}
				}
			}
		})
	}
}

func TestGame_setMines(t *testing.T) {
	rand.Seed(time.Now().Unix())

	board := &minesweeper.Board{
		Cols:  4,
		Rows:  6,
		Mines: 5,
	}

	ui := ui.Initialize()

	g := &Game{
		Board: board,
		UI:    ui,
	}
	g.initBoard()

	var x, y int32
	t.Run("Board should be empty", func(t *testing.T) {
		for x = 0; x < g.Board.Cols; x++ {
			for y = 0; y < g.Board.Rows; y++ {
				if g.Board.Tiles[x][y] != minesweeper.Empty {
					t.Errorf("Tile should be empty at this point")
				}
			}
		}
	})

	g.setMines()
	countedMines := 0
	t.Run("Board should be populated", func(t *testing.T) {
		for x = 0; x < g.Board.Cols; x++ {
			for y = 0; y < g.Board.Rows; y++ {
				if g.Board.Tiles[x][y] == minesweeper.Mine {
					countedMines++
				}
			}
		}
		if countedMines != g.Board.Mines {
			t.Errorf("I counted %v mines instead of %v", countedMines, g.Board.Mines)
		}
	})
}
