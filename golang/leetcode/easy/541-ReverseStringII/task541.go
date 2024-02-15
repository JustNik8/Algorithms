package main

func reverseStr(s string, k int) string {
	ans := []byte(s)

	for i := 0; i < len(s); i += 2 * k {
		l, r := i, i+k-1

		if r >= len(s) {
			r = len(s) - 1
		}

		for l < r {
			ans[l], ans[r] = ans[r], ans[l]
			l++
			r--
		}
	}

	return string(ans)
}
