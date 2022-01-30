package main

import (
	"fmt"
	"strings"
)

func findConconantCharCount(word string) int {

	wordLower := strings.ToLower(word)
	vovelCount := 0
	conconantCount := 0
	vovels := "aeioü"

	for _, c := range wordLower {
		if c <= 122 && c >= 97 {
			if strings.Contains(vovels, string(c)) {
				vovelCount++
			} else {
				conconantCount++
			}
		}
	}

	return conconantCount
}

func main() {
	var result = findConconantCharCount("**--AaBbCcdi-**")
	fmt.Println(result)
}
