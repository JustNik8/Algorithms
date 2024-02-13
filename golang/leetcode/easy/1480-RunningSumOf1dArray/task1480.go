package main

func runningSum(nums []int) []int {
	pref := make([]int, len(nums)+1)

	for i := 1; i < len(pref); i++ {
		numIdx := i - 1
		pref[i] = nums[numIdx] + pref[i-1]
	}

	return pref[1:]
}
