package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	first := make(map[string]string, n)
	second := make(map[string]string, n)

	for i := 0; i < n; i++ {
		var firstWord string
		var secondWord string
		fmt.Fscan(in, &firstWord, &secondWord)

		first[firstWord] = secondWord
		second[secondWord] = firstWord
	}

	var word string
	fmt.Fscan(in, &word)

	if len(first[word]) != 0 {
		fmt.Print(first[word])
	} else {
		fmt.Print(second[word])
	}
}
