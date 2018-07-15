package game

import (
	"fmt"

	"github.com/tommyblue/minesweeper"
)

// Game represents the main game object
type Game struct {
	Board          *minesweeper.Board
	State          *minesweeper.GameState
	MaskedBoard    *minesweeper.Board
	UI             minesweeper.UI
	EventCallbacks *eventCallbacks
}

// Setup the game, setting mouse click callbacks and the initial state
func Setup(ui minesweeper.UI) minesweeper.Game {

	game := &Game{
		UI: ui,
	}
	game.EventCallbacks = &eventCallbacks{
		leftClick:  game.leftClickOnTile,
		rightClick: game.rightClickOnTile,
	}
	game.setInitialState()

	return game
}

// Start the game running the main game loop
func (g *Game) Start() {
	if minesweeper.IsDebug() {
		fmt.Println("Starting game...")
	}
	// loop
	g.UI.StartRunning()
	for g.UI.ShouldRun() {
		g.UI.ManageInput()
		g.updateState()
		g.UI.Draw(g.State.CurrentState, g.MaskedBoard)
	}
}

// Exit is called at the end of the game. Make cleanups here
func (g *Game) Exit() {
	if minesweeper.IsDebug() {
		fmt.Println("Closing game...")
	}
}

func (g *Game) updateState() {
	g.UI.UpdateState(g.EventCallbacks)
}

func (g *Game) leftClickOnTile(tileClick *minesweeper.Position, mouseClick *minesweeper.Position) {
	switch g.State.CurrentState {
	case minesweeper.InitialScreen, minesweeper.Lost:
		new := g.UI.GetButton("button_new")
		quit := g.UI.GetButton("button_quit")
		if mouseClick.ClickedOn(minesweeper.Rect{
			X0: new.X,
			X1: new.X + new.W,
			Y0: new.Y,
			Y1: new.Y + new.H,
		}) {
			g.setState(minesweeper.SelectLevel)
		} else if mouseClick.ClickedOn(minesweeper.Rect{
			X0: quit.X,
			X1: quit.X + quit.W,
			Y0: quit.Y + quit.H + 4,
			Y1: quit.Y + quit.H + 4 + quit.H,
		}) {
			g.UI.StopRunning()
		}
		break
	case minesweeper.SelectLevel:
		easy := g.UI.GetButton("button_easy")
		medium := g.UI.GetButton("button_medium")
		hard := g.UI.GetButton("button_hard")
		if mouseClick.ClickedOn(minesweeper.Rect{
			X0: easy.X,
			X1: easy.X + easy.W,
			Y0: easy.Y,
			Y1: easy.Y + easy.H,
		}) {
			g.selectLevel(minesweeper.EasyLevel)
		} else if mouseClick.ClickedOn(minesweeper.Rect{
			X0: medium.X,
			X1: medium.X + medium.W,
			Y0: medium.Y + easy.H + 4,
			Y1: medium.Y + easy.H + 4 + medium.H,
		}) {
			g.selectLevel(minesweeper.MediumLevel)
		} else if mouseClick.ClickedOn(minesweeper.Rect{
			X0: hard.X,
			X1: hard.X + hard.W,
			Y0: hard.Y + easy.H + 4 + medium.H + 4,
			Y1: hard.Y + easy.H + 4 + medium.H + 4 + hard.H,
		}) {
			g.selectLevel(minesweeper.HardLevel)
		} else {
			// Click outside buttons
			break
		}
		g.initLevel()
		g.setState(minesweeper.InAGame)
		break
	case minesweeper.InAGame:
		x := tileClick.X
		y := tileClick.Y
		canClick := g.State.DiscoveredTiles[x][y] != true &&
			g.MaskedBoard.Tiles[x][y] != minesweeper.Flag
		if canClick {
			g.State.DiscoveredTiles[x][y] = true
			switch g.Board.Tiles[x][y] {
			case minesweeper.Mine:
				g.mineExplodedAt(x, y)
				g.setState(minesweeper.Lost)
				break
			case minesweeper.Empty:
				g.expandEmptyClick(x, y)
				break
			}
			g.updateMaskedBoard()
		}
		break
	}
}

func (g *Game) mineExplodedAt(x, y int32) {
	g.Board.Tiles[x][y] = minesweeper.Explosion
	g.showAllMines()
}

func (g *Game) showAllMines() {
	var x, y int32
	for x = 0; x < g.Board.Cols; x++ {
		for y = 0; y < g.Board.Rows; y++ {
			if g.Board.Tiles[x][y] == minesweeper.Mine {
				g.State.DiscoveredTiles[x][y] = true
			}
		}
	}
}

func (g *Game) resetDiscoveredTiles() {
	var x, y int32
	for x = 0; x < g.Board.Cols; x++ {
		for y = 0; y < g.Board.Rows; y++ {
			g.State.DiscoveredTiles[x][y] = false
		}
	}
}

func (g *Game) rightClickOnTile(tileClick *minesweeper.Position, mouseClick *minesweeper.Position) {
	switch g.State.CurrentState {
	case minesweeper.InAGame:
		x := tileClick.X
		y := tileClick.Y
		if g.State.DiscoveredTiles[x][y] != true {
			if g.MaskedBoard.Tiles[x][y] != minesweeper.Flag {
				g.MaskedBoard.Tiles[x][y] = minesweeper.Flag
			} else {
				g.MaskedBoard.Tiles[x][y] = minesweeper.Unknown
			}
			g.updateMaskedBoard()
		}
		break
	}
}

func (g *Game) expandEmptyClick(x, y int32) {
	// look around and make recursive calls
	if minesweeper.IsDebug() {
		fmt.Printf("Expanding to %d, %d\n", x, y)
	}

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
