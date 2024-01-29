package main

func removeDuplicates(nums []int) int {
	ptr := 0
	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[ptr] {
			continue
		}

		k++
		ptr++
		nums[ptr] = nums[i]
	}

	return k
}
