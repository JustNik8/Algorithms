package main

func findMiddleIndex(nums []int) int {
	pref := make([]int, len(nums)+1)

	for i := 1; i < len(pref); i++ {
		numIdx := i - 1
		pref[i] = pref[i-1] + nums[numIdx]
	}

	n := len(pref)
	for i := 1; i < len(pref); i++ {
		if pref[i-1]-pref[0] == pref[n-1]-pref[i] {
			return i - 1
		}
	}

	return -1
}

func findMiddleIndex2(nums []int) int {
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}

	leftSum := 0
	for i, v := range nums {
		totalSum -= v

		if leftSum == totalSum {
			return i
		}

		leftSum += v
	}

	return -1
}
