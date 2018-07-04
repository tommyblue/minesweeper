package ui

import "github.com/tommyblue/minesweeper"

// UpdateState checks if any event happened and updates the game state accordingly
func (ui *UI) UpdateState(ev minesweeper.GameEventCallbacks) {
	// TODO: move to channels
	if ui.event != nil {
		switch ui.event.evType {
		case mouseLeftClick:
			ev.OnLeftClick(ui.event.tile, ui.event.mouseClick)
			break
		case mouseRightClick:
			ev.OnRightClick(ui.event.tile, ui.event.mouseClick)
			break
		}
		// Reset the event
		ui.event = nil
	}
}
