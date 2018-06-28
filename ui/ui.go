package ui

import (
	"github.com/tommyblue/minesweeper"
	"github.com/tommyblue/minesweeper/graphy"
)

// Initialize the ui
func Initialize(tileSize int32) *UI {
	graphyConf := &graphy.GraphyConf{
		TileSize: tileSize,
		Window: &graphy.Window{
			Width:  800,
			Height: 600,
		},
		Debug: minesweeper.IsDebug(),
		Title: "Minesweeper",
		Fonts: map[string]graphy.FontConfig{
			"mono": {
				Path: "../assets/fonts/mono.ttf",
				Size: 14,
			},
		},
		BackgroundColor: &[4]uint8{255, 255, 255, 255},
		BackgroundImage: getAbsolutePath("../assets/images/bg.jpg"),
		ImagesToCache:   getImagesToCache(),
	}
	graphy.InitGraphy(graphyConf)
	ui := &UI{tileSize: tileSize}
	mapInputToFn(ui)

	return ui
}

// ShouldRun checks whether the game loop should be running
func (ui *UI) ShouldRun() bool {
	return ui.isRunning
}

// StartRunning starts the game loop
func (ui *UI) StartRunning() {
	ui.isRunning = true
}
