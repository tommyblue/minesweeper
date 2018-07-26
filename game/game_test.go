package game

import (
	"testing"

	"github.com/tommyblue/minesweeper"

	"github.com/tommyblue/minesweeper/fixtures"
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

func TestGame_updateState(t *testing.T) {
	ui := &fixtures.UIMock{}

	g := &Game{
		UI: ui,
	}

	g.updateState()

	if g.UI.(*fixtures.UIMock).UpdateStateCalled != 1 {
		t.Errorf("UpdateState called %d times", ui.UpdateStateCalled)
	}
}

func TestGame_Start(t *testing.T) {
	ui := &fixtures.UIMock{
		LoopsToMake: 1,
	}

	g := &Game{
		UI:    ui,
		State: &minesweeper.GameState{},
	}

	g.Start()

	if g.UI.(*fixtures.UIMock).StartRunningCalled != 1 {
		t.Errorf("StartRunning called %d times", ui.StartRunningCalled)
	}

	if g.UI.(*fixtures.UIMock).ShouldRunCalled == 0 {
		t.Errorf("ShouldRun called 0 times")
	}

	if g.UI.(*fixtures.UIMock).ManageInputCalled == 0 {
		t.Errorf("ManageInput called 0 times")
	}

	if g.UI.(*fixtures.UIMock).UpdateStateCalled == 0 {
		t.Errorf("UpdateState called 0 times")
	}

	if g.UI.(*fixtures.UIMock).DrawCalled == 0 {
		t.Errorf("Draw called 0 times")
	}
}
