package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/28/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	dp := make([][]uint64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]uint64, m)
	}

	dpi := []int{2, 1}
	dpj := []int{1, 2}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < len(dpi); k++ {
				neighbourI := i - dpi[k]
				neighbourJ := j - dpj[k]
				if neighbourI >= 0 && neighbourJ >= 0 {
					dp[i][j] += dp[neighbourI][neighbourJ]
				}
			}
		}
	}

	fmt.Println(dp[n-1][m-1])

}

