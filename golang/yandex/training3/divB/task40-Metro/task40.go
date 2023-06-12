package main

import (
	"bufio"
	"fmt"
	"os"
)

type empty struct{}

type set map[int]empty

var lines = make(map[int]set)

// https://contest.yandex.ru/contest/45468/problems/40/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var stationCnt, linesCnt int
	fmt.Fscan(reader, &stationCnt, &linesCnt)

	graph := make([][]int, linesCnt+1)

	for line := 1; line <= linesCnt; line++ {
		var cnt int
		fmt.Fscan(reader, &cnt)

		lines[line] = make(set)
		for i := 0; i < cnt; i++ {
			var station int
			fmt.Fscan(reader, &station)

			lines[line][station] = empty{}
		}
	}

	for i := 1; i <= linesCnt; i++ {
		line := lines[i]
		for j := 1; j <= linesCnt; j++ {
			if i == j {
				continue
			}

			anotherLine := lines[j]

			for station, _ := range line {
				if _, ok := anotherLine[station]; ok {
					graph[i] = append(graph[i], j)
					break
				}
			}
		}
	}

	var start, target int
	fmt.Fscan(reader, &start, &target)

	fmt.Println(graph)
}
