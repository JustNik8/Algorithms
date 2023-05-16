package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/8/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var k int
	fmt.Fscan(reader, &k)

	var minX, maxX, minY, maxY int

	var x, y int
	fmt.Fscan(reader, &x, &y)
	minX, maxX = x, x
	minY, maxY = y, y

	for i := 0; i < k-1; i++ {
		fmt.Fscan(reader, &x, &y)

		minX = min(minX, x)
		maxX = max(maxX, x)
		minY = min(minY, y)
		maxY = max(maxY, y)
	}

	fmt.Println(minX, minY, maxX, maxY)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
