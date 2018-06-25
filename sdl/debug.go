package sdl

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func debugFPS() {
	if ui.conf.Debug {
		var zeroTime time.Time
		if lastTimestamp == zeroTime {
			lastTimestamp = time.Now()
		} else {
			if time.Since(lastTimestamp) < time.Duration(time.Second) {
				countSinceLast++
			} else {
				lastTimestamp = time.Now()
				if os.Getenv("FPSLOG") == "1" {
					log.Printf("[%v] update FPS: %v -> %v\n", lastTimestamp, lastFPS, countSinceLast+1)
				}
				lastFPS = countSinceLast + 1
				countSinceLast = 0
			}
			drawFPS(fmt.Sprintf("%v", lastFPS))
		}
	}
}

func drawFPS(text string) error {
	fontColor := sdl.Color{R: 0, G: 0, B: 0, A: 0}
	textSurface, err := defaultFont.RenderUTF8Blended(fmt.Sprintf("%s FPS", text), fontColor)
	if err != nil {
		return err
	}
	defer textSurface.Free()

	textTexture, err := ui.renderer.CreateTextureFromSurface(textSurface)
	if err != nil {
		return err
	}
	defer textTexture.Destroy()

	src := sdl.Rect{X: 0, Y: 0, W: textSurface.W, H: textSurface.H}
	dest := sdl.Rect{X: 0, Y: 0, W: textSurface.W, H: textSurface.H}
	ui.renderer.Copy(textTexture, &src, &dest)

	return nil
}
