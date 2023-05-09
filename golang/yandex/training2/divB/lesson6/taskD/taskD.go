package main

import (
	"bufio"
	"fmt"
	"os"
)

type params struct {
	dimaCut    uint64
	dimaChill  uint64
	fedyaCut   uint64
	fedyaChill uint64
	trees      uint64
}

// https://contest.yandex.ru/contest/29188/problems/D/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var dimaCut, dimaChill, fedyaCut, fedyaChill, trees uint64

	fmt.Fscan(reader, &dimaCut, &dimaChill, &fedyaCut, &fedyaChill, &trees)

	p := params{
		dimaCut:    dimaCut,
		dimaChill:  dimaChill,
		fedyaCut:   fedyaCut,
		fedyaChill: fedyaChill,
		trees:      trees,
	}

	r := trees * 2 / (dimaCut + 1)
	days := leftBinSearch(0, r, check, p)
	fmt.Println(days)
}

func check(m uint64, p params) bool {
	dima := m*p.dimaCut - (m/p.dimaChill)*p.dimaCut
	fedya := m*p.fedyaCut - (m/p.fedyaChill)*p.fedyaCut

	return dima+fedya >= p.trees
}

func leftBinSearch(l, r uint64, check func(uint64, params) bool, p params) uint64 {
	for l < r {
		m := (l + r) / 2
		if check(m, p) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
