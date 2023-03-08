package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/27794/problems/D/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, r int
	fmt.Fscan(reader, &n, &r)

	dists := make([]int, n)
	for i := 0; i < n; i++ {
		var dist int
		fmt.Fscan(reader, &dist)
		dists[i] = dist
	}

	left := 0
	right := 1

	cnt := 0
	for right < n {
		if dists[right]-dists[left] <= r {
			right++
		} else {
			if dists[right]-dists[left] > r {
				cnt += n - right
			}
			left++
		}
	}

	fmt.Println(cnt)
}
