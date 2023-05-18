package main

import "fmt"

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	sum := findSubMatrixSum(mat, 0, 0, 1, 2)
	fmt.Println(sum)
}

func preprocess(mat [][]int) [][]int {
	// MxN matrix
	m := len(mat)
	n := len(mat[0])

	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
	}

	sum[0][0] = mat[0][0]

	// preprocess the first row
	for j := 1; j < n; j++ {
		sum[0][j] = mat[0][j] + sum[0][j-1]
	}

	// preprocess the first column
	for i := 1; i < m; i++ {
		sum[i][0] = mat[i][0] + sum[i-1][0]
	}

	// preprocess the rest of the matrix
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sum[i][j] = mat[i][j] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}
	return sum
}

// Calculate the sum of all elements in a submatrix in constant time
func findSubMatrixSum(mat [][]int, leftX int, leftY int, rightX int, rightY int) int {
	//base case
	if mat == nil || len(mat) == 0 {
		return 0
	}

	//preprocess the matrix
	sum := preprocess(mat)

	// total = sum[rightX][rightY] - sum[rightX][leftY-1] - sum[leftX-1][rightY] + sum[leftX-1][leftY-1]
	total := sum[rightX][rightY]
	if leftY-1 >= 0 {
		total -= sum[rightX][leftY-1]
	}
	if leftX-1 >= 0 {
		total -= sum[leftX-1][rightY]
	}
	if leftX-1 >= 0 && leftY-1 >= 0 {
		total += sum[leftX-1][leftY-1]
	}
	return total
}
