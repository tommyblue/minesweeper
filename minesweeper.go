package minesweeper

import "os"

type Tile int

const (
	Empty     Tile = 0
	N1        Tile = 1
	N2        Tile = 2
	N3        Tile = 3
	N4        Tile = 4
	N5        Tile = 5
	N6        Tile = 6
	N7        Tile = 7
	N8        Tile = 8
	Mine      Tile = 9
	Explosion Tile = 10
	Flag      Tile = 11
	Unknown   Tile = 12
)

type Board struct {
	Cols  int32
	Rows  int32
	Mines int
	Tiles [][]Tile
}

type GameState struct {
	DiscoveredTiles [][]bool
}

type Game interface {
	Start()
	Exit()
}

type GameEventCallbacks interface {
	OnLeftClick(x, y int32)
	OnRightClick(x, y int32)
}

type UI interface {
	Draw(*Board)
	ManageInput()
	UpdateState(GameEventCallbacks)
	ShouldRun() bool
	StartRunning()
}

func IsDebug() bool {
	return os.Getenv("DEBUG") == "1"
}
