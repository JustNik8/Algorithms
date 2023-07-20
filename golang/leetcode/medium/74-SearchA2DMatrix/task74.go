package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	ROWS, COLS := len(matrix), len(matrix[0])

	top, bottom := 0, ROWS-1

	for top <= bottom {
		row := (top + bottom) / 2
		if target > matrix[row][COLS-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bottom = row - 1
		} else {
			break
		}
	}

	if top > bottom {
		return false
	}

	row := (top + bottom) / 2
	left, right := 0, COLS-1

	for left <= right {
		m := (left + right) / 2
		if target > matrix[row][m] {
			left = m + 1
		} else if target < matrix[row][m] {
			right = m - 1
		} else {
			return true
		}
	}

	return false
}
