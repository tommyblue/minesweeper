package graphy

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type FontConfig struct {
	Path string
	Size int
}

type GraphyConf struct {
	Title           string
	Fonts           map[string]FontConfig
	Debug           bool
	BackgroundColor *[4]uint8
	ImagesToCache   *map[string]string
}

type imageStruct struct {
	id    string
	image *sdl.Texture
	rect  sdl.Rect
}

type sdlWrapper struct {
	cache    *map[string]imageStruct
	conf     *GraphyConf
	window   *sdl.Window
	renderer *sdl.Renderer

	countedFrames uint32
	timer         time.Time
	previousTimer time.Time
}

type Tile struct {
	ImageID string
	PosX    int32
	PosY    int32
}

type Matrix struct {
	Tiles *[]Tile
}
