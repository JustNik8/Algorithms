package main

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}
