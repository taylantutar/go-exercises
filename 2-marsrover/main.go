package main

import "fmt"

func main() {
	fmt.Println("******************")

	rover := rover{
		roverState: NorthState{
			coordinate: &coordinate{
				X: 1,
				Y: 1,
			},
		},
	}

	for i := 0; i < 2; i++ {
		rover.goForward()
	}
}
