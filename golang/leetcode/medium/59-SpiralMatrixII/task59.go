package main

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	left, right := 0, n-1
	top, bottom := 0, n-1

	dx, dy := 1, 0
	x, y := 0, 0
	counter := 1

	for i := 0; i < n*n; i++ {
		if x+dx > right {
			dx, dy = 0, 1
			top++
		} else if y+dy > bottom {
			dx, dy = -1, 0
			right--
		} else if x+dx < left {
			dx, dy = 0, -1
			bottom--
		} else if y+dy < top {
			dx, dy = 1, 0
			left++
		}

		ans[y][x] = counter
		counter++

		x += dx
		y += dy
	}

	return ans
}
