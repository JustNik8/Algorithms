package main

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			ans := checkPalindrome(s, l+1, r) || checkPalindrome(s, l, r-1)
			return ans
		}
		l++
		r--
	}

	return true
}

func checkPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
