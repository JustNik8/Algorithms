package main

func rob(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	rob1 := calculateDP(nums[1:size])
	rob2 := calculateDP(nums[:size-1])

	if rob1 > rob2 {
		return rob1
	} else {
		return rob2
	}
}

func calculateDP(nums []int) int {
	size := len(nums)
	dpSize := size + 2
	dp := make([]int, dpSize)

	for i := 2; i < dpSize; i++ {
		if dp[i-1] > dp[i-2]+nums[i-2] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-2] + nums[i-2]
		}
	}

	return dp[dpSize-1]
}
