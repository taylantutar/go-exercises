package main

import (
	"fmt"
)

func main() {
	fmt.Println("1. Kelime")
	w1, err := ReadTextFromConsole()

	if err != nil {
		fmt.Printf("%d", err)
	}

	fmt.Println("1. kelime :" + w1)

	fmt.Println("2. Kelime")
	w2, err := ReadTextFromConsole()

	if err != nil {
		fmt.Printf("%d", err)
	}

	fmt.Println("2. kelime :" + w2)

}
