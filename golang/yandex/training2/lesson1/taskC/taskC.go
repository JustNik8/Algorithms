package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/28730/problems/C/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var x, y, z int

	fmt.Fscan(reader, &x, &y, &z)

	var res = 0
	if (x > 12 || y > 12) || x == y {
		res = 1
	}

	fmt.Println(res)
}
