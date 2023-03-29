package main

import (
	"bufio"
	"fmt"
	"os"
)

// Пока не решил. Здесь неправильный код!

func main() {
	reader := bufio.NewReader(os.Stdin)
	var squareHeight, squareWidth, tileCnt uint64
	fmt.Fscan(reader, &squareHeight, &squareWidth, &tileCnt)

	params := checkParams{
		squareHeight: squareHeight,
		squareWidth:  squareHeight,
		tileCnt:      tileCnt,
	}
	r := squareHeight / 2
	ans := leftBinSearch(0, r, check, params)
	fmt.Print(ans)
}

func leftBinSearch(l uint64, r uint64, check func(uint64, checkParams) bool, params checkParams) uint64 {
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
	if (params.squareWidth)-2*m <= 0 ||
		(params.squareHeight)-2*m <= 0 {
		return false
	}
	totalTiles := (params.squareWidth-2*m)*2 + (params.squareHeight-2*m)*2 + 4*m*m

	return totalTiles <= params.tileCnt
}

type checkParams struct {
	squareHeight uint64
	squareWidth  uint64
	tileCnt      uint64
}
