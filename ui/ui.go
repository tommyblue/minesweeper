package ui

import (
	"os"

	"github.com/tommyblue/minesweeper/sdl"
)

type UI struct {
	isRunning bool
}

func Initialize() *UI {
	sdlConf := &sdl.SdlConf{
		Debug: (os.Getenv("DEBUG") == "1"),
		Title: "Minesweeper",
		Fonts: map[string]sdl.FontConfig{
			"mono": {
				Path: "../assets/fonts/mono.ttf",
				Size: 14,
			},
		},
		BackgroundColor: &[4]uint8{255, 255, 255, 255},
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
