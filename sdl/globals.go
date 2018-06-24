package sdl

import (
	"time"

	"github.com/veandco/go-sdl2/ttf"
)

var font *ttf.Font

// FPS are the frame per seconds target
const FPS = 60

// TICKSPERFRAME is the number of ticks required to reach FPS
const TICKSPERFRAME = 1000.0 / FPS

var lastTimestamp time.Time
var lastFPS int
var countSinceLast int
