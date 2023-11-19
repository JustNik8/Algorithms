package main

import "fmt"

// effective solution
func jump(nums []int) int {
	size := len(nums)

	maxJump := 0
	end := 0
	step := 0

	for i := 0; i < size-1; i++ {
		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}

		if i == end {
			end = maxJump
			step++
		}
	}

	return step
}

// not effective solution
func jumpDP(nums []int) int {
	size := len(nums)
	dp := make([]int, size)

	for i := 0; i < size-1; i++ {
		jump := i + nums[i]
		if jump > size-1 {
			jump = size - 1
		}
		for j := i + 1; j <= jump; j++ {
			if dp[j] == 0 {
				dp[j] = dp[i] + 1
			} else {
				dp[j] = min(dp[j], dp[i]+1)
			}
		}
	}

	fmt.Println(dp)
	return dp[size-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
