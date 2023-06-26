package main

import "sort"

func main() {

}

func threeSum(nums []int) [][]int {
	size := len(nums)

	sort.Ints(nums)
	ans := make([][]int, 0)

	for p1 := 0; p1 < size-2; p1++ {
		if p1 > 0 && nums[p1] == nums[p1-1] {
			continue
		}

		p2 := p1 + 1
		p3 := size - 1

		p1Num := nums[p1]
		target := -p1Num

		for p2 < p3 {
			p2Num := nums[p2]
			p3Num := nums[p3]

			if p2Num+p3Num == target {
				ans = append(ans, []int{p1Num, p2Num, p3Num})
				p2++

				for p2 < p3 && nums[p2] == nums[p2-1] {
					p2++
				}
			} else if p2Num+p3Num < target {
				p2++
			} else {
				p3--
			}
		}
	}

	return ans
}
