package main

import "fmt"

// https://leetcode.com/problems/maximize-distance-to-closest-person/
func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}))
}

func maxDistToClosest(seats []int) int {
	size := len(seats)
	left := -1
	right := 0

	maxDist := 0
	for right < size {
		if seats[right] == 1 {
			if left == -1 {
				maxDist = right
			} else if (right-left)/2 > maxDist {
				maxDist = (right - left) / 2
			}
			left = right
		}

		right++
	}
	if size-left-1 > maxDist {
		maxDist = size - left - 1
	}

	return maxDist
}
