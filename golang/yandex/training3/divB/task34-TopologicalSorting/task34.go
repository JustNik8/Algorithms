package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	white = 0
	gray  = 1
	black = 2
)

type empty struct{}

type set map[int]empty

var sorted = make([]int, 0)
var hasCycle = false

// https://contest.yandex.ru/contest/45468/problems/34/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var vertexCnt, edgeCnt int
	fmt.Fscan(reader, &vertexCnt, &edgeCnt)

	graph := make([]set, vertexCnt+1)
	colors := make([]int, vertexCnt+1)

	for i := 0; i < vertexCnt+1; i++ {
		graph[i] = make(set)
	}

	for i := 0; i < edgeCnt; i++ {
		var v1, v2 int
		fmt.Fscan(reader, &v1, &v2)

		graph[v1][v2] = empty{}
	}

	for i := 1; i < vertexCnt+1; i++ {
		if colors[i] == white {
			dfs(graph, &colors, i)
		}
	}

	if hasCycle {
		fmt.Println(-1)
		return
	}
	for i := len(sorted) - 1; i >= 0; i-- {
		fmt.Print(sorted[i], " ")
	}
}

func dfs(graph []set, colors *[]int, current int) {
	if hasCycle {
		return
	}
	(*colors)[current] = gray

	for neig, _ := range graph[current] {
		if (*colors)[neig] == gray {
			hasCycle = true
			return
		} else if (*colors)[neig] == white {
			dfs(graph, colors, neig)
		}
	}

	(*colors)[current] = black
	sorted = append(sorted, current)
}
