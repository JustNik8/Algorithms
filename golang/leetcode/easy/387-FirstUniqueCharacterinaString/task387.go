package main

func firstUniqChar(s string) int {
	chars := [26]int{}

	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if chars[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
