package main

import (
	"bufio"
	"fmt"
	"os"
)

var dists []int

// https://contest.yandex.ru/contest/45468/problems/36/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	graph := make([][]int, n+1)
	dists = make([]int, n+1)

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
	bfs(graph, start, start)

	fmt.Println(dists[target])
}

func bfs(graph [][]int, current, prev int) {

	for i := 1; i < len(graph); i++ {
		if graph[current][i] == 1 && i != prev {
			if dists[i] == -1 || dists[current]+1 < dists[i] {
				dists[i] = dists[current] + 1
				bfs(graph, i, current)
			}

		}
	}

}
