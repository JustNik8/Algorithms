package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/28730/problems/D/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var studentCnt int

	fmt.Fscan(reader, &studentCnt)

	median := studentCnt / 2

	ans := 0
	for i := 0; i <= median; i++ {
		fmt.Fscan(reader, &ans)
	}

	fmt.Println(ans)
}
