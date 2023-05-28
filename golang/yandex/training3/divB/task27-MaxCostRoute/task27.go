package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/27/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &table[i][j])
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = table[0][0]
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + table[0][i]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + table[i][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			left := dp[i][j-1] + table[i][j]
			up := dp[i-1][j] + table[i][j]
			dp[i][j] = int(math.Max(float64(left), float64(up)))
		}
	}

	fmt.Println(dp[n-1][m-1])

	ans := make([]string, 0, m+n-2)
	curI, curJ := n-1, m-1
	for curI != 0 || curJ != 0 {
		if curI == 0 {
			ans = append(ans, "R")
			curJ--
		} else if curJ == 0 {
			ans = append(ans, "D")
			curI--
		} else {
			left := dp[curI][curJ-1]
			up := dp[curI-1][curJ]
			if left > up {
				ans = append(ans, "R")
				curJ--
			} else {
				ans = append(ans, "D")
				curI--
			}
		}
	}

	for i := m + n - 3; i >= 0; i-- {
		fmt.Print(ans[i], " ")
	}

}
