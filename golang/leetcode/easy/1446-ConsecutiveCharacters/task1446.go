package main

func maxPower(s string) int {

	maxPower := 1
	power := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			power++
		} else {
			maxPower = max(maxPower, power)
			power = 1
		}
	}

	maxPower = max(maxPower, power)
	return maxPower
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
