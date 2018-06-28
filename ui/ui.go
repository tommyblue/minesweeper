package ui

import (
	"github.com/tommyblue/minesweeper"
	"github.com/tommyblue/minesweeper/graphy"
)

type tile struct {
	x int32
	y int32
}

type event struct {
	evType string
	tile   *tile
}

type UI struct {
	isRunning bool
	tileSize  int32
	event     *event
}

func Initialize(tileSize int32) *UI {
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
	ui := &UI{tileSize: tileSize}
	mapInputToFn(ui)

	return ui
}

func (ui *UI) ShouldRun() bool {
	return ui.isRunning
}

func (ui *UI) StartRunning() {
	ui.isRunning = true
}
