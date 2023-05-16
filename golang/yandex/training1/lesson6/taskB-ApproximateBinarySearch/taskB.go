package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://contest.yandex.ru/contest/27844/problems/B/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(reader, &num)
		nums[i] = num
	}

	for i := 0; i < k; i++ {
		var num int
		fmt.Fscan(reader, &num)

		ans := rightBinSearch(
			0,
			n-1,
			check,
			checkParams{seq: nums, number: num},
		)
		left := nums[ans]
		if ans < n-1 {
			right := nums[ans+1]
			if math.Abs(float64(num-right)) < math.Abs(float64(num-left)) {
				fmt.Println(right)
			} else {
				fmt.Println(left)
			}
		} else {
			fmt.Println(left)
		}
	}
}

func check(m int, params checkParams) bool {
	return params.seq[m] <= params.number
}

func rightBinSearch(l int, r int, check func(int, checkParams) bool, checkParams checkParams) int {
	for l < r {
		m := (l + r + 1) / 2
		if check(m, checkParams) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

type checkParams struct {
	seq    []int
	number int
}
