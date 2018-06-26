package minesweeper

type Tile int

type Board struct {
	Cols  int32
	Rows  int32
	Mines int32
	Tiles [][]Tile
}

type GameState struct {
	DiscoveredTiles [][]bool
}

type Game interface {
	Start()
	Exit()
	Quit()
}

type UI interface {
	Draw(*Board)
	ManageInput()
	ShouldRun() bool
	StartRunning()
	StopRunning()
}
