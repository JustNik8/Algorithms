package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type cond struct {
	power int
	price int
}

// https://contest.yandex.ru/contest/27794/problems/G/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	classes := make([]int, n)
	for i := 0; i < n; i++ {
		var class int
		fmt.Fscan(reader, &class)
		classes[i] = class
	}
	sort.Ints(classes)

	var m int
	fmt.Fscan(reader, &m)
	conds := make([]cond, m)
	for i := 0; i < m; i++ {
		var power, price int
		fmt.Fscan(reader, &power, &price)
		conds[i] = cond{
			power: power,
			price: price,
		}
	}

	sort.SliceStable(conds, func(i, j int) bool {
		return conds[i].power < conds[j].power
	})

	var classPointer = 0
	var condPointer = findPointerForBestPrice(conds, 0)
	totalSum := 0
	for classPointer < n {
		if conds[condPointer].power >= classes[classPointer] {
			totalSum += conds[condPointer].price
			classPointer++
		} else {
			condPointer = findPointerForBestPrice(conds, condPointer+1)
		}
	}
	fmt.Println(totalSum)
}

func findPointerForBestPrice(conds []cond, start int) int {
	size := len(conds)
	bestPrice := conds[start].price
	pointer := start
	for i := start + 1; i < size; i++ {
		if conds[i].price <= bestPrice {
			bestPrice = conds[i].price
			pointer = i
		}
	}
	return pointer
}
