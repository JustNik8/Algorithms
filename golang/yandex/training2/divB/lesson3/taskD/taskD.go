package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type empty struct{}

// https://contest.yandex.ru/contest/28964/problems/D/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var maxGuessNum int

	fmt.Fscan(reader, &maxGuessNum)

	guessNums := make(map[int]empty)
	for i := 0; i < maxGuessNum; i++ {
		guessNums[i+1] = empty{}
	}

	str := readString(reader)
	for str != "HELP" {
		nums := getNums(str)

		str = readString(reader)
		if str == "NO" {
			for num := range nums {
				delete(guessNums, num)
			}
		} else if str == "YES" {
			for num := range guessNums {
				_, found := nums[num]
				if !found {
					delete(guessNums, num)
				}
			}
		}
	}

	sortKeys := make([]int, len(guessNums))
	i := 0
	for k := range guessNums {
		sortKeys[i] = k
		i++
	}

	sort.Ints(sortKeys)
	for _, k := range sortKeys {
		fmt.Printf("%d ", k)
	}
}

func readString(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)

}

func getNums(str string) map[int]empty {
	strNums := strings.Split(str, " ")
	nums := make(map[int]empty)
	for i := 0; i < len(strNums); i++ {
		num, _ := strconv.Atoi(strNums[i])
		nums[num] = empty{}
	}

	return nums
}
