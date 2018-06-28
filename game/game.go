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
		switch g.Board.Tiles[x][y] {
		case minesweeper.Mine:
			// boom
			break
		case minesweeper.Empty:
			g.expandEmptyClick(x, y)
			break
		}
		g.updateMaskedBoard()
	}
}

func (g *Game) rightClickOnTile(x, y int32) {
	if g.State.DiscoveredTiles[x][y] != true {
		if g.MaskedBoard.Tiles[x][y] != minesweeper.Flag {
			g.MaskedBoard.Tiles[x][y] = minesweeper.Flag
		} else {
			g.MaskedBoard.Tiles[x][y] = minesweeper.Unknown
		}
		g.updateMaskedBoard()
	}
}

func (g *Game) expandEmptyClick(x, y int32) {
	// look around and make recursive calls
	fmt.Printf("Expanding to %d, %d\n", x, y)

	for _, coord := range g.getCoordsToExpand(x, y) {
		newX := coord[0]
		newY := coord[1]
		if g.State.DiscoveredTiles[newX][newY] == false {
			g.State.DiscoveredTiles[newX][newY] = true

			if g.Board.Tiles[newX][newY] == minesweeper.Empty {
				g.expandEmptyClick(newX, newY)
			}
		}
	}
}

func (g *Game) getCoordsToExpand(x, y int32) [][2]int32 {
	coords := [][2]int32{}
	if x-1 >= 0 {
		coords = append(coords, [2]int32{x - 1, y})
	}
	if y-1 >= 0 {
		coords = append(coords, [2]int32{x, y - 1})
	}
	if x-1 >= 0 && y-1 >= 0 {
		coords = append(coords, [2]int32{x - 1, y - 1})
	}
	if x+1 < g.Board.Cols {
		coords = append(coords, [2]int32{x + 1, y})
	}
	if y+1 < g.Board.Rows {
		coords = append(coords, [2]int32{x, y + 1})
	}
	if x+1 < g.Board.Cols && y+1 < g.Board.Rows {
		coords = append(coords, [2]int32{x + 1, y + 1})
	}
	if x-1 >= 0 && y+1 < g.Board.Rows {
		coords = append(coords, [2]int32{x - 1, y + 1})
	}
	if x+1 < g.Board.Cols && y-1 >= 0 {
		coords = append(coords, [2]int32{x + 1, y - 1})
	}
	return coords
}
