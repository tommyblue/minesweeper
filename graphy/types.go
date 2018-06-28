package graphy

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type FontConfig struct {
	Path string
	Size int
}

type Window struct {
	Height int32
	Width  int32
}

type GraphyConf struct {
	TileSize        int32
	Title           string
	Fonts           map[string]FontConfig
	Debug           bool
	Window          *Window
	BackgroundColor *[4]uint8
	BackgroundImage string
	ImagesToCache   *map[string]string
}

type imageStruct struct {
	id    string
	image *sdl.Texture
	rect  sdl.Rect
}

type imageOffset struct {
	x int32
	y int32
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
