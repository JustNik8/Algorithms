package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	prefSum := make([]int, n+1)
	prefSumReversed := make([]int, n+1)
	prevY := 0
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		var diff = y - prevY
		if diff > 0 {
			prefSum[i] = prefSum[i-1] + diff
			prefSumReversed[i] = prefSumReversed[i-1]
		} else {
			prefSumReversed[i] = prefSumReversed[i-1] + (-diff)
			prefSum[i] = prefSum[i-1]
		}

		prevY = y
	}

	var m int
	fmt.Fscan(reader, &m)

	for i := 0; i < m; i++ {
		var start, end int
		fmt.Fscan(reader, &start, &end)
		if end-start >= 0 {
			fmt.Println(prefSum[end] - prefSum[start])
		} else {
			fmt.Println(prefSumReversed[start] - prefSumReversed[end])
		}
	}
}
