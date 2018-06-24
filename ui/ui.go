package ui

import (
	"github.com/tommyblue/minesweeper/sdl"
)

type UI struct {
	isRunning bool
}

func Initialize() *UI {
	sdlConf := &sdl.SdlConf{
		Title: "Minesweeper",
		Fonts: map[string]sdl.FontConfig{
			"mono": {
				Path: "../assets/fonts/mono.ttf",
				Size: 14,
			},
		},
	}
	sdl.InitSdl(sdlConf)
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
