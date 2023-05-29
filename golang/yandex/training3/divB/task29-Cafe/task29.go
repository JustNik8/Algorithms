package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://contest.yandex.ru/contest/45468/problems/29/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	costs := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &costs[i+1])
	}
	if n == 0 {
		fmt.Println(0)
		fmt.Println(0, 0)
		return
	} else if n == 1 {
		fmt.Println(costs[1])
		fmt.Println(1, 0)
		return
	}

	dp := make([][]int, n+1)
	transfers := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = 100_000
		}
		transfers[i] = make([]int, n+1)

	}

	dp[0][0] = 0

	for i := 1; i < n+1; i++ {
		for j := 0; j < n; j++ {
			if costs[i] <= 100 {
				// dp[i][j] = min(dp[i-1][j]+costs[i], dp[i-1][j+1])
				if dp[i-1][j]+costs[i] < dp[i-1][j+1] {
					dp[i][j] = dp[i-1][j] + costs[i]
					transfers[i][j] = j
				} else {
					dp[i][j] = dp[i-1][j+1]
					transfers[i][j] = j + 1
				}

			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][j+1]
					transfers[i][j] = j + 1
					continue
				}
				//dp[i][j] = min(dp[i-1][j-1]+costs[i], dp[i-1][j+1])
				if dp[i-1][j-1]+costs[i] < dp[i-1][j+1] {
					dp[i][j] = dp[i-1][j-1] + costs[i]
					transfers[i][j] = j - 1
				} else {
					dp[i][j] = dp[i-1][j+1]
					transfers[i][j] = j + 1
				}
			}
		}
	}

	minSum := dp[n][0]
	minSumIndex := 0
	for i := 1; i < n+1; i++ {
		if dp[n][i] <= minSum {
			minSum = dp[n][i]
			minSumIndex = i
		}
	}

	days := make([]int, 0)
	index := minSumIndex
	for i := n; i >= 0; i-- {
		next := transfers[i][index]
		if next-index > 0 {
			days = append(days, i)
		}
		index = next
	}

	fmt.Println(minSum)
	fmt.Println(minSumIndex, len(days))

	sort.Ints(days)
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
}
