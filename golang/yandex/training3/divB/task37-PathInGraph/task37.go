package main

import (
	"bufio"
	"fmt"
	"os"
)

var dists []int
var prevs []int

// https://contest.yandex.ru/contest/45468/problems/37/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	graph := make([][]int, n+1)
	dists = make([]int, n+1)
	prevs = make([]int, n+1)

	for i := 0; i < n+1; i++ {
		dists[i] = -1
	}

	for i := 1; i < n+1; i++ {
		graph[i] = make([]int, n+1)
		for j := 1; j < n+1; j++ {
			fmt.Fscan(reader, &graph[i][j])
		}
	}

	var start, target int
	fmt.Fscan(reader, &start, &target)

	dists[start] = 0
	prevs[start] = 0

	bfs(graph, start, start)
	fmt.Println(dists[target])

	if dists[target] != -1 && dists[target] != 0 {
		ans := make([]int, 0)
		ans = append(ans, target)

		v := prevs[target]
		for v != start {
			ans = append(ans, v)
			v = prevs[v]
		}
		ans = append(ans, start)

		for i := len(ans) - 1; i >= 0; i-- {
			fmt.Print(ans[i], " ")
		}
	}
}

func bfs(graph [][]int, current, prev int) {
	for i := 1; i < len(graph); i++ {
		if graph[current][i] == 1 && i != prev {
			if dists[i] == -1 || dists[current]+1 < dists[i] {
				dists[i] = dists[current] + 1
				prevs[i] = current
				bfs(graph, i, current)
			}
		}
	}
}
