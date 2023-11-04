package main

const (
	fresh  = 1
	rotten = 2
)

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	q := make([][2]int, 0)
	freshCnt := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == fresh {
				freshCnt++
			} else if grid[r][c] == rotten {
				q = append(q, [2]int{r, c})
			}
		}
	}

	time := 0
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for freshCnt > 0 && len(q) != 0 {
		rottenCnt := len(q)

		for i := 0; i < rottenCnt; i++ {
			cell := q[0]
			q = q[1:]

			for _, d := range directions {
				r, c := cell[0]+d[0], cell[1]+d[1]

				if r >= 0 && c >= 0 && r < rows && c < cols &&
					grid[r][c] == fresh {
					grid[r][c] = rotten
					freshCnt--

					q = append(q, [2]int{r, c})
				}
			}
		}
		time++
	}

	if freshCnt == 0 {
		return time
	}
	return -1
}
