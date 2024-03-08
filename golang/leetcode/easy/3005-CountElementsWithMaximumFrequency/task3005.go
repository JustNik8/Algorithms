package main

func maxFrequencyElements(nums []int) int {
	freqs := make(map[int]int)

	maxFreq := 0

	for _, num := range nums {
		freqs[num] += 1

		maxFreq = max(maxFreq, freqs[num])
	}

	ans := 0
	for _, freq := range freqs {
		if freq == maxFreq {
			ans += freq
		}
	}

	return ans
}
