package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://contest.yandex.ru/contest/27844/problems/H/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k uint64
	fmt.Fscan(reader, &n, &k)

	lens := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		var length uint64
		fmt.Fscan(reader, &length)
		lens[i] = length
	}

	params := checkParams{
		lens: lens,
		k:    k,
	}

	r := uint64(math.Pow10(8))
	ans := rightBinSearch(0, r, check, params)
	fmt.Print(ans)
}

func rightBinSearch(l uint64, r uint64, check func(uint64, checkParams) bool, params checkParams) uint64 {
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

func check(m uint64, params checkParams) bool {
	totalCnt := uint64(0)
	for i := 0; i < len(params.lens); i++ {
		totalCnt += params.lens[i] / m
	}

	return totalCnt >= params.k
}

type checkParams struct {
	lens []uint64
	k    uint64
}
