package main

type empty struct{}

type set map[int]empty

// Space - O(n+m), Time - O(n*m)
func setZeroes(matrix [][]int) {
	zeroRows := make(set)
	zeroCols := make(set)

	n := len(matrix)
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				zeroRows[i] = empty{}
				zeroCols[j] = empty{}
			}
		}
	}

	for row := range zeroRows {
		for i := 0; i < m; i++ {
			matrix[row][i] = 0
		}
	}

	for col := range zeroCols {
		for i := 0; i < n; i++ {
			matrix[i][col] = 0
		}
	}
}

// Space - O(1), Time - O(n*m)
func setZeroes2(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])

	isFirstRowZero := false
	isFirstColZero := false

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != 0 {
				continue
			}

			if i == 0 {
				isFirstRowZero = true
			}
			if j == 0 {
				isFirstColZero = true
			}

			matrix[i][0] = 0
			matrix[0][j] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if isFirstRowZero {
		for r := 0; r < m; r++ {
			matrix[0][r] = 0
		}
	}

	if isFirstColZero {
		for c := 0; c < n; c++ {
			matrix[c][0] = 0
		}
	}
}
