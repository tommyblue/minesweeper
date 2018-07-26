package fixtures

import "github.com/tommyblue/minesweeper"

type UIMock struct {
	UpdateStateCalled  int
	StartRunningCalled int
	ManageInputCalled  int
	ShouldRunCalled    int
	DrawCalled         int
	LoopsToMake        int
}

func (ui *UIMock) Draw(s minesweeper.State, b *minesweeper.Board) {
	ui.DrawCalled++
}
func (ui *UIMock) ManageInput() {
	ui.ManageInputCalled++
}
func (ui *UIMock) UpdateState(e minesweeper.GameEventCallbacks) {
	ui.UpdateStateCalled++
}
func (ui *UIMock) ShouldRun() bool {
	ui.ShouldRunCalled++
	if ui.LoopsToMake > 0 {
		ui.LoopsToMake--
		return true
	}
	return false
}
func (ui *UIMock) StartRunning() {
	ui.StartRunningCalled++
}
func (ui *UIMock) StopRunning() {

}
func (ui *UIMock) GetButton(string) *minesweeper.Button {
	return &minesweeper.Button{}
}
func (ui *UIMock) SetTileSizes(int32, int32) {

}
func (ui *UIMock) GetCols() int32 {
	return 0
}
func (ui *UIMock) GetRows() int32 {
	return 0
}
