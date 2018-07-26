package game

import (
	"testing"

	"github.com/tommyblue/minesweeper"
)

func TestGame_OnLeftClick(t *testing.T) {
	leftClick := 0
	rightClick := 0

	// prepare mock
	m := eventCallbacks{
		leftClick:  func(*minesweeper.Position, *minesweeper.Position) { leftClick++ },
		rightClick: func(*minesweeper.Position, *minesweeper.Position) { rightClick++ },
	}

	p := &minesweeper.Position{}

	m.OnLeftClick(p, p)
	if leftClick != 1 {
		t.Errorf("Left click not working")
	}

	m.OnRightClick(p, p)
	if rightClick != 1 {
		t.Errorf("Right click not working")
	}
}
