package main

import (
	"bufio"
	"fmt"
	"os"
)

type empty struct{}

type set map[int]empty

// Хранит компоненты связности
var comp = make([][]int, 0)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var vertexCnt, edgeCnt int
	fmt.Fscan(reader, &vertexCnt, &edgeCnt)

	visited := make([]bool, vertexCnt+1)
	graph := make([]set, vertexCnt+1)

	for i := 0; i < vertexCnt+1; i++ {
		graph[i] = make(set)
	}

	for i := 0; i < edgeCnt; i++ {
		var v1, v2 int
		fmt.Fscan(reader, &v1, &v2)

		graph[v1][v2] = empty{}
		graph[v2][v1] = empty{}
	}

	compIndex := -1

	for i := 1; i < vertexCnt+1; i++ {
		if !(visited)[i] {
			compIndex++
			comp = append(comp, make([]int, 0))
			dfs(graph, &visited, i, compIndex)
		}
	}

	fmt.Println(len(comp))
	for i := 0; i < len(comp); i++ {
		currentComp := comp[i]
		fmt.Println(len(currentComp))
		for j := 0; j < len(currentComp); j++ {
			fmt.Print(currentComp[j], " ")
		}
		fmt.Println()
	}

}

func dfs(graph []set, visited *[]bool, current int, compIndex int) {
	(*visited)[current] = true

	comp[compIndex] = append(comp[compIndex], current)

	for neig, _ := range graph[current] {
		if !(*visited)[neig] {
			dfs(graph, visited, neig, compIndex)
		}
	}
}
