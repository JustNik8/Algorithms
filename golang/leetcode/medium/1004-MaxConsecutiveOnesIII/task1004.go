package main

func longestOnes(nums []int, k int) int {
	maxOnes := 0
	flips := 0
	l, r := 0, 0

	for r < len(nums) {
		if nums[r] == 0 {
			flips++
		}

		for flips > k {
			if nums[l] == 0 {
				flips--
			}
			l++
		}

		maxOnes = max(maxOnes, r-l+1)
		r++
	}

	return maxOnes
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
