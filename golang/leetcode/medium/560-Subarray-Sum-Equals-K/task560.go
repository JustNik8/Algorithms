package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{-1, -1, 1}, 0))
}

func subarraySum(nums []int, k int) int {
	size := len(nums)
	pref := make(map[int]int)
	ans := 0

	pref[0] = 1
	currSum := 0
	for i := 0; i < size; i++ {
		currSum += nums[i]

		diff := currSum - k

		ans += pref[diff]
		pref[currSum] += 1
	}

	return ans
}
