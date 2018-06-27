package ui

import (
	"github.com/tommyblue/minesweeper"
	"github.com/tommyblue/minesweeper/graphy"
)

type UI struct {
	isRunning bool
}

func Initialize() *UI {
	graphyConf := &graphy.GraphyConf{
		Debug: minesweeper.IsDebug(),
		Title: "Minesweeper",
		Fonts: map[string]graphy.FontConfig{
			"mono": {
				Path: "../assets/fonts/mono.ttf",
				Size: 14,
			},
		},
		BackgroundColor: &[4]uint8{255, 255, 255, 255},
		ImagesToCache:   getImagesToCache(),
	}
	graphy.InitGraphy(graphyConf)
	ui := &UI{}
	initInput(ui)

	/* TODO:
	- preload images
	*/
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
