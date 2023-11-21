package main

func singleNumber(nums []int) int {
	n := 0
	for _, num := range nums {
		n ^= num
	}

	return n
}
