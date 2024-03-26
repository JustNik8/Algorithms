package main

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target || nums[l] == target {
			return true
		}

		if nums[l] < nums[m] {
			// Left portion
			if target < nums[l] || target > nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else if nums[r] > nums[m] {
			// Right portion
			if target > nums[r] || target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			l++
		}
	}

	return false
}
