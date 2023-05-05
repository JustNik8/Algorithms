package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://contest.yandex.ru/contest/28970/problems/C/
func main() {
	words := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		words[word] += 1
	}

	keys := make([]string, 0)
	for k := range words {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		keyI := keys[i]
		keyJ := keys[j]

		if words[keyI] == words[keyJ] {
			return keyI < keyJ
		} else {
			return words[keyI] > words[keyJ]
		}
	})

	for _, k := range keys {
		fmt.Println(k)
	}
}
