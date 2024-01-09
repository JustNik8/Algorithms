package main

func missingNumber1Slower(nums []int) int {
	ans := 0
	for i := 0; i < len(nums)+1; i++ {
		ans += i
	}

	for i := 0; i < len(nums); i++ {
		ans -= nums[i]
	}

	return ans
}

// Выше описсано более подробное решение. А это решение более быстрое, но более хитрое
func missingNumber1(nums []int) int {
	ans := len(nums)

	for i := 0; i < len(nums); i++ {
		ans += i - nums[i]
	}

	return ans
}
