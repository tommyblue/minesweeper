package ui

import (
	"math"

	"github.com/tommyblue/minesweeper/graphy"
)

/*
ManageInput manages the input. To do so calls graphy.ManageInput passing the input global var
to graphy.
The input var must be a pointer to the graphyInputInterface struct, which implements the
graphy.Input interface
*/
func (ui *UI) ManageInput() {
	graphy.ManageInput(input)
}

/*
This function initializes the input global var and maps the input events happening from graphy to
functions in the ui package, actually doing things.
This is required because the callbacks from graphy don't have knowledge of the `ui` object and
this is the point of linking between the graphy interface implementation and the ui
*/
func mapInputToFn(ui *UI) {
	if input == nil {
		input = &graphyInputInterface{
			quitFn:               ui.stopRunning,
			mouseLeftClickDownFn: ui.mouseClickAt,
		}
	}
}

/*
The following functions map the callbacks called by graphy (through the input interface)
*/
func (i *graphyInputInterface) MouseLeftClickDown(x, y int32) {
	i.mouseLeftClickDownFn(x, y)
}
func (i *graphyInputInterface) Quit() {
	i.quitFn()
}

/*
Below this comment are implemented the functions called by the graphy interface
*/
func (ui *UI) stopRunning() {
	ui.isRunning = false
}
func (ui *UI) mouseClickAt(x, y int32) {
	tileX := int32(math.Floor(float64(x) / float64(ui.tileSize)))
	tileY := int32(math.Floor(float64(y) / float64(ui.tileSize)))
	ui.event = &event{
		evType: "mouseLeftClick",
		tile: &tile{
			x: tileX,
			y: tileY,
		},
	}
}
