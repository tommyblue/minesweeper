package ui

import (
	"fmt"

	"github.com/tommyblue/matrigo"
	"github.com/tommyblue/minesweeper"
)

func (ui *UI) Draw(state minesweeper.State, board *minesweeper.Board) {
	switch state {
	case minesweeper.InitialScreen:
		ui.drawInitialScreen()
		break
	case minesweeper.InAGame:
		ui.drawGame(board)
		break
	}
}

func (ui *UI) drawInitialScreen() {
	matrix := &matrigo.Matrix{
		Tiles: &[]matrigo.Tile{
			{
				ImageID: "button_new",
				PosX:    0,
				PosY:    0,
			},
		},
	}
	matrigo.Draw(matrix)
}

func (ui *UI) drawGame(board *minesweeper.Board) {
	// prepare matrix
	var tiles []matrigo.Tile
	for x, tile := range board.Tiles {
		for y, t := range tile {
			tiles = append(tiles, matrigo.Tile{
				ImageID: string(tileImages[t]),
				PosX:    int32(x),
				PosY:    int32(y),
			})
		}
	}
	matrix := &matrigo.Matrix{
		Tiles: &tiles,
	}
	matrigo.Draw(matrix)
}

// returns a map with key of the image and its path
func getImagesToCache() *map[string]string {
	ret := make(map[string]string)
	for _, v := range tileImages {
		ret[string(v)] = getAbsolutePath(fmt.Sprintf("../assets/images/%s.png", v))
	}
	ret["button_new"] = getAbsolutePath("../assets/images/button_new.png")
	return &ret
}
