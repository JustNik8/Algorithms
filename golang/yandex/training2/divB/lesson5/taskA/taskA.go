package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/29075/problems/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(reader, &n, &q)

	pref := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		var num int
		fmt.Fscan(reader, &num)
		pref[i] = pref[i-1] + num
	}

	for i := 0; i < q; i++ {
		var left, right int
		fmt.Fscan(reader, &left, &right)

		fmt.Println(pref[right] - pref[left-1])
	}
}
