package main

func longestCommonSubsequence(text1 string, text2 string) int {
	rows := len(text1) + 1
	cols := len(text2) + 1

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			if text1[row-1] == text2[col-1] {
				dp[row][col] = dp[row-1][col-1] + 1
			} else {
				dp[row][col] = max(dp[row][col-1], dp[row-1][col])
			}
		}
	}

	return dp[rows-1][cols-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
