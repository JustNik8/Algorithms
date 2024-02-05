package main

/**
 * @param nums: a list of integer
 * @return: return a integer, denote  the maximum number of consecutive 1s
 */
func FindMaxConsecutiveOnes(nums []int) int {
	zeros := 0
	maxOnes := 0

	l, r := 0, 0

	for r < len(nums) {
		if nums[r] == 0 {
			zeros++
		}

		for zeros > 1 {
			if nums[l] == 0 {
				zeros--
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
