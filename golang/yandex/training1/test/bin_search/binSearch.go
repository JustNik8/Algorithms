package main

import "fmt"

func main() {
	const SIZE = 10
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	params := checkParams{
		seq:    arr,
		number: 8,
	}

	fmt.Println("Left Bin Search:")
	leftAns := leftBinSearch(0, SIZE-1, leftCheck, params)
	fmt.Printf("Answer, index = %d number = %d\n", leftAns, arr[leftAns])

	fmt.Println("\n Right Bin Search:")
	rightAns := rightBinSearch(0, SIZE-1, rightCheck, params)
	fmt.Printf("Answer, index = %d number = %d\n", rightAns, arr[rightAns])
}

func leftCheck(m int, params checkParams) bool {
	return params.seq[m] >= params.number
}

func rightCheck(m int, params checkParams) bool {
	return params.seq[m] <= params.number
}

func leftBinSearch(l int, r int, check func(int, checkParams) bool, params checkParams) int {
	for l < r {
		m := (l + r) / 2
		fmt.Printf("Middle = %d\n", m)
		if check(m, params) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func rightBinSearch(l int, r int, check func(int, checkParams) bool, params checkParams) int {
	for l < r {
		m := (l + r + 1) / 2
		fmt.Printf("Middle = %d\n", m)
		if check(m, params) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

type checkParams struct {
	seq    []int
	number int
}
