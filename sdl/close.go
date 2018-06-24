package sdl

import "github.com/veandco/go-sdl2/sdl"

func (ui *sdlWrapper) closeRenderer() error {
	return ui.renderer.Destroy()
}

func (ui *sdlWrapper) closeWindow() error {
	return ui.window.Destroy()
}

func (ui *sdlWrapper) closeSdl() {
	sdl.Quit()
}
