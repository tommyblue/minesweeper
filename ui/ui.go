package ui

import (
	"github.com/tommyblue/minesweeper/sdl"
)

type UI struct {
	isRunning bool
}

func Initialize() *UI {
	sdl.InitSdl("Minesweeper", "../assets/fonts/mono.ttf")
	ui := &UI{}
	initInput(ui)
	return ui
}

func (ui *UI) ShouldRun() bool {
	return ui.isRunning
}

func (ui *UI) StartRunning() {
	ui.isRunning = true
}

func (ui *UI) StopRunning() {
	ui.isRunning = false
}
