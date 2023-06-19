package main

type empty struct{}

func containsDuplicate(nums []int) bool {
	set := make(map[int]empty)

	for _, num := range nums {
		_, contains := set[num]

		if contains {
			return true
		}

		set[num] = empty{}
	}

	return false
}
