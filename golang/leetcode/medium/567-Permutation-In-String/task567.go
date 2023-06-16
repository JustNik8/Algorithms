package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("adc", "dcda"))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	size := len(s1)
	right := size - 1
	left := 0

	s1Letters := make(map[byte]int)
	s2Letters := make(map[byte]int)

	for i := 0; i < size; i++ {
		letter := s1[i]
		s1Letters[letter] += 1
	}

	for i := 0; i <= right; i++ {
		letter := s2[i]
		s2Letters[letter] += 1
	}

	for right < len(s2) {
		if compare(s1Letters, s2Letters) {
			return true
		}

		fmt.Println(s1Letters)
		fmt.Println(s2Letters)

		leftLetter := s2[left]
		s2Letters[leftLetter] -= 1
		left++

		right++
		if right == len(s2) {
			break
		}
		rightLetter := s2[right]
		s2Letters[rightLetter] += 1
	}

	return false
}

func compare(s1Letters, s2Letters map[byte]int) bool {
	for letter, cnt := range s1Letters {
		s2LettersCnt := s2Letters[letter]

		if cnt != s2LettersCnt {
			return false
		}
	}

	return true
}
