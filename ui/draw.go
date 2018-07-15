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
	case minesweeper.SelectLevel:
		ui.drawLevelSelection()
		break
	case minesweeper.InAGame:
		ui.drawGame(board)
	case minesweeper.Lost:
		ui.drawLost(board)
		break
	}
}

func (ui *UI) drawLost(board *minesweeper.Board) {
	newButton := ui.getInitialButtons()
	gameBoard := ui.getGameToDraw(board)
	tiles := &[]matrigo.Tile{}
	*tiles = append(*tiles, *gameBoard.Tiles...)
	*tiles = append(*tiles, *newButton.Tiles...)
	matrix := &matrigo.Matrix{
		Tiles: tiles,
	}
	matrigo.Draw(matrix)
}

func (ui *UI) drawInitialScreen() {
	matrix := ui.getInitialButtons()
	matrigo.Draw(matrix)
}

func (ui *UI) getInitialButtons() *matrigo.Matrix {
	var tiles []matrigo.Tile
	for i, id := range []string{"button_new", "button_quit"} {
		b := ui.GetButton(id)
		// Draw buttons in the centered position
		b.X, b.Y = centerImgPosition(b.W, b.H, int32(i), int32(2))
		tile := matrigo.Tile{
			ImageID: id,
			PosX:    0,
			PosY:    int32(i),
			OffsetX: b.X,
			OffsetY: b.Y,
		}
		tiles = append(tiles, tile)
	}
	matrix := &matrigo.Matrix{
		Tiles: &tiles,
	}
	return matrix
}

func (ui *UI) drawLevelSelection() {
	var tiles []matrigo.Tile
	for i, id := range []string{"button_easy", "button_medium", "button_hard"} {
		b := ui.GetButton(id)
		// Draw buttons in the centered position
		b.X, b.Y = centerImgPosition(b.W, b.H, int32(i), int32(3))
		tile := matrigo.Tile{
			ImageID: id,
			PosX:    0,
			PosY:    int32(i),
			OffsetX: b.X,
			OffsetY: b.Y,
		}
		tiles = append(tiles, tile)
	}
	matrix := &matrigo.Matrix{
		Tiles: &tiles,
	}
	matrigo.Draw(matrix)
}

func (ui *UI) drawGame(board *minesweeper.Board) {
	matrix := ui.getGameToDraw(board)
	matrigo.Draw(matrix)
}

func (ui *UI) getGameToDraw(board *minesweeper.Board) *matrigo.Matrix {
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
	return matrix
}

// returns a map with key of the image and its path
func getImagesToCache() *map[string]string {
	ret := make(map[string]string)
	for _, v := range tileImages {
		ret[string(v)] = getAbsolutePath(fmt.Sprintf("../assets/images/tiles/%s.png", v))
	}
	for k, v := range buttons {
		ret[k] = v.Src
	}
	ret["button_easy"] = getAbsolutePath("../assets/images/buttons/easy.png")
	ret["button_medium"] = getAbsolutePath("../assets/images/buttons/medium.png")
	ret["button_hard"] = getAbsolutePath("../assets/images/buttons/hard.png")
	return &ret
}
