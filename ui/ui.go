package ui

import (
	"github.com/tommyblue/matrigo"
	"github.com/tommyblue/minesweeper"
)

// Initialize the ui
func Initialize(tileSize int32) *UI {
	w := &matrigo.Window{
		Width:  800,
		Height: 600,
	}
	matrigoConf := &matrigo.Conf{
		TileSize: tileSize,
		Window:   w,
		Debug:    minesweeper.IsDebug(),
		Title:    "Minesweeper",
		Fonts: map[string]matrigo.FontConfig{
			"mono": {
				Path: getAbsolutePath("../assets/fonts/mono.ttf"),
				Size: 14,
			},
		},
		BackgroundColor: &[4]uint8{255, 255, 255, 255},
		BackgroundImage: getAbsolutePath("../assets/images/bg.jpg"),
		ImagesToCache:   getImagesToCache(),
		SyncFPS:         true,
	}

	err := matrigo.Init(matrigoConf)
	if err != nil {
		panic(err)
	}

	ui := &UI{
		window:   w,
		tileSize: tileSize,
	}
	mapInputToFn(ui)

	return ui
}

func (ui *UI) SetTileSizes(cols, rows int32) {
	ui.cols = cols
	ui.rows = rows
}

func (ui *UI) GetCols() int32 {
	return ui.cols
}

func (ui *UI) GetRows() int32 {
	return ui.rows
}

// ShouldRun checks whether the game loop should be running
func (ui *UI) ShouldRun() bool {
	return ui.isRunning
}

// StartRunning starts the game loop
func (ui *UI) StartRunning() {
	ui.isRunning = true
}

// Close everything
func (ui *UI) Close() {
	matrigo.Close()
}

// GetButton return a button struct from the initialized buttons
func (ui *UI) GetButton(buttonID string) *minesweeper.Button {
	return buttons[buttonID]
}
