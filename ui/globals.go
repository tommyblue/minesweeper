package ui

import "github.com/tommyblue/minesweeper"

type tileImageName string

// tileImages is a map holding the name of the images
var tileImages = map[minesweeper.Tile]tileImageName{
	minesweeper.Empty:     "empty",
	minesweeper.N1:        "n1",
	minesweeper.N2:        "n2",
	minesweeper.N3:        "n3",
	minesweeper.N4:        "n4",
	minesweeper.N5:        "n5",
	minesweeper.N6:        "n6",
	minesweeper.N7:        "n7",
	minesweeper.N8:        "n8",
	minesweeper.Mine:      "mine",
	minesweeper.Explosion: "explosion",
	minesweeper.Flag:      "flag",
	minesweeper.Unknown:   "unknown",
}
