package ui

import (
	"math"

	"github.com/tommyblue/matrigo"
	"github.com/tommyblue/minesweeper"
)

/*
ManageInput manages the input. To do so calls matrigo.ManageInput passing the input global var
to matrigo.
The input var must be a pointer to the matrigoInputInterface struct, which implements the
matrigo.Input interface
*/
func (ui *UI) ManageInput() {
	matrigo.ManageInput(input)
}

/*
This function initializes the input global var and maps the input events happening from matrigo to
functions in the ui package, actually doing things.
This is required because the callbacks from matrigo don't have knowledge of the `ui` object and
this is the point of linking between the matrigo interface implementation and the ui
*/
func mapInputToFn(ui *UI) {
	if input == nil {
		input = &matrigoInputInterface{
			quitFn:                ui.StopRunning,
			mouseLeftClickDownFn:  ui.mouseLeftClickAt,
			mouseRightClickDownFn: ui.mouseRightClickAt,
		}
	}
}

/*
The following functions map the callbacks called by matrigo (through the input interface)
*/
func (i *matrigoInputInterface) MouseLeftClickDown(x, y int32) {
	i.mouseLeftClickDownFn(x, y)
}
func (i *matrigoInputInterface) MouseRightClickDown(x, y int32) {
	i.mouseRightClickDownFn(x, y)
}
func (i *matrigoInputInterface) Quit() {
	i.quitFn()
}

/*
Below this comment are implemented the functions called by the matrigo interface
*/
func (ui *UI) StopRunning() {
	ui.isRunning = false
}
func (ui *UI) mouseLeftClickAt(x, y int32) {
	ui.mouseClickAt(x, y, mouseLeftClick)
}
func (ui *UI) mouseRightClickAt(x, y int32) {
	ui.mouseClickAt(x, y, mouseRightClick)
}
func (ui *UI) mouseClickAt(x, y int32, clickType eventType) {
	newX := x - (ui.window.Width / 2) + ((ui.cols * ui.tileSize) / 2)
	newY := y - (ui.window.Height / 2) + ((ui.rows * ui.tileSize) / 2)

	tileX := int32(math.Floor(float64(newX) / float64(ui.tileSize)))
	tileY := int32(math.Floor(float64(newY) / float64(ui.tileSize)))

	ui.event = &event{
		evType: clickType,
		tile: &minesweeper.Position{
			X: tileX,
			Y: tileY,
		},
		mouseClick: &minesweeper.Position{
			X: x,
			Y: y,
		},
	}
}
