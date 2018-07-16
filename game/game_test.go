package game

import (
	"testing"

	"github.com/tommyblue/minesweeper/ui"
)

func TestGame_Setup(t *testing.T) {
	ui := ui.Initialize(24)
	defer ui.Close()
	g := Setup(ui)
	if g.UI == nil {
		t.Errorf("Wrong game ui initialization")
	}
	if g.EventCallbacks == nil {
		t.Errorf("Wrong callbacks initialization")
	}
	if g.State == nil {
		t.Errorf("Wrong State initialization")
	}
}
