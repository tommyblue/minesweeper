package sdl

import (
	"time"

	"github.com/tommyblue/sokoban/utils"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func (ui *sdlWrapper) initSdl() error {
	return sdl.Init(sdl.INIT_EVERYTHING)
}

func (ui *sdlWrapper) initFonts() error {
	// Initialize TTF
	ttf.Init()
	var err error
	filepath := utils.GetRelativePath(ui.fontPath)
	font, err = ttf.OpenFont(filepath, 14)
	return err
}

func (ui *sdlWrapper) initWindow() error {
	var err error
	ui.window, err = sdl.CreateWindow(
		ui.title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800,
		600,
		sdl.WINDOW_SHOWN,
	)
	return err
}

func (ui *sdlWrapper) initRenderer() error {
	var err error
	ui.renderer, err = sdl.CreateRenderer(ui.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}

	err = ui.renderer.SetDrawColor(0, 255, 0, 255)
	if err != nil {
		return err
	}
	err = ui.renderer.Clear()

	return err
}

func (ui *sdlWrapper) initTimers() error {
	ui.timer = time.Now()
	ui.previousTimer = ui.timer
	return nil
}
