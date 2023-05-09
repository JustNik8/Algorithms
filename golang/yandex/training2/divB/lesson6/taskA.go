package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type checkParams struct {
	arr      []int
	leftNum  int
	rightNum int
}

// https://contest.yandex.ru/contest/29188/problems/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(reader, &num)
		arr[i] = num
	}

	sort.Ints(arr)

	var k int
	fmt.Fscan(reader, &k)

	for i := 0; i < k; i++ {
		var left, right int
		fmt.Fscan(reader, &left, &right)

		params := checkParams{arr: arr, leftNum: left, rightNum: right}
		leftIndex := leftBinSearch(0, n-1, checkGreater, params)
		rightIndex := rightBinSearch(0, n-1, checkLess, params)
		if leftIndex != rightIndex {
			fmt.Printf("%d ", rightIndex-leftIndex+1)
		} else if arr[leftIndex] >= left && arr[leftIndex] <= right {
			fmt.Printf("%d ", 1)
		} else {
			fmt.Printf("%d ", 0)
		}
	}
}

func checkGreater(m int, params checkParams) bool {
	arr := params.arr
	leftNum := params.leftNum
	return arr[m] >= leftNum
}

func checkLess(m int, params checkParams) bool {
	arr := params.arr
	rightNum := params.rightNum
	return arr[m] <= rightNum
}

func leftBinSearch(l, r int, check func(int, checkParams) bool, params checkParams) int {
	for l < r {
		m := (l + r) / 2
		if check(m, params) {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func rightBinSearch(l, r int, check func(int, checkParams) bool, params checkParams) int {
	for l < r {
		m := (l + r + 1) / 2
		if check(m, params) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}
