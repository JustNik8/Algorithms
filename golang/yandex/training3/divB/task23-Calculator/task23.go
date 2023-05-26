package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/23/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	if n == 1 {
		fmt.Println(0)
		fmt.Println(1)
		return
	}

	size := int(math.Max(float64(n+1), float64(4)))

	dp := make([]uint32, size)
	prev := make([]uint32, size)

	dp[1], dp[2], dp[3] = 0, 1, 1
	prev[1], prev[2], prev[3] = 0, 1, 1

	for i := 2; i < n; i++ {
		cnt := dp[i] + 1
		num := i * 3
		if num <= n && (dp[num] == 0 || dp[num] > cnt) {
			dp[num] = cnt
			prev[num] = uint32(i)
		}

		num = i * 2
		if num <= n && (dp[num] == 0 || dp[num] > cnt) {
			dp[num] = cnt
			prev[num] = uint32(i)
		}

		num = i + 1
		if num <= n && (dp[num] == 0 || dp[num] > cnt) {
			dp[num] = cnt
			prev[num] = uint32(i)
		}
	}

	fmt.Println(dp[n])

	ansSize := dp[n] + 1
	ans := make([]uint32, ansSize)
	ans[ansSize-1] = uint32(n)

	ansIndex := ansSize - 2
	pointer := prev[n]
	for pointer != 0 {
		ans[ansIndex] = pointer
		pointer = prev[pointer]
		ansIndex--
	}

	for i := uint32(0); i < ansSize; i++ {
		fmt.Print(ans[i], " ")
	}
}
