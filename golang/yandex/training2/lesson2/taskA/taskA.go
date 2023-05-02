package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/28738/problems/
func main() {
	reader := bufio.NewReader(os.Stdin)

	max := 0
	maxCnt := 0
	for {
		var num int
		fmt.Fscan(reader, &num)
		if num == 0 {
			break
		}

		if num > max {
			max = num
			maxCnt = 1
		} else if num == max {
			maxCnt++
		}
	}

	fmt.Println(maxCnt)
}
