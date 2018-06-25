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
	Title           string
	Fonts           map[string]FontConfig
	Debug           bool
	BackgroundColor *[4]uint8
}

type sdlWrapper struct {
	conf     *SdlConf
	window   *sdl.Window
	renderer *sdl.Renderer

	countedFrames uint32
	timer         time.Time
	previousTimer time.Time
}
