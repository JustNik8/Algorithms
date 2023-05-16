package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/28738/problems/C/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var str string

	fmt.Fscan(reader, &str)

	// cognitive
	cost := 0
	size := len(str)
	for i := 0; i < size/2; i++ {
		if str[i] != str[size-i-1] {
			cost++
		}
	}
	fmt.Println(cost)
}
