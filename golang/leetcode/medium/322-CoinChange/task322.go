package main

func coinChange(coins []int, amount int) int {
	coinsCnt := len(coins)
	dp := make([]int, amount+1)
	for i := 0; i < coinsCnt; i++ {
		move := coins[i]
		if move <= amount {
			dp[move] = 1
		}
	}

	for i := 1; i < amount; i++ {
		if dp[i] == 0 {
			continue
		}
		for j := 0; j < coinsCnt; j++ {
			move := i + coins[j]
			if move <= amount {
				if dp[move] == 0 {
					dp[move] = dp[i] + 1
				} else {
					dp[move] = min(dp[move], dp[i]+1)
				}
			}
		}
	}

	if dp[amount] == 0 && amount != 0 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
