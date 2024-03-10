package main

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)

	for _, n := range nums1 {
		if _, exists := m[n]; !exists {
			m[n] = true
		}
	}

	ans := make([]int, 0)
	for _, n := range nums2 {
		notUsed := m[n]
		if notUsed {
			ans = append(ans, n)
			m[n] = false
		}
	}

	return ans
}
