package main

import (
	"bufio"
	"fmt"
	"os"
)

// http://e-maxx.ru/algo/maximum_average_segment
// https://contest.yandex.ru/contest/29075/problems/B/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	sum := 0
	maxSum := -1000000000
	pref := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		var num int
		fmt.Fscan(reader, &num)

		sum = num + pref[i-1]
		pref[i] = sum

		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
		pref[i] = sum
	}

	fmt.Println(maxSum)
}
