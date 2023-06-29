package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	if len(t) > len(s) || s == "" {
		return ""
	}

	tChars := make(map[byte]int)
	sChars := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		char := t[i]
		tChars[char] += 1
	}

	needMatches := len(tChars)
	matches := 0
	window := ""

	left := 0
	for right := 0; right < len(s); right++ {
		char := s[right]

		tCnt, exists := tChars[char]

		if exists {
			sChars[char] += 1
			sCnt := sChars[char]

			if tCnt == sCnt {
				matches++
			}
		}

		for matches == needMatches {
			currLen := right - left + 1
			if currLen < len(window) || window == "" {
				window = s[left : right+1]
			}

			leftChar := s[left]
			_, okInT := tChars[leftChar]

			if okInT {
				sChars[leftChar] -= 1

				if sChars[leftChar] < tChars[leftChar] {
					matches--
				}
			}

			left++
		}
	}

	return window
}
