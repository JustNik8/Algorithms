package main

func sortedSquares(nums []int) []int {
	size := len(nums)
	ans := make([]int, size)

	left, right := 0, size-1
	index := size - 1

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]

		if leftSq > rightSq {
			ans[index] = leftSq
			left++
		} else {
			ans[index] = rightSq
			right--
		}

		index--
	}

	return ans
}
