package ui

type tile struct {
	x int32
	y int32
}

type eventType int

const (
	mouseLeftClick = iota
	mouseRightClick
)

type event struct {
	evType eventType
	tile   *tile
}

type UI struct {
	isRunning bool
	tileSize  int32
	event     *event
}

// This struct must implement graphy.Input interface, managing all possible input events.
// It contains the methods to be mapped to actual ui methods.
type graphyInputInterface struct {
	mouseLeftClickDownFn func(x, y int32)
	quitFn               func()
}

type tileImageName string