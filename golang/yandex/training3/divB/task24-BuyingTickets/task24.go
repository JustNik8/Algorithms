package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://contest.yandex.ru/contest/45468/problems/24/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	dp := make([]int, n+3)
	a := make([]int, n+3)
	b := make([]int, n+3)
	c := make([]int, n+3)

	a[0], a[1], a[2] = 10000, 10000, 10000
	b[0], b[1], b[2] = 10000, 10000, 10000
	c[0], c[1], c[2] = 10000, 10000, 10000

	for i := 3; i < n+3; i++ {
		var aCur, bCur, cCur int
		fmt.Fscan(reader, &aCur, &bCur, &cCur)
		a[i], b[i], c[i] = aCur, bCur, cCur
		vars := make([]int, 3)
		vars[0] = dp[i-3] + c[i-2]
		vars[1] = dp[i-2] + b[i-1]
		vars[2] = dp[i-1] + a[i]

		sort.Ints(vars)
		dp[i] = vars[0]
	}

	fmt.Println(dp[n+2])
}
