package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type empty struct{}

// https://contest.yandex.ru/contest/28964/problems/B/
func main() {
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strNums := strings.Split(str, " ")

	nums := make([]int, len(strNums))
	for i := 0; i < len(strNums); i++ {
		num, _ := strconv.Atoi(strNums[i])
		nums[i] = num
	}

	uniqueNums := make(map[int]empty)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		_, found := uniqueNums[num]
		if !found {
			uniqueNums[num] = empty{}
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
