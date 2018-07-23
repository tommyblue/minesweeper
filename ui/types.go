package ui

import (
	"github.com/tommyblue/matrigo"
	"github.com/tommyblue/minesweeper"
)

type eventType int

const (
	mouseLeftClick eventType = iota
	mouseRightClick
)

type event struct {
	evType     eventType
	tile       *minesweeper.Position
	mouseClick *minesweeper.Position
}

type UI struct {
	isRunning bool
	tileSize  int32
	event     *event
	window    *matrigo.Window
	cols      int32
	rows      int32
}

// This struct must implement matrigo.Input interface, managing all possible input events.
// It contains the methods to be mapped to actual ui methods.
type matrigoInputInterface struct {
	mouseLeftClickDownFn  func(x, y int32)
	mouseRightClickDownFn func(x, y int32)
	quitFn                func()
}

type tileImageName string
