package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c uint64
	fmt.Fscan(reader, &a, &b, &c)

	r := uint64(math.Pow10(16))
	params := checkParams{
		a: a,
		b: b,
		c: c,
	}
	ans := leftBinSearch(0, r, check, params)
	fmt.Println(ans)
}

func leftBinSearch(l uint64, r uint64, check func(uint64, checkParams) bool, params checkParams) uint64 {
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

func check(m uint64, params checkParams) bool {
	sum2 := params.a * 2
	sum3 := params.b * 3
	sum4 := params.c * 4
	sum5 := m * 5

	totalSum := sum2 + sum3 + sum4 + sum5
	cnt := params.a + params.b + params.c + m

	intPart := totalSum / cnt
	remains := totalSum % cnt

	if intPart >= 4 {
		return true
	} else if intPart >= 3 {
		half := cnt / 2
		return cnt-remains <= half
	} else {
		return false
	}
}

type checkParams struct {
	a uint64
	b uint64
	c uint64
}
