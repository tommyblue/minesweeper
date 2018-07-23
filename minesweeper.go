package minesweeper

import (
	"os"
)

type Tile int

// Not using a iota to be sure that each number corresponds to the right status
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
	Mines int32
	Tiles [][]Tile
}

type State int

const (
	InitialScreen State = iota
	SelectLevel
	InAGame
	Lost
	Win
)

type Level int

const (
	EasyLevel Level = iota
	MediumLevel
	HardLevel
)

type GameState struct {
	DiscoveredTiles [][]bool
	CurrentState    State
	SelectedLevel   Level
}

type GameEventCallbacks interface {
	OnLeftClick(*Position, *Position)
	OnRightClick(*Position, *Position)
}

type UI interface {
	Draw(State, *Board)
	ManageInput()
	UpdateState(GameEventCallbacks)
	ShouldRun() bool
	StartRunning()
	StopRunning()
	GetButton(string) *Button
	SetTileSizes(int32, int32)
	GetCols() int32
	GetRows() int32
}

type Button struct {
	W   int32
	H   int32
	Src string
	X   int32 // X position where the button has been drawn
	Y   int32 // Y position where the button has been drawn
}

// Position defines a 2D (X,Y) position
type Position struct {
	X int32
	Y int32
}

// Rect defines a rectangle
type Rect struct {
	X0 int32
	Y0 int32
	X1 int32
	Y1 int32
}

func (p *Position) ClickedOn(clickArea Rect) bool {
	return p.X > clickArea.X0 && p.X <= clickArea.X1 &&
		p.Y > clickArea.Y0 && p.Y <= clickArea.Y1
}

func IsDebug() bool {
	return os.Getenv("DEBUG") == "1"
}
