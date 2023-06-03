package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	row int
	col int
}

var queue = make([]point, 0)
var dists [][]int

const movesCnt = 8

var di = []int{-2, -2, -1, 1, 2, 2, 1, -1}
var dj = []int{-1, 1, 2, 2, -1, 1, -2, -2}

// https://contest.yandex.ru/contest/45468/problems/38/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var rows, cols int
	fmt.Fscan(reader, &rows, &cols)

	var feederRow, feederCol int
	fmt.Fscan(reader, &feederRow, &feederCol)

	var fleasCnt int
	fmt.Fscan(reader, &fleasCnt)

	dists = make([][]int, rows+1)
	for i := 0; i < rows+1; i++ {
		dists[i] = make([]int, cols+1)
		for j := 0; j < cols+1; j++ {
			dists[i][j] = -1
		}
	}

	dists[feederRow][feederCol] = 0
	queue = append(queue, point{row: feederRow, col: feederCol})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 0; i < movesCnt; i++ {
			neigI := current.row + di[i]
			neigJ := current.col + dj[i]

			if neigI >= 1 && neigJ >= 1 && neigI <= rows && neigJ <= cols {
				if dists[neigI][neigJ] == -1 || dists[current.row][current.col]+1 < dists[neigI][neigJ] {
					dists[neigI][neigJ] = dists[current.row][current.col] + 1
					queue = append(queue, point{row: neigI, col: neigJ})
				}
			}
		}
	}

	sum := 0
	for i := 0; i < fleasCnt; i++ {
		var fleaRow, fleaCol int
		fmt.Fscan(reader, &fleaRow, &fleaCol)

		fleaDist := dists[fleaRow][fleaCol]

		if fleaDist == -1 {
			sum = -1
			break
		} else {
			sum += fleaDist
		}
	}

	fmt.Println(sum)

	// В задаче это не нужно выводить. Здесь это написано, для того, чтобы понимать как выглядит поле дистанций
	print(dists)
}

func print(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[1]); j++ {
			fmt.Print("\t", a[i][j])
		}
		fmt.Println()
	}
}
