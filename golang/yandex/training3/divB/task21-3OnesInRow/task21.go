package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/21/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	dp := make([]uint64, 36)
	dp[1] = 2
	dp[2] = 4
	dp[3] = 7

	for i := 4; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	fmt.Println(dp[n])
}
