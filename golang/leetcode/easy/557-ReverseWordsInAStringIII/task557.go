package main

import "strings"

func reverseWords(s string) string {
	ans := make([]string, 0)
	for _, v := range strings.Split(s, " ") {
		temp := []byte(v)

		l, r := 0, len(temp)-1
		for l < r {
			temp[l], temp[r] = temp[r], temp[l]
			l++
			r--
		}

		ans = append(ans, string(temp))
	}

	return strings.Join(ans, " ")
}
