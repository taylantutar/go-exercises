package main

import (
	"bufio"
	"os"
)

func ReadTextFromConsole() (text string, err error) {
	reader := bufio.NewReader(os.Stdin)

	text, err = reader.ReadString('\n')

	if err != nil {
		return text, err
	}

	return text, nil

}
