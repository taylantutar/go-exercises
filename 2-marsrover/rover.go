package main

type rover struct {
	roverState IRoverState
}

func (r rover) TurnRight() {
	r.roverState.turnRight()
}

func (r rover) TurnLeft() {
	r.roverState.turnLeft()
}

func (r rover) goForward() {
	r.roverState.goForward()
}

func (r rover) goBack() {
	r.roverState.goBack()
}
