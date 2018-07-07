package main

import (
	"math/rand"
	"time"

	"github.com/tommyblue/minesweeper/game"
	"github.com/tommyblue/minesweeper/ui"
)

func main() {
	rand.Seed(time.Now().Unix())

	ui := ui.Initialize(24)
	defer ui.Close()
	game := game.Setup(ui)

	game.Start()
	defer game.Exit()
}
