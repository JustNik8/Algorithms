package main

func main() {

}

func countBits(n int) []int {
	dp := make([]int, n+1)

	for num := 0; num <= n; num++ {
		offset := num >> 1
		dp[num] = dp[offset] + (num & 1)
	}

	return dp
}
