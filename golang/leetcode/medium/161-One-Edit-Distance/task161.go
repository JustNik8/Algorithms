package main

func LengthOfLongestSubstringTwoDistinct(s string) int {
	counts := make(map[byte]int)
	diffCnt := 0
	maxLen := 0

	left, right := 0, 0
	for right < len(s) {
		if diffCnt <= 2 {
			rightLetter := s[right]
			cnt := counts[rightLetter]

			if cnt == 0 {
				diffCnt++
			}
			counts[rightLetter] += 1

			if diffCnt <= 2 {
				maxLen = max(maxLen, right-left+1)
			}

			right++
		} else {
			leftLetter := s[left]
			counts[leftLetter] -= 1

			if counts[leftLetter] == 0 {
				diffCnt--
			}

			left++
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
