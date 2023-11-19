package main

func lengthOfLIS(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = 1
	}

	lis := dp[0]
	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		if dp[i] > lis {
			lis = dp[i]
		}
	}

	return lis
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
