package ui

import "github.com/tommyblue/minesweeper"

// UpdateState checks if any event happened and updates the game state accordingly
func (ui *UI) UpdateState(ev minesweeper.GameEventCallbacks) {
	if ui.event != nil {
		switch ui.event.evType {
		case mouseLeftClick:
			ev.OnLeftClick(ui.event.tile.x, ui.event.tile.y)
			break
		case mouseRightClick:
			ev.OnRightClick(ui.event.tile.x, ui.event.tile.y)
			break
		}
		// Reset the event
		ui.event = nil
	}
}
