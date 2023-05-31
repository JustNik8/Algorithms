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

func main() {
	reader := bufio.NewReader(os.Stdin)

	// n - количество студентов (вершины)
	// m - количество пар студентов, которые обменивались записками (ребра)
	var n, m int
	fmt.Fscan(reader, &n, &m)

	graph := make([]set, n+1)
	colors := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		graph[i] = make(set)
	}

	for i := 0; i < m; i++ {
		var s1, s2 int
		fmt.Fscan(reader, &s1, &s2)

		graph[s1][s2] = empty{}
		graph[s2][s1] = empty{}
	}

	for i := 1; i < n+1; i++ {
		if colors[i] == white {
			paint(graph, &colors, i, gray)
			if !canPaint {
				break
			}
		}
	}

	if canPaint {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

var canPaint = true

func paint(graph []set, colors *[]int, current int, color int) {
	if !canPaint {
		return
	}
	(*colors)[current] = color

	for neig, _ := range graph[current] {
		if (*colors)[neig] == white {
			paint(graph, colors, neig, chooseColor(color))
		} else if color == (*colors)[neig] {
			canPaint = false
			return
		}
	}

}

func chooseColor(color int) int {
	return 3 - color
}
