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

// This struct must implement graphy.Input interface, managing all possible input events.
// It contains the methods to be mapped to actual ui methods.
type graphyInputInterface struct {
	mouseLeftClickDownFn func(x, y int32)
	quitFn               func()
}

// global var that is passed to graphy to make callbacks on events. It must implement the
// graphy.Input interface
var input *graphyInputInterface
