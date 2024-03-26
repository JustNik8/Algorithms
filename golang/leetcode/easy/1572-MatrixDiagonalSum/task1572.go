package main

func diagonalSum(mat [][]int) int {
	sum := 0
	size := len(mat)

	for i := 0; i < size; i++ {
		firstPos, secondPos := i, size-i-1

		sum += mat[i][firstPos]
		if firstPos != secondPos {
			sum += mat[i][secondPos]
		}
	}

	return sum
}
