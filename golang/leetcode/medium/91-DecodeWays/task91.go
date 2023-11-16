package main

func numDecodings(s string) int {
	size := len(s)
	dp := make([]int, size+1)
	dp[0] = 1

	for j := 0; j < size; j++ {
		i := j + 1
		if s[j] == '0' {
			dp[i] = 0
		} else {
			dp[i] = dp[i-1]
		}

		if j-1 >= 0 && validate(s[j-1:j+1]) {
			dp[i] += dp[i-2]
		}
	}

	return dp[size]
}

func validate(num string) bool {
	if num[0] == '1' {
		return true
	}

	return num[0] == '2' && num[1] <= '6'
}
