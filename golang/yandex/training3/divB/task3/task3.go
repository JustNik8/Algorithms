package main

import (
	"fmt"
	"sort"
)

// TL
func main() {
	var n int
	fmt.Scan(&n)

	diegoCardSet := make(map[int]bool)
	var collection []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		if _, ok := diegoCardSet[num]; !ok {
			diegoCardSet[num] = true
			collection = append(collection, num)
		}
	}

	sort.Ints(collection)

	var k int
	fmt.Scan(&k)
	buyers := make([]int, k)

	for i := range buyers {
		fmt.Scan(&buyers[i])
	}

	for _, card := range buyers {
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
