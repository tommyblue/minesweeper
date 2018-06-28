package ui

// UpdateState checks if any event happened and updates the game state accordingly
func (ui *UI) UpdateState(onClickOnTile func(x, y int32)) {
	if ui.event != nil {
		if ui.event.evType == mouseLeftClick {
			onClickOnTile(ui.event.tile.x, ui.event.tile.y)
		}
		// Reset the event
		ui.event = nil
	}
}
