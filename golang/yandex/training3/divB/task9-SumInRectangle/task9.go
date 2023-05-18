package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/9/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, m)
		mat[i] = row
		var num int
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &num)
			mat[i][j] = num
		}
	}

	sum := preprocess(mat)
	for i := 0; i < k; i++ {
		var leftX, leftY, rightX, rightY int
		fmt.Fscan(reader, &leftX, &leftY, &rightX, &rightY)
		fmt.Println(findSubMatrixSum(sum, leftX-1, leftY-1, rightX-1, rightY-1))
	}

}

func findSubMatrixSum(sum [][]int, leftX int, leftY int, rightX int, rightY int) int {

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
