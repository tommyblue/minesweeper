package ui

import "github.com/tommyblue/minesweeper"

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

// global var that is passed to matrigo to make callbacks on events. It must implement the
// matrigo.Input interface
var input *matrigoInputInterface

var buttons = map[string]*minesweeper.Button{
	"button_new": &minesweeper.Button{
		Src: getAbsolutePath("../assets/images/buttons/new_game.png"),
		W:   190,
		H:   49,
	},
	"button_quit": &minesweeper.Button{
		Src: getAbsolutePath("../assets/images/buttons/quit.png"),
		W:   190,
		H:   49,
	},
	"button_easy": &minesweeper.Button{
		Src: getAbsolutePath("../assets/images/buttons/easy.png"),
		W:   190,
		H:   49,
	},
	"button_medium": &minesweeper.Button{
		Src: getAbsolutePath("../assets/images/buttons/medium.png"),
		W:   190,
		H:   49,
	},
	"button_hard": &minesweeper.Button{
		Src: getAbsolutePath("../assets/images/buttons/hard.png"),
		W:   190,
		H:   49,
	},
}
