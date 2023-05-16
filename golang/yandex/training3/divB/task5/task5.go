package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/5/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	letters := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &letters[i])
	}

	beauty := 0
	for i := 0; i < n-1; i++ {
		beauty += min(letters[i], letters[i+1])
	}

	fmt.Println(beauty)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
