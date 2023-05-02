package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://contest.yandex.ru/contest/28730/problems/B/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var stationCnt, enter, exit int

	fmt.Fscan(reader, &stationCnt, &enter, &exit)

	straight := math.Abs(float64(enter)-float64(exit)) - 1

	minStation := math.Min(float64(enter), float64(exit))
	maxStation := math.Max(float64(enter), float64(exit))
	reverse := float64(stationCnt) - maxStation + minStation - 1.0

	res := int(math.Min(straight, reverse))
	fmt.Println(res)
}
