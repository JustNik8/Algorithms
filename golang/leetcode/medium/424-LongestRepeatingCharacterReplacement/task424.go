package main

func main() {
	characterReplacement("BAAAB", 2)
}

func characterReplacement(s string, k int) int {
	size := len(s)
	counts := make(map[byte]int)

	left, right := 0, 0
	maxFreq := 0
	ans := 0

	for right < size {
		char := s[right]
		counts[char] += 1

		maxFreq = max(maxFreq, counts[char])

		if (right-left+1)-maxFreq > k {
			counts[s[left]] -= 1
			left++
		}

		ans = max(ans, right-left+1)
		right++
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
