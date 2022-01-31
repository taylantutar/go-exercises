package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("******************")

	rover := rover{
		roverState: EastState{
			coordinate: &coordinate{
				X: 1,
				Y: 1,
			},
		},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case 'R':
			rover.TurnRight()
			break
		case 'L':
			rover.TurnLeft()
			break
		case 'F':
			rover.goForward()
			break
		case 'B':
			rover.goBack()
			break
		}
	}
}
