package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closest := nums[0] + nums[1] + nums[2]
	for p1 := 0; p1 < len(nums)-2; p1++ {
		p2 := p1 + 1
		p3 := len(nums) - 1

		for p2 < p3 {
			delta := abs(target - (nums[p1] + nums[p2] + nums[p3]))
			if abs(target-closest) > delta {
				closest = nums[p1] + nums[p2] + nums[p3]
				if closest == target {
					return closest
				}
			}

			if nums[p1]+nums[p2]+nums[p3] < target {
				p2++
			} else {
				p3--
			}
		}
	}

	return closest
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
