package graphy

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Input interface {
	MouseLeftClickDown(x, y int32)
	MouseRightClickDown(x, y int32)
	Quit()
}

func ManageInput(inputInterface Input) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.MouseButtonEvent:
			switch event.(*sdl.MouseButtonEvent).Type {
			case sdl.MOUSEBUTTONDOWN:
				switch ev.Button {
				case sdl.BUTTON_LEFT:
					if ui.conf.Debug {
						fmt.Printf("Mouse left click at (%d, %d)\n", ev.X, ev.Y)
					}
					inputInterface.MouseLeftClickDown(ev.X, ev.Y)
					break
				case sdl.BUTTON_RIGHT:
					if ui.conf.Debug {
						fmt.Printf("Mouse right click at (%d, %d)\n", ev.X, ev.Y)
					}
					inputInterface.MouseRightClickDown(ev.X, ev.Y)
					break
				}
				break
			}
			break
		case *sdl.QuitEvent:
			println("Quit")
			inputInterface.Quit()
			break
		}
	}
}
