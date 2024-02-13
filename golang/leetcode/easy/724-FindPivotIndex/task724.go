package main

func pivotIndex(nums []int) int {
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}

	leftSum := 0
	for idx, v := range nums {
		totalSum -= v

		if leftSum == totalSum {
			return idx
		}

		leftSum += v
	}

	return -1
}
