package ui

import "github.com/tommyblue/minesweeper/sdl"

type sdlInputInterface struct {
	quitFn func()
}

var input *sdlInputInterface

func (ui *UI) ManageInput() {
	sdl.ManageInput(input)
}

func (i *sdlInputInterface) Quit() {
	i.quitFn()
}

func initInput(ui *UI) {
	if input == nil {
		input = &sdlInputInterface{
			quitFn: ui.StopRunning,
		}
	}
}
