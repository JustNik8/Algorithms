package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/17/
func main() {
	const cnt = 5

	reader := bufio.NewReader(os.Stdin)

	first := make([]int, cnt)
	second := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		fmt.Fscan(reader, &first[i])
	}
	for i := 0; i < cnt; i++ {
		fmt.Fscan(reader, &second[i])
	}

	steps := 0
	for len(first) > 0 && len(second) > 0 {
		if steps > 1000000 {
			fmt.Println("botva")
			return
		}
		steps++

		firstCard := first[0]
		secondCard := second[0]

		first = first[1:]
		second = second[1:]

		if firstCard == 0 && secondCard == 9 {
			first = append(first, firstCard, secondCard)
		} else if secondCard == 0 && firstCard == 9 {
			second = append(second, firstCard, secondCard)
		} else if firstCard > secondCard {
			first = append(first, firstCard, secondCard)
		} else {
			second = append(second, firstCard, secondCard)

		}
	}

	if len(first) == 0 {
		fmt.Println("second", steps)
	} else {
		fmt.Println("first", steps)
	}
}
