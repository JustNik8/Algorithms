package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/30/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	first := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &first[i])
	}

	var m int
	fmt.Fscan(reader, &m)

	second := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(reader, &second[i])
	}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if second[j] == first[i] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	length := dp[n][m]
	ans := make([]int, 0, length)

	curI, curJ := n, m
	for dp[curI][curJ] > 0 {
		currentLen := dp[curI][curJ]
		if dp[curI-1][curJ] == currentLen {
			curI = curI - 1
		} else if dp[curI][curJ-1] == currentLen {
			curJ = curJ - 1
		} else {
			ans = append(ans, second[curJ])
			curI = curI - 1
			curJ = curJ - 1
		}
	}

	for i := length - 1; i >= 0; i-- {
		fmt.Print(ans[i], " ")
	}
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
