package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://contest.yandex.ru/contest/28730/problems/E/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var d int
	var xPoint, yPoint int

	fmt.Fscan(reader, &d)
	fmt.Fscan(reader, &xPoint, &yPoint)

	ans := 0

	value := direct(xPoint, d)
	if xPoint >= 0 && yPoint >= 0 && yPoint <= value {
		ans = 0
	} else {
		ans = calcClosestPoint(xPoint, yPoint, d)
	}

	fmt.Println(ans)
}

func direct(x int, d int) int {
	return -x + d
}

func calcDest(x1 int, y1 int, x2 int, y2 int) float64 {
	formula := math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)
	return math.Sqrt(formula)
}

func calcClosestPoint(xPoint int, yPoint int, d int) int {
	var dests [3]float64
	dests[0] = calcDest(xPoint, yPoint, 0, 0)
	dests[1] = calcDest(xPoint, yPoint, d, 0)
	dests[2] = calcDest(xPoint, yPoint, 0, d)

	minIndex := 0
	for i := 1; i < 3; i++ {
		if dests[i] < dests[minIndex] {
			minIndex = i
		}
	}

	return minIndex + 1
}
