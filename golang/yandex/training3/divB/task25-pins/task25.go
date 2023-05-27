package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// https://contest.yandex.ru/contest/45468/problems/25/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	pins := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &pins[i])
	}

	sort.Ints(pins)

	dp := make([]int, n)
	dp[0] = pins[1] - pins[0]
	dp[1] = pins[1] - pins[0]

	for i := 2; i < n; i++ {
		dist := pins[i] - pins[i-1]

		dp[i] = int(math.Min(float64(dist+dp[i-2]), float64(dist+dp[i-1])))
	}

	fmt.Println(dp[n-1])
}
