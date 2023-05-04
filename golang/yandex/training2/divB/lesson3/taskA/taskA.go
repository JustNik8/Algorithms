package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var empty struct{}

// https://contest.yandex.ru/contest/28964/problems/
func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	strNums := strings.Split(strings.TrimSpace(str), " ")

	nums := make(map[int]struct{})
	for i := 0; i < len(strNums); i++ {
		num, _ := strconv.Atoi(strNums[i])
		nums[num] = empty
	}

	str, _ = reader.ReadString('\n')
	strNums = strings.Split(strings.TrimSpace(str), " ")

	cnt := 0
	for i := 0; i < len(strNums); i++ {
		num, _ := strconv.Atoi(strNums[i])
		_, found := nums[num]
		if found {
			cnt++
		}
	}
	
	fmt.Print(cnt)
}
