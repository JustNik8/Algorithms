package main

import (
	"fmt"
)

func main() {
	first := make(map[string]string)
	second := make(map[string]string)

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var firstWord string
		var secondWord string
		fmt.Scan(&firstWord)
		fmt.Scan(&secondWord)

		first[firstWord] = secondWord
		second[secondWord] = firstWord
	}

	var word string
	fmt.Scan(&word)

	if len(first[word]) != 0 {
		fmt.Println(first[word])
	} else {
		fmt.Println(second[word])
	}
}
