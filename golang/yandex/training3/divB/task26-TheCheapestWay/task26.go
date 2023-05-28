package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/26/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int

	fmt.Fscan(reader, &n, &m)

	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &table[i][j])
		}
	}

	l := 0
	r := 50_000
	for l < r {
		m = (l + r) / 2
		if checkReach(&table, m) {
			r = m
		} else {
			l = m + 1
		}
	}

	fmt.Println(l)
}

func checkReach(table *[][]int, food int) bool {
	n := len(*table)
	m := len((*table)[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = (*table)[0][0]
	for i := 1; i < m; i++ {
		dp[0][i] += dp[0][i-1] + (*table)[0][i]
	}

	for i := 1; i < n; i++ {
		dp[i][0] += dp[i-1][0] + (*table)[i][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			left := dp[i-1][j] + (*table)[i][j]
			right := dp[i][j-1] + (*table)[i][j]
			dp[i][j] = int(math.Min(float64(left), float64(right)))
		}
	}

	return food-dp[n-1][m-1] >= 0
}
