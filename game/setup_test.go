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

	g := &Game{
		Board: board,
		UI:    nil,
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

	boards := []minesweeper.Board{
		{Cols: 4, Rows: 4, Mines: 5},
		{Cols: 10, Rows: 10, Mines: 10},
		{Cols: 10, Rows: 10, Mines: 100},
		{Cols: 20, Rows: 15, Mines: 80},
		{Cols: 25, Rows: 30, Mines: 250},
		{Cols: 40, Rows: 40, Mines: 600},
	}

	for _, b := range boards {
		board := &minesweeper.Board{
			Cols:  b.Cols,
			Rows:  b.Cols,
			Mines: b.Mines,
		}

		ui := ui.Initialize(24)
		g := &Game{
			Board: board,
			UI:    ui,
		}
		g.setInitialState()
		g.initBoard()
		g.initBoardTiles()

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
		var countedMines int32 = 0
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
				printTiles(g.Board.Tiles)
			}
		})
	}
}

func TestGame_setInitialState(t *testing.T) {
	g := &Game{}
	g.setInitialState()

	if g.State.CurrentState != minesweeper.InitialScreen || g.State.SelectedLevel != minesweeper.EasyLevel {
		t.Errorf("Game wasn't correctly initialized")
	}
}

func TestGame_initLevel(t *testing.T) {
	ui := ui.Initialize(24)
	g := &Game{UI: ui}
	g.setInitialState()
	g.initLevel()

	if g.State.DiscoveredTiles == nil {
		t.Errorf("DiscoveredTiles wasn't initialized")
	}
	type levelTest struct {
		level minesweeper.Level
		cols  int32
		rows  int32
		mines int32
	}
	tests := []levelTest{
		{minesweeper.EasyLevel, 20, 16, 30},
		{minesweeper.MediumLevel, 26, 20, 60},
		{minesweeper.HardLevel, 32, 24, 150},
	}
	for _, test := range tests {
		g.State.SelectedLevel = test.level
		g.initLevel()
		if g.Board.Cols != test.cols || g.Board.Rows != test.rows || g.Board.Mines != test.mines {
			t.Errorf("Wrong level numbers for %v", test.level)
		}
	}
}

func TestGame_updateMaskedBoard(t *testing.T) {
	ui := ui.Initialize(24)
	g := &Game{UI: ui}
	g.setInitialState()
	g.initBoard()
	g.initBoardTiles()
	g.initMaskedBoard()
	g.initDiscoveredTiles()

	type testStruct struct {
		discovered int
		won        bool
	}
	tt := []testStruct{
		{5, false},
		{20*16 - 30, true},
	}
	for _, test := range tt {
		_discoverTiles(g, test.discovered)
		if _countDiscoveredOnMask(g) == test.discovered {
			t.Errorf("Wrong number of discovered before update")
		}
		g.updateMaskedBoard()
		if _countDiscoveredOnMask(g) != test.discovered {
			t.Errorf("Wrong number of discovered after update")
		}
		if (g.State.CurrentState == minesweeper.Win) != test.won {
			t.Errorf("Wrong win state: %v", g.State.CurrentState)
		}
	}
}

func _discoverTiles(g *Game, tiles2discover int) {
	still2discover := tiles2discover
	var x, y int32
	for x = 0; x < g.Board.Cols; x++ {
		for y = 0; y < g.Board.Rows; y++ {
			if still2discover > 0 {
				g.State.DiscoveredTiles[x][y] = true
				still2discover--
			}
		}
	}
}

func _countDiscoveredOnMask(g *Game) int {
	discovered := 0
	var x, y int32
	for x = 0; x < g.MaskedBoard.Cols; x++ {
		for y = 0; y < g.MaskedBoard.Rows; y++ {
			a := g.MaskedBoard.Tiles[x][y]
			b := g.Board.Tiles[x][y]
			if a == b {
				discovered++
			}
		}
	}
	return discovered
}
