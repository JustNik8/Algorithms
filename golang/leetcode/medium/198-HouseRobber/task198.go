func rob(nums []int) int {
	n := len(nums)

	dp := make([]int, n+2)

	for i := 2; i < n+2; i++ {
		if dp[i-1] > dp[i-2]+nums[i-2] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-2] + nums[i-2]
		}
	}

	return dp[n+1]
}