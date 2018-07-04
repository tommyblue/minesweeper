package game

import (
	"github.com/tommyblue/minesweeper"
)

type eventCallbacks struct {
	leftClick  func(*minesweeper.Position, *minesweeper.Position)
	rightClick func(*minesweeper.Position, *minesweeper.Position)
}

func (ev eventCallbacks) OnLeftClick(tileClick *minesweeper.Position, mouseClick *minesweeper.Position) {
	ev.leftClick(tileClick, mouseClick)
}

func (ev eventCallbacks) OnRightClick(tileClick *minesweeper.Position, mouseClick *minesweeper.Position) {
	ev.rightClick(tileClick, mouseClick)
}
