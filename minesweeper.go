package minesweeper

type Board struct {
	Height int
	Width  int
	Mines  int
}

type Game interface {
	Start()
	Exit()
	Quit()
}

type UI interface {
	Draw()
	ManageInput()
	ShouldRun() bool
	StartRunning()
	StopRunning()
}
