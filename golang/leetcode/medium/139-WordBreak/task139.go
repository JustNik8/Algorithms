package main

func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	dp := make([]bool, size+1)

	dp[size] = true

	for i := size - 1; i >= 0; i-- {
		for _, w := range wordDict {
			wSize := len(w)
			if i+wSize <= size && s[i:i+wSize] == w {
				dp[i] = dp[i+wSize]
			}
			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}
