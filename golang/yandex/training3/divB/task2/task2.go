package main

import (
	"bufio"
	"fmt"
	"os"
)

// TL
// https://contest.yandex.ru/contest/45468/problems/2/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var k int
	fmt.Fscan(reader, &k)

	var str string
	fmt.Fscan(reader, &str)

	maxBeauty := 0
	for i := 0; i < len(str); i++ {
		pointer := i
		letter := str[i]
		changed := 0
		for changed < k && pointer < len(str) {
			currentLetter := str[pointer]
			if letter != currentLetter {
				changed++
			}
			pointer++
		}
		for pointer < len(str) && str[pointer] == letter {
			pointer++
		}

		if pointer-i > maxBeauty {
			maxBeauty = pointer - i
		}
	}

	fmt.Println(maxBeauty)
}
