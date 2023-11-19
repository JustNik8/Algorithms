package main

func canJump(nums []int) bool {
	size := len(nums)
	// Pointer
	p := size - 1

	for i := size - 2; i >= 0; i-- {
		if i+nums[i] >= p {
			p = i
		}
	}

	return p == 0
}
