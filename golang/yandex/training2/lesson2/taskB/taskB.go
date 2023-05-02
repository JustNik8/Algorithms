package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	office = 0
	house  = 1
	shop   = 2
	size   = 10
)

// https://contest.yandex.ru/contest/28738/problems/B/
func main() {
	reader := bufio.NewReader(os.Stdin)

	buildings := make([]int, size)
	for i := 0; i < size; i++ {
		var building int
		fmt.Fscan(reader, &building)
		buildings[i] = building
	}

	maxDestShop := 0
	for i := 0; i < size; i++ {
		if buildings[i] == house {
			destShop := findClosestShop(i, buildings)
			if destShop > maxDestShop {
				maxDestShop = destShop
			}
		}
	}

	fmt.Println(maxDestShop)
}

func findClosestShop(index int, buildings []int) int {
	closestShop := 99
	shopIndex := index

	for shopIndex < size && buildings[shopIndex] != shop {
		shopIndex++
	}
	if shopIndex < size {
		closestShop = shopIndex - index
	}

	shopIndex = index
	for shopIndex >= 0 && buildings[shopIndex] != shop {
		shopIndex--
	}
	if shopIndex >= 0 && index-shopIndex < closestShop {
		closestShop = index - shopIndex
	}

	return closestShop
}
