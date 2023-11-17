package main

// More memory efficient solution
func maxProduct(nums []int) int {
	size := len(nums)
	prevMax, prevMin := nums[0], nums[0]

	maxProduct := nums[0]

	for i := 1; i < size; i++ {
		min := prevMin * nums[i]
		max := prevMax * nums[i]

		if nums[i] < 0 {
			min, max = max, min
		}

		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}

		prevMax, prevMin = max, min
		if max > maxProduct {
			maxProduct = max
		}
	}

	return maxProduct
}

// Less memory efficient solution
type value struct {
	min int
	max int
}

func maxProductWithSlice(nums []int) int {
	size := len(nums)
	dp := make([]value, size)
	dp[0] = value{min: nums[0], max: nums[0]}

	maxProduct := nums[0]

	for i := 1; i < size; i++ {
		min := dp[i-1].min * nums[i]
		max := dp[i-1].max * nums[i]

		if nums[i] < 0 {
			min, max = max, min
		}

		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}

		dp[i] = value{min: min, max: max}
		if max > maxProduct {
			maxProduct = max
		}
	}

	return maxProduct
}
