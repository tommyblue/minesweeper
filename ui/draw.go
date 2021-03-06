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
	case minesweeper.Win:
		ui.drawWin()
		break
	}
}

func (ui *UI) drawLost(board *minesweeper.Board) {
	gameBoard := ui.getGameToDraw(board)
	tiles := []*matrigo.Tile{}
	tiles = append(tiles, gameBoard.Tiles...)
	matrix := &matrigo.Matrix{
		Tiles: tiles,
	}
	ui.centerMatrix(matrix)
	newButton := ui.getInitialButtons()
	matrix.Tiles = append(matrix.Tiles, newButton.Tiles...)
	matrigo.Draw(matrix)
}

func (ui *UI) drawInitialScreen() {
	matrix := ui.getInitialButtons()
	matrigo.Draw(matrix)
}

func (ui *UI) drawWin() {
	matrix := ui.getWinButtons()
	matrigo.Draw(matrix)
}

func (ui *UI) getInitialButtons() *matrigo.Matrix {
	return ui.getButtons([]string{"button_new", "button_quit"})
}

func (ui *UI) getWinButtons() *matrigo.Matrix {
	return ui.getButtons([]string{"button_win"})
}

func (ui *UI) getButtons(buttons []string) *matrigo.Matrix {
	var tiles []*matrigo.Tile
	for i, id := range buttons {
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
		tiles = append(tiles, &tile)
	}
	matrix := &matrigo.Matrix{
		Tiles: tiles,
	}
	return matrix
}

func (ui *UI) drawLevelSelection() {
	matrix := ui.getButtons([]string{"button_easy", "button_medium", "button_hard"})
	matrigo.Draw(matrix)
}

func (ui *UI) drawGame(board *minesweeper.Board) {
	matrix := ui.getGameToDraw(board)
	ui.centerMatrix(matrix)
	matrigo.Draw(matrix)
}

func (ui *UI) centerMatrix(matrix *matrigo.Matrix) {
	for _, t := range matrix.Tiles {
		t.OffsetX = (ui.window.Width / 2) - ((ui.cols * ui.tileSize) / 2)
		t.OffsetY = (ui.window.Height / 2) - ((ui.rows * ui.tileSize) / 2)
	}
}

func (ui *UI) getGameToDraw(board *minesweeper.Board) *matrigo.Matrix {
	// prepare matrix
	var tiles []*matrigo.Tile
	for x, tile := range board.Tiles {
		for y, t := range tile {
			tiles = append(tiles, &matrigo.Tile{
				ImageID: string(tileImages[t]),
				PosX:    int32(x),
				PosY:    int32(y),
			})
		}
	}
	matrix := &matrigo.Matrix{
		Tiles: tiles,
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
	ret["button_win"] = getAbsolutePath("../assets/images/buttons/win.png")
	ret["button_new"] = getAbsolutePath("../assets/images/buttons/new_game.png")
	ret["button_quit"] = getAbsolutePath("../assets/images/buttons/quit.png")
	return &ret
}
