package main

func main() {

}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/?envType=list&envId=oceyii8d
func lengthOfLongestSubstring(s string) int {
	chars := make(map[uint8]int)

	left, right := 0, 0

	size := len(s)
	maxLen := 0
	currentLen := 0

	for right < size {
		rightChar := s[right]
		cnt := chars[rightChar]

		if cnt >= 1 {
			if currentLen > maxLen {
				maxLen = currentLen
			}

			currentLen--

			leftChar := s[left]
			chars[leftChar] -= 1

			left++
		} else {
			chars[rightChar] += 1
			currentLen++

			right++
		}
	}

	if currentLen > maxLen {
		maxLen = currentLen
	}

	return maxLen
}
