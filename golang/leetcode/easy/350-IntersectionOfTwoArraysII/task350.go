package main

func intersect(nums1 []int, nums2 []int) []int {
	freqs := make(map[int]int)
	ans := make([]int, 0)

	for _, num := range nums1 {
		freqs[num] += 1
	}

	for _, num := range nums2 {
		if freqs[num] > 0 {
			ans = append(ans, num)
			freqs[num]--
		}
	}

	return ans
}
