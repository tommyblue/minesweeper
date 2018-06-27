package ui

import (
	"fmt"

	"github.com/tommyblue/minesweeper"
	"github.com/tommyblue/minesweeper/graphy"
)

func (ui *UI) Draw(board *minesweeper.Board) {
	// prepare matrix
	var tiles []graphy.Tile
	for x, tile := range board.Tiles {
		for y, t := range tile {
			tiles = append(tiles, graphy.Tile{
				ImageID: string(tileImages[t]),
				PosX:    int32(x),
				PosY:    int32(y),
			})
		}
	}
	matrix := &graphy.Matrix{
		Tiles: &tiles,
	}
	graphy.Draw(matrix)
}

// returns a map with key of the image and its path
func getImagesToCache() *map[string]string {
	ret := make(map[string]string)
	for _, v := range tileImages {
		ret[string(v)] = getAbsolutePath(fmt.Sprintf("../assets/images/%s.png", v))
	}
	return &ret
}
