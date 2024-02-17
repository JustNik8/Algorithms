package main

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 0; i < n; i++ {
		for j := 1; i+j*j <= n; j++ {
			square := j * j
			if dp[i+square] == 0 {
				dp[i+square] = dp[i] + 1
			} else {
				dp[i+square] = min(dp[i+square], dp[i]+1)
			}
		}
	}

	return dp[n]
}
