package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// https://contest.yandex.ru/contest/45468/problems/29/
// WA 9 :(
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	costs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &costs[i])
	}
	if n == 0 {
		fmt.Println(0)
		fmt.Println(0, 0)
		return
	} else if n == 1 {
		fmt.Println(costs[0])
		fmt.Println(1, 0)
		return
	}

	dp := make([][]int, n+1)
	transfers := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 100_000
		}
		transfers[i] = make([]int, n)
		for j := 0; j < n; j++ {
			transfers[i][j] = 0
		}
	}

	if costs[0] <= 100 {
		dp[0][0] = costs[0]
	} else {
		dp[1][0] = costs[0]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if costs[i] == 0 {
				dp[j][i] = min(dp[j][i], dp[j][i-1]+costs[i])
			} else if costs[i] <= 100 {
				dp[j][i] = min3(dp[j][i], dp[j][i-1]+costs[i], dp[j+1][i-1])
			} else {
				dp[j][i] = min(dp[j][i], dp[j+1][i-1])
				dp[j+1][i] = min(dp[j+1][i], dp[j][i-1]+costs[i])
			}

			if dp[j][i] == dp[j][i-1]+costs[i] {
				transfers[j][i] = j
			} else if dp[j][i] == dp[j+1][i-1] {
				transfers[j][i] = j + 1
			}
		}
	}

	print(dp)
	print(transfers)

	minSum := dp[0][n-1]
	minSumIndex := 0
	for i := 1; i < n; i++ {
		if dp[i][n-1] <= minSum {
			minSum = dp[i][n-1]
			minSumIndex = i
		}
	}

	index := minSumIndex
	usedCoupons := 0
	days := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		next := transfers[index][i]
		if next-index > 0 {
			usedCoupons++
			days = append(days, i+1)
		}
		index = next
	}

	fmt.Println(minSum)
	fmt.Println(minSumIndex, usedCoupons)

	sort.Ints(days)
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func min3(a, b, c int) int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	return arr[0]
}

func print(dp [][]int) {
	fmt.Println()
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			fmt.Print(dp[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}
