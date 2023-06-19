package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for i := 0; i < len(s); i++ {
		sLetter := s[i] - 'a'
		tLetter := t[i] - 'a'

		freq[sLetter]++
		freq[tLetter]--
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] != 0 {
			return false
		}
	}

	return true
}
