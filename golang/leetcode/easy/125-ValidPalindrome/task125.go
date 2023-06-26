package main

import "unicode"

func main() {

}

func isPalindrome(s string) bool {
	size := len(s)
	left, right := 0, size-1

	for left < right {
		leftChar := s[left]
		rightChar := s[right]

		if !isAlphabet(leftChar) {
			left++
		} else if !isAlphabet(rightChar) {
			right--
		} else {
			leftLC := unicode.ToLower(rune(leftChar))
			rightLC := unicode.ToLower(rune(rightChar))

			if leftLC != rightLC {
				return false
			}

			left++
			right--
		}
	}

	return true
}

func isAlphabet(char byte) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9')
}
