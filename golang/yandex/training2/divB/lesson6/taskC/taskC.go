package main

import (
	"bufio"
	"fmt"
	"os"
)

type equationParams struct {
	a float64
	b float64
	c float64
	d float64
}

// https://contest.yandex.ru/contest/29188/problems/C/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c, d float64
	fmt.Fscan(reader, &a, &b, &c, &d)

	if d == 0 {
		fmt.Println(0.0)
		return
	}

	p := equationParams{a: a, b: b, c: c, d: d}
	r := 1.0
	for equation(r, p)*equation(-r, p) >= 0 {
		r *= 2
	}
	l := -r
	fmt.Println(solve(l, r, p))

}

func solve(l, r float64, p equationParams) float64 {
	var m float64
	for r-l > 0.000001 {
		m = (l + r) / 2
		leftY := equation(l, p)
		mediumY := equation(m, p)

		if leftY*mediumY <= 0 {
			r = m
		} else {
			l = m
		}
	}
	return m
}

func equation(x float64, ep equationParams) float64 {
	return ep.a*x*x*x + ep.b*x*x + ep.c*x + ep.d
}
