package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/27844/problems/D/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, a, b, w, h uint64

	fmt.Fscan(reader, &n, &a, &b, &w, &h)
	params := checkParams{
		a: a,
		b: b,
		w: w,
		h: h,
		n: n,
	}
	ans := rightBinSearch(0, w+h, check, params)
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
	aNew := params.a + 2*m
	bNew := params.b + 2*m

	hCnt1 := params.h / aNew
	wCnt1 := params.w / bNew

	hCnt2 := params.w / aNew
	wCnt2 := params.h / bNew

	return hCnt1*wCnt1 >= params.n || hCnt2*wCnt2 >= params.n
}

type checkParams struct {
	a uint64
	b uint64
	w uint64
	h uint64
	n uint64
}
