package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) <= 3 {
		return [][]int{}
	}

	ans := make([][]int, 0)
	quad := make([]int, 0)
	sort.Ints(nums)

	var kSum func(k, start, localTarget int)
	kSum = func(k, start, localTarget int) {
		if k != 2 {
			for i := start; i < len(nums)-k+1; i++ {
				if i > start && nums[i] == nums[i-1] {
					continue
				}
				quad = append(quad, nums[i])
				kSum(k-1, i+1, localTarget-nums[i])
				quad = quad[:len(quad)-1]
			}
			return
		}

		l, r := start, len(nums)-1
		for l < r {
			if nums[l]+nums[r] < localTarget {
				l++
			} else if nums[l]+nums[r] > localTarget {
				r--
			} else {
				t := make([]int, 0)
				t = append(t, quad...)
				t = append(t, nums[l], nums[r])
				ans = append(ans, t)
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	kSum(4, 0, target)
	return ans
}
