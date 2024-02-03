package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	p1 := 0
	p2 := 0

	for p2 < len(t) {
		if s[p1] == t[p2] {
			p1++
		}
		if p1 == len(s) {
			return true
		}

		p2++
	}

	return false
}
