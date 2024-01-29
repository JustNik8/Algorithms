package main

func majorityElement(nums []int) int {
	m := make(map[int]int)
	size := len(nums) / 2

	for _, num := range nums {
		m[num] += 1

		if m[num] > size {
			return num
		}
	}

	return 0
}
