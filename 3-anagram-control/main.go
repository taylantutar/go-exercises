package main

import (
	"algorithm-exercises/3-anagram-control/helper"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("1. kelimeyi giriniz")
	w1, _ := helper.ReadTextFromConsole()
	fmt.Println("1. kelime :" + w1)

	fmt.Println("2. kelimeyi giriniz")
	w2, _ := helper.ReadTextFromConsole()
	fmt.Println("2. kelime :" + w2)

	if controlAnagram(strings.ToLower(w1), strings.ToLower(w2)) {
		fmt.Println("Bu iki kelime anagramdır..")
	} else {
		fmt.Println("Bu iki kelime anagram değildir!")
	}
}

func controlAnagram(w1 string, w2 string) bool {

	if len(w1) != len(w2) {
		return false
	}

	m1 := make(map[string]int)

	fmt.Println(m1)

	fmt.Println("---------------")

	for _, c := range w1 {
		fmt.Println(string(c))
		if m1[string(c)] == 0 {
			m1[string(c)] = 1
		} else {
			value := m1[string(c)]
			value += 1
			m1[string(c)] = value
		}
	}

	fmt.Println(m1)

	// for key, value := range m1 {
	// 	if key != "" {
	// 		fmt.Println(key, ":", value)
	// 	}
	// }

	return true
}
