package main

import (
	"bufio"
	"fmt"
	"os"
)

var hasCycle = false
var path = make([]int, 0)

const (
	white = 0
	gray  = 1
	black = 2
)

// https://contest.yandex.ru/contest/45468/problems/35/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	graph := make([][]int, n+1)
	colors := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		graph[i] = make([]int, n+1)
		for j := 1; j < n+1; j++ {
			fmt.Fscan(reader, &graph[i][j])
		}
	}

	for i := 1; i < n+1; i++ {
		if colors[i] == white {
			dfs(graph, &colors, i, 0)
		}
	}

	if hasCycle {
		fmt.Println("YES")
		size := len(path)
		cycleVertex := path[size-1]
		i := size - 2
		for path[i] != cycleVertex {
			i--
		}
		fmt.Println(size - i - 1)
		for ; i < size-1; i++ {
			fmt.Print(path[i], " ")
		}
	} else {
		fmt.Println("NO")
	}
}

func dfs(graph [][]int, colors *[]int, current int, prev int) {
	if hasCycle {
		return
	}
	(*colors)[current] = gray
	path = append(path, current)

	neigs := graph[current]

	for neig, hasEdge := range neigs {
		if hasEdge == 0 || neig == prev {
			continue
		}

		if (*colors)[neig] == gray && !hasCycle {
			hasCycle = true
			path = append(path, neig)
			return
		} else if (*colors)[neig] == white && !hasCycle {
			dfs(graph, colors, neig, current)
		}
	}

	if hasCycle {
		return
	}

	(*colors)[current] = black
	path = path[:len(path)-1]
}
