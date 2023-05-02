package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/28738/problems/D/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var l, k int
	fmt.Fscan(reader, &l, &k)

	legs := make([]int, k)
	for i := 0; i < k; i++ {
		var leg int
		fmt.Fscan(reader, &leg)
		legs[i] = leg
	}

	middle := l / 2

	pointer := 0
	for {
		leg := legs[pointer]
		if leg == middle && l%2 == 1 {
			fmt.Print(leg)
			return
		} else if leg < middle {
			pointer++
		} else {
			break
		}
	}

	fmt.Printf("%d %d", legs[pointer-1], legs[pointer])

}
