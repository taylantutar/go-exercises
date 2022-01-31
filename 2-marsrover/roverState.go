package main

import "fmt"

type IRoverState interface {
	goForward()
	goBack()
	turnRight()
	turnLeft()
}

type NorthState struct {
	coordinate *coordinate
}

func (state NorthState) goForward() {
	state.coordinate.Y += 1
	fmt.Println("x: ", state.coordinate.X, "y: ", state.coordinate.Y)
}

func (state NorthState) goBack() {
	state.coordinate.Y -= 1
	fmt.Println("x: ", state.coordinate.X, "y: ", state.coordinate.Y)
}

func (state NorthState) turnRight() {
	fmt.Println("Turn Right")
}

func (state NorthState) turnLeft() {
	fmt.Println("Turn Left")
}
