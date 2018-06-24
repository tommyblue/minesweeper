package sdl

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type FontConfig struct {
	Path string
	Size int
}
type SdlConf struct {
	Title string
	Fonts map[string]FontConfig
}

type sdlWrapper struct {
	conf     *SdlConf
	window   *sdl.Window
	renderer *sdl.Renderer

	countedFrames uint32
	timer         time.Time
	previousTimer time.Time
}

var ui *sdlWrapper

// Init UI components
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
	ui.renderer.SetDrawColor(167, 125, 83, 255)
	ui.renderer.Clear()
}

func (ui *sdlWrapper) syncFPS() {
	// Reset timers
	ui.previousTimer = ui.timer
	ui.timer = time.Now()

	tick := time.Now()
	elapsedMS := float64(tick.Sub(ui.timer)) / float64(time.Millisecond)
	if sleep := TICKSPERFRAME - elapsedMS; sleep > 0 {
		d := time.Duration(sleep)
		sdl.Delay(uint32(d))
	}

	debugFPS()
}
