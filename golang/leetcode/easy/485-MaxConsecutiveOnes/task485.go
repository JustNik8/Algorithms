package main

func findMaxConsecutiveOnes(nums []int) int {
	maxOnes := 0
	ones := 0

	for _, num := range nums {
		if num == 1 {
			ones++
		} else {
			maxOnes = max(maxOnes, ones)
			ones = 0
		}
	}
	maxOnes = max(maxOnes, ones)

	return maxOnes
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
