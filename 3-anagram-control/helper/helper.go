package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadTextFromConsole() (text string, err error) {
	reader := bufio.NewReader(os.Stdin)

	text, err = reader.ReadString('\n')

	if err != nil {
		return text, err
	}

	text = strings.TrimSpace(text)
	fmt.Println(len(text))

	return text, nil

}
