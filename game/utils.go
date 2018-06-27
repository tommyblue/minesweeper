package game

import (
	"fmt"

	"github.com/tommyblue/minesweeper"
)

func printTiles(tiles [][]minesweeper.Tile) {
	for _, x := range tiles {
		fmt.Printf("%v\n", x)
	}
}
