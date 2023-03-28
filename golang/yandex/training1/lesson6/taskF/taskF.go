package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, x, y uint64
	fmt.Fscan(reader, &n, &x, &y)

	r := uint64(math.Pow10(16))
	params := checkParams{
		n: n - 1,
		x: x,
		y: y,
	}
	ans := leftBinSearch(0, r, check, params)
	ans += uint64(math.Min(float64(x), float64(y)))
	fmt.Print(ans)
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
	maxX := m / params.x
	maxY := m / params.y
	return maxX+maxY >= params.n
}

type checkParams struct {
	n uint64
	x uint64
	y uint64
}
