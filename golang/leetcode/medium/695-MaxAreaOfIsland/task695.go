package main

func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= rows {
			return 0
		}
		if j < 0 || j >= cols {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0
		cnt := 0

		cnt += dfs(i-1, j)
		cnt += dfs(i+1, j)
		cnt += dfs(i, j-1)
		cnt += dfs(i, j+1)

		return cnt + 1
	}

	maxArea := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
