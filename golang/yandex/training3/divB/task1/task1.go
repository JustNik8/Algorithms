package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://contest.yandex.ru/contest/45468/problems/1/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	letters := make(map[uint8]int)
	maxCnt := 0
	for scanner.Scan() {
		word := scanner.Text()

		for i := 0; i < len(word); i++ {
			letter := word[i]
			if letter < '!' || letter > '~' {
				continue
			}
			letters[letter] += 1

			cnt := letters[letter]
			if cnt > maxCnt {
				maxCnt = cnt
			}
		}
	}

	sortedLetters := make([]uint8, 0, len(letters))
	for l := range letters {
		sortedLetters = append(sortedLetters, l)
	}

	sort.Slice(sortedLetters, func(i, j int) bool {
		return sortedLetters[i] < sortedLetters[j]
	})

	for i := maxCnt; i > 0; i-- {
		for j := 0; j < len(sortedLetters); j++ {
			letter := sortedLetters[j]
			cnt := letters[letter]
			if cnt >= i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	for i := 0; i < len(sortedLetters); i++ {
		fmt.Print(string(sortedLetters[i]))
	}
}
