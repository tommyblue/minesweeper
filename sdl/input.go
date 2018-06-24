package sdl

import "github.com/veandco/go-sdl2/sdl"

type Input interface {
	Quit()
}

func ManageInput(inputInterface Input) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			inputInterface.Quit()
			break
		}
	}
}
