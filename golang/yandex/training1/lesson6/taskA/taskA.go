package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)

	firstSlice := make([]int, n)
	for i := 0; i < n; i++ {
		var elem int
		fmt.Fscan(reader, &elem)
		firstSlice[i] = elem
	}

	for i := 0; i < k; i++ {
		var elem int
		fmt.Fscan(reader, &elem)
		index := LeftBinSearch(0, n-1, check,
			checkParams{
				seq:    firstSlice,
				number: elem,
			},
		)
		if firstSlice[index] == elem {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func check(m int, params checkParams) bool {
	number := params.seq[m]
	if number < params.number {
		return false
	} else {
		return true
	}
}

func LeftBinSearch(l int, r int, check func(int, checkParams) bool, checkParams checkParams) int {
	for l < r {
		m := (l + r) / 2
		if check(m, checkParams) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

type checkParams struct {
	seq    []int
	number int
}
