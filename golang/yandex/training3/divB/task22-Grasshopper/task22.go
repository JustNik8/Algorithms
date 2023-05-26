package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/22/
func main() {
	reader := bufio.NewReader(os.Stdin)

	dp := make([]uint64, 31)

	var n, k int
	fmt.Fscan(reader, &n, &k)

	dp[0], dp[1] = 1, 1

	for i := 2; i <= k; i++ {
		dp[i] = 2 * dp[i-1]
	}

	for i := k + 1; i < n; i++ {
		sum := uint64(0)
		for j := 0; j <= k; j++ {
			sum += dp[i-j]
		}
		dp[i] = sum
	}

	fmt.Println(dp[n-1])
}
