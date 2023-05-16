package main

import (
	"bufio"
	"fmt"
	"os"
)

type empty struct{}
type byteSet map[byte]empty
type carNumbers struct {
	carNumber string
	matchCnt  int
}

// https://contest.yandex.ru/contest/28964/problems/E/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var m int
	fmt.Fscan(reader, &m)

	witness := make([]byteSet, m)
	for i := 0; i < m; i++ {
		set := byteSet{}
		witness[i] = set
		var str string
		fmt.Fscan(reader, &str)
		for j := 0; j < len(str); j++ {
			char := str[j]
			set[char] = empty{}
		}
	}

	var n int
	fmt.Fscan(reader, &n)
	carNums := make([]carNumbers, n)
	maxMatch := 0

	for i := 0; i < n; i++ {
		var carNumber string
		fmt.Fscan(reader, &carNumber)
		matchCnt := 0
		for _, set := range witness {
			copySet := byteSet{}
			for b := range set {
				copySet[b] = empty{}
			}

			for k := 0; k < len(carNumber); k++ {
				delete(copySet, carNumber[k])
			}

			if len(copySet) == 0 {
				matchCnt++
			}
		}

		carNums[i] = carNumbers{carNumber: carNumber, matchCnt: matchCnt}
		if matchCnt > maxMatch {
			maxMatch = matchCnt
		}
	}

	for _, carNum := range carNums {
		if carNum.matchCnt == maxMatch {
			fmt.Println(carNum.carNumber)
		}
	}
}
