package main

import "fmt"

type IRoverState interface {
	goForward()
	goBack()
	turnRight() IRoverState
	turnLeft() IRoverState
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

func (state NorthState) turnRight() IRoverState {
	return EastState{
		coordinate: state.coordinate,
	}
}

func (state NorthState) turnLeft() IRoverState {
	return WestState{
		coordinate: state.coordinate,
	}
}

type SouthState struct {
	coordinate *coordinate
}

func (state SouthState) goForward() {
	state.coordinate.Y -= 1
	fmt.Println("x: ", state.coordinate.X, "y: ", state.coordinate.Y)
}

func (state SouthState) goBack() {
	state.coordinate.Y += 1
	fmt.Println("x: ", state.coordinate.X, "y: ", state.coordinate.Y)
}

func (state SouthState) turnRight() IRoverState {
	return WestState{
		coordinate: state.coordinate,
	}
}

func (state SouthState) turnLeft() IRoverState {
	return EastState{
		coordinate: state.coordinate,
	}
}

type WestState struct {
	coordinate *coordinate
}

func (state WestState) goForward() {
	state.coordinate.X -= 1
	fmt.Println("x: ", state.coordinate.X, "y: ", state.coordinate.Y)
}

func (state WestState) goBack() {
	state.coordinate.X += 1
	fmt.Println("x: ", state.coordinate.X, "y: ", state.coordinate.Y)
}

func (state WestState) turnRight() IRoverState {
	return NorthState{
		coordinate: state.coordinate,
	}
}

func (state WestState) turnLeft() IRoverState {
	return SouthState{
		coordinate: state.coordinate,
	}
}

type EastState struct {
	coordinate *coordinate
}

func (state EastState) goForward() {
	state.coordinate.X += 1
	fmt.Println("x: ", state.coordinate.X, "y: ", state.coordinate.Y)
}

func (state EastState) goBack() {
	state.coordinate.X -= 1
	fmt.Println("x: ", state.coordinate.X, "y: ", state.coordinate.Y)
}

func (state EastState) turnRight() IRoverState {
	return SouthState{
		coordinate: state.coordinate,
	}
}

func (state EastState) turnLeft() IRoverState {
	return NorthState{
		coordinate: state.coordinate,
	}
}
