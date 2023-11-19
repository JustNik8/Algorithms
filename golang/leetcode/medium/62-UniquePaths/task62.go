package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if col-1 >= 0 {
				dp[row][col] += dp[row][col-1]
			}
			if row-1 >= 0 {
				dp[row][col] += dp[row-1][col]
			}
		}
	}

	return dp[m-1][n-1]
}
