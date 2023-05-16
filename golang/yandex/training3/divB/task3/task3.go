package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type empty struct{}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	cardSet := make(map[int]empty)
	collection := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(reader, &num)
		if _, ok := cardSet[num]; !ok {
			cardSet[num] = empty{}
			collection = append(collection, num)
		}
	}

	sort.Ints(collection)

	var k int
	fmt.Fscan(reader, &k)

	for i := 0; i < k; i++ {
		var card int
		fmt.Fscan(reader, &card)
		fmt.Println(binarySearch(&collection, card))
	}
}

func binarySearch(nums *[]int, target int) int {
	left := 0
	right := len(*nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if (*nums)[mid] == target {
			return mid
		}

		if (*nums)[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
