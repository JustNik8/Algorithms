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
