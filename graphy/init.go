package graphy

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func (ui *sdlWrapper) initSdl() error {
	return sdl.Init(sdl.INIT_EVERYTHING)
}

func (ui *sdlWrapper) initFonts() error {
	fonts = make(map[string]*ttf.Font)
	// Initialize TTF
	ttf.Init()
	for fontName, fontConf := range ui.conf.Fonts {
		filepath := getAbsolutePath(fontConf.Path)
		font, err := ttf.OpenFont(filepath, fontConf.Size)
		if err != nil {
			return err
		}
		fonts[fontName] = font
		if defaultFont == nil {
			defaultFont = font
		}

	}
	return nil
}

func (ui *sdlWrapper) initWindow() error {
	var err error
	ui.window, err = sdl.CreateWindow(
		ui.conf.Title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		ui.conf.Window.Width,
		ui.conf.Window.Height,
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

	err = ui.renderer.SetDrawColor(0, 0, 0, 255)
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
