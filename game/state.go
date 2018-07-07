package game

import (
	"github.com/tommyblue/minesweeper"
)

func (g *Game) setState(newState minesweeper.State) {
	g.State.CurrentState = newState
}

func (g *Game) selectLevel(level minesweeper.Level) {
	g.State.SelectedLevel = level
}
