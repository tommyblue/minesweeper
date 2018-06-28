package game

type eventCallbacks struct {
	leftClick  func(x, y int32)
	rightClick func(x, y int32)
}

func (ev eventCallbacks) OnLeftClick(x, y int32) {
	ev.leftClick(x, y)
}

func (ev eventCallbacks) OnRightClick(x, y int32) {
	ev.rightClick(x, y)
}
