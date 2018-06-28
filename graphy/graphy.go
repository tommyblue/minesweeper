package graphy

func InitGraphy(c *GraphyConf) error {
	ui = &sdlWrapper{conf: c}

	err := ui.initSdl()
	if err != nil {
		return err
	}

	err = ui.initFonts()
	if err != nil {
		return err
	}

	err = ui.initWindow()
	if err != nil {
		return err
	}

	err = ui.initRenderer()
	if err != nil {
		return err
	}

	err = ui.initTimers()
	if err != nil {
		return err
	}

	err = ui.warmCache()

	return err
}

func Close() error {
	err := ui.closeRenderer()
	if err != nil {
		return err
	}

	err = ui.closeWindow()
	if err != nil {
		return err
	}

	ui.closeSdl()
	return nil
}

func Draw(matrix *Matrix) {
	ui.drawBackground()

	for _, tile := range *matrix.Tiles {
		ui.drawImage(tile, getMatrixOffset(matrix))
	}

	ui.syncFPS()
	ui.countedFrames++

	ui.renderer.Present()
	if ui.conf.BackgroundColor != nil {
		ui.renderer.SetDrawColor(
			ui.conf.BackgroundColor[0],
			ui.conf.BackgroundColor[1],
			ui.conf.BackgroundColor[2],
			ui.conf.BackgroundColor[3],
		)
	}
	ui.renderer.Clear()
}

/*
Calculate draw offset of the matrix, currently commented out because mouse events must be
re-calculated based on offset
*/
func getMatrixOffset(matrix *Matrix) *imageOffset {
	// matrixSide := int32(math.Sqrt(float64(len(*matrix.Tiles))))
	// matrixWidth := matrixSide * ui.conf.TileSize
	// xOffset := int32((ui.conf.Window.Width - matrixWidth) / 2)
	// offset := &imageOffset{
	// 	x: xOffset,
	// 	y: 0,
	// }
	offset := &imageOffset{
		x: 0,
		y: 0,
	}
	return offset
}
