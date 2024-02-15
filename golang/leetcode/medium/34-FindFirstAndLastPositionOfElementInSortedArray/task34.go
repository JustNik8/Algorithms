package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	ans := make([]int, 2)
	ans[0], ans[1] = -1, -1

	size := len(nums)
	l, r := 0, size-1

	for l <= r {
		m := (l + r) / 2
		fmt.Println(l, r, m)

		if (m == 0 && nums[m] == target) ||
			(nums[m] == target && nums[m-1] != target) {
			ans[0] = m
			break
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if ans[0] == -1 {
		return ans
	}

	l, r = ans[0], size-1
	for l <= r {
		m := (l + r) / 2

		if (m == size-1 && nums[m] == target) ||
			(nums[m] == target && nums[m+1] != target) {
			ans[1] = m
			break
		} else if nums[m] == target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return ans
}
