package main

import "fmt"

func main() {
	mathings := map[uint8]uint8{
		'(': ')',
	}

	fmt.Println(mathings)
}

func longestPalindrome(s string) string {
	maxLeft, maxRight := 0, 0
	size := len(s)

	for i := 0; i < size; i++ {
		left, right := i, i

		for left >= 0 && right < size {
			if s[left] != s[right] {
				break
			}

			currentLen := right - left + 1

			if currentLen > maxRight-maxLeft+1 {
				maxRight, maxLeft = right, left
			}

			left--
			right++
		}

		left, right = i, i+1
		for left >= 0 && right < size {
			if s[left] != s[right] {
				break
			}

			currentLen := right - left + 1

			if currentLen > maxRight-maxLeft+1 {
				maxRight, maxLeft = right, left
			}

			left--
			right++
		}
	}

	return s[maxLeft : maxRight+1]
}
