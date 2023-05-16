package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/28964/problems/C/
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

	numsCnt := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		numsCnt[num] += 1
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		cnt := numsCnt[num]
		if cnt == 1 {
			fmt.Printf("%d ", num)
		}
	}
}
