package main

func subarraysDivByK(nums []int, k int) int {
	pref := make(map[int]int)
	pref[0] = 1

	sum := 0
	ans := 0
	for i := range nums {
		sum += nums[i]

		rem := sum % k
		if rem < 0 {
			rem += k
		}

		ans += pref[rem]

		pref[rem] += 1
	}

	return ans
}
