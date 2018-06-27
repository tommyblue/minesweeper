package ui

import "github.com/tommyblue/minesweeper/graphy"

type graphyInputInterface struct {
	quitFn func()
}

var input *graphyInputInterface

func (ui *UI) ManageInput() {
	graphy.ManageInput(input)
}

func (i *graphyInputInterface) Quit() {
	i.quitFn()
}

func initInput(ui *UI) {
	if input == nil {
		input = &graphyInputInterface{
			quitFn: ui.StopRunning,
		}
	}
}
