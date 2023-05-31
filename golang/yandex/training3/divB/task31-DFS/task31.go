package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type empty struct{}

type set map[int]empty

// Хранит компоненту связности. Да не оч хорошо делать глобальную переме
var comp = make([]int, 0)

// https://contest.yandex.ru/contest/45468/problems/31/
func main() {
	reader := bufio.NewReader(os.Stdin)

	// n - количество вершин, m - ребра
	var n, m int
	fmt.Fscan(reader, &n, &m)

	visited := make([]bool, n+1)
	graph := make([]set, n+1)

	for i := 0; i < n+1; i++ {
		graph[i] = make(set)
	}

	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Fscan(reader, &v1, &v2)

		graph[v1][v2] = empty{}
		graph[v2][v1] = empty{}
	}

	dfs(graph, &visited, 1)

	sort.Ints(comp)
	fmt.Println(len(comp))
	for i := 0; i < len(comp); i++ {
		fmt.Print(comp[i], " ")
	}
}

func dfs(graph []set, visited *[]bool, current int) {
	(*visited)[current] = true
	comp = append(comp, current)
	for neig, _ := range graph[current] {
		if !(*visited)[neig] {
			dfs(graph, visited, neig)
		}
	}
}
