package main

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	ans := make([]int, 0)

	var sDict [26]int
	var pDict [26]int

	for i := 0; i < len(p); i++ {
		sDict[s[i]-'a']++
		pDict[p[i]-'a']++
	}

	if isEqual(sDict, pDict) {
		ans = append(ans, 0)
	}

	for l := 0; l+len(p) < len(s); l++ {
		r := l + len(p)

		sDict[s[l]-'a']--
		sDict[s[r]-'a']++

		if isEqual(sDict, pDict) {
			ans = append(ans, l+1)
		}
	}

	return ans
}

func isEqual(a [26]int, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
