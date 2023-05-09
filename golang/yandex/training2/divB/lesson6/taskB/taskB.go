package main

import (
	"bufio"
	"fmt"
	"os"
)

type params struct {
	nums []int
	num  int
}

// https://contest.yandex.ru/contest/29188/problems/B/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(reader, &num)
		nums[i] = num
	}

	var m int
	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var num int
		fmt.Fscan(reader, &num)
		p := params{nums: nums, num: num}
		left := leftBinSearch(0, n-1, checkNum, p)
		if left == -1 {
			fmt.Println(0, 0)
		} else {
			right := rightBinSearch(left-1, n-1, isEqual, p)
			fmt.Println(left+1, right+1)
		}

	}
}

func checkNum(m int, p params) bool {
	return p.nums[m] >= p.num
}

func isEqual(m int, p params) bool {
	return p.nums[m] == p.num
}

func leftBinSearch(l, r int, check func(int, params) bool, p params) int {
	for l < r {
		m := (l + r) / 2
		if check(m, p) {
			r = m
		} else {
			l = m + 1
		}
	}
	if p.nums[l] != p.num {
		return -1
	}
	return l
}

func rightBinSearch(l, r int, check func(int, params) bool, p params) int {
	for l < r {
		m := (l + r + 1) / 2
		if check(m, p) {
			l = m
		} else {
			r = m - 1
		}
	}
	if p.nums[l] != p.num {
		return -1
	}
	return l
}
