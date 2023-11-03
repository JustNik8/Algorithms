package main

func main() {

}

// Более удачное решение. Нет доп. памяти на посещенные вершины.
// в dfs краевые случаи описаны
func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows {
			return
		}
		if j < 0 || j >= cols {
			return
		}
		if grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)

	}

	cnt := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				cnt++
			}
		}
	}

	return cnt
}

// Мое первое решение. Здесь острова и посещенные острова хранятся в разных слайсах
func numIslandsBad(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	visit := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visit[i] = make([]bool, cols)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		visit[i][j] = true

		left := j - 1
		right := j + 1
		top := i - 1
		bottom := i + 1

		if left >= 0 && grid[i][left] == '1' && !visit[i][left] {
			dfs(i, left)
		}
		if right < cols && grid[i][right] == '1' && !visit[i][right] {
			dfs(i, right)
		}
		if top >= 0 && grid[top][j] == '1' && !visit[top][j] {
			dfs(top, j)
		}
		if bottom < rows && grid[bottom][j] == '1' && !visit[bottom][j] {
			dfs(bottom, j)
		}

	}

	cnt := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visit[i][j] {
				dfs(i, j)
				cnt++
			}
		}
	}

	return cnt
}
