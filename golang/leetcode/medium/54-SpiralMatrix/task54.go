package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	n, m := len(matrix), len(matrix[0])

	x, y, dx, dy := 0, 0, 1, 0

	ans := make([]int, 0, len(matrix)*len(matrix[0]))
	for i := 0; i < n*m; i++ {
		fmt.Println("YX:", y, dy, x, dx)
		if x+dx > right {
			dx, dy = 0, 1
			top++
		} else if x+dx < left {
			dx, dy = 0, -1
			bottom--
		} else if y+dy > bottom {
			dx, dy = -1, 0
			right--
		} else if y+dy < top {
			dx, dy = 1, 0
			left++
		}

		ans = append(ans, matrix[y][x])
		fmt.Println(ans)
		x += dx
		y += dy
	}

	return ans
}
