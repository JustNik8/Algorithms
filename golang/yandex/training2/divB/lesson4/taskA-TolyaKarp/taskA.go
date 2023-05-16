package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://contest.yandex.ru/contest/28970/problems/A/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	parcels := make(map[int64]int64)
	for i := 0; i < n; i++ {
		var color, number int64
		fmt.Fscan(reader, &color, &number)
		parcels[color] += number
	}

	sortedKeys := make([]int64, 0)
	for color := range parcels {
		sortedKeys = append(sortedKeys, color)
	}

	sort.Slice(sortedKeys, func(i, j int) bool {
		return sortedKeys[i] < sortedKeys[j]
	})
	for _, color := range sortedKeys {
		number := parcels[color]
		fmt.Printf("%d %d\n", color, number)
	}
}
