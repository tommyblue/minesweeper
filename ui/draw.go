package ui

import (
	"github.com/tommyblue/minesweeper"
	"github.com/tommyblue/minesweeper/sdl"
)

func (ui *UI) Draw(*minesweeper.Board) {
	sdl.Draw()
}
