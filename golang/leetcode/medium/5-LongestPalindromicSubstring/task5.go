package main

func main() {
}

// Более красивое решение
func longestPalindromeClear(s string) string {
	maxLeft, maxRight := 0, 0
	size := len(s)

	for i := 0; i < size; i++ {
		maxLeft, maxRight = calculate(s, i, i, maxLeft, maxRight)
		maxLeft, maxRight = calculate(s, i, i+1, maxLeft, maxRight)
	}

	return s[maxLeft : maxRight+1]
}

func calculate(s string, left, right, maxLeft, maxRight int) (int, int) {
	size := len(s)
	for left >= 0 && right < size {
		if s[left] != s[right] {
			break
		}

		if right-left+1 > maxRight-maxLeft+1 {
			maxLeft, maxRight = left, right
		}

		left--
		right++
	}

	return maxLeft, maxRight
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
