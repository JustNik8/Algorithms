package main

func partitionLabels(s string) []int {
	ans := make([]int, 0)

	maxPositions := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		maxPositions[s[i]] = i
	}

	l, r := 0, maxPositions[s[0]]
	start := 0

	for {
		for l != r {
			l++
			r = max(r, maxPositions[s[l]])
		}

		ans = append(ans, r-start+1)
		l++
		if l == len(s) {
			break
		}
		start = l
		r = max(r, maxPositions[s[l]])
	}

	return ans
}
