package ui

import (
	"github.com/tommyblue/minesweeper"
	"github.com/tommyblue/minesweeper/graphy"
)

func (ui *UI) Draw(*minesweeper.Board) {
	graphy.Draw()
}
