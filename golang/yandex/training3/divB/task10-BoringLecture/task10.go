package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var word string
	fmt.Fscan(reader, &word)

	words := make(map[uint8]int)

	for i := 0; i < len(word); i++ {
		letter := word[i]
		pref := i + 1
		post := len(word) - i

		words[letter] += pref * post
	}

	keys := make([]uint8, 0, len(words))
	for l, _ := range words {
		keys = append(keys, l)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, key := range keys {
		fmt.Printf("%s: %d\n", string(key), words[key])
	}
}
