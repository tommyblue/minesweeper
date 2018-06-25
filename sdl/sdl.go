package sdl

func InitSdl(c *SdlConf) error {
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

func Draw() {
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
