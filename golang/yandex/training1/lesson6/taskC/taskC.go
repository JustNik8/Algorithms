package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://contest.yandex.ru/contest/27844/problems/C/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var w, h, n uint64
	fmt.Fscan(reader, &w, &h, &n)

	l := uint64(0)
	r := uint64(math.Pow10(18))
	for l < r {
		m := (l + r) / 2
		if check(m, w, h, n) {
			r = m
		} else {
			l = m + 1
		}
	}

	fmt.Print(l)
}

func check(m uint64, w uint64, h uint64, n uint64) bool {
	mw := m / w
	mh := m / h
	return mw*mh >= n
}
