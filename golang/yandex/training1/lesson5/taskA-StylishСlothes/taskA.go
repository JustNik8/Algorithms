package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://contest.yandex.ru/contest/27794/problems/A/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	shirts := make([]int, n)
	for i := 0; i < n; i++ {
		var shirt int
		fmt.Fscan(reader, &shirt)
		shirts[i] = shirt
	}

	var m int
	fmt.Fscan(reader, &m)

	pants := make([]int, m)
	for i := 0; i < m; i++ {
		var pant int
		fmt.Fscan(reader, &pant)
		pants[i] = pant
	}

	var pointer1 = 0
	var pointer2 = 0
	var chosenShirt = shirts[0]
	var chosenPants = pants[0]
	var min = math.Abs(float64(chosenShirt - chosenPants))

	for i := 0; i < n+m; i++ {
		diff := math.Abs(float64(shirts[pointer1] - pants[pointer2]))
		if diff < min {
			chosenShirt = shirts[pointer1]
			chosenPants = pants[pointer2]
			min = diff
		}

		if pointer1 < n-1 && shirts[pointer1] < pants[pointer2] {
			pointer1++
		} else if pointer2 < m-1 {
			pointer2++
		}
	}

	fmt.Printf("%d %d", chosenShirt, chosenPants)
}
