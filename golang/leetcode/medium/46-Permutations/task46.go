package main

func main() {

}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)

	if len(nums) == 1 {
		dst := make([]int, 1)
		copy(dst, nums)
		return [][]int{dst}
	}

	for i := 0; i < len(nums); i++ {
		n := nums[0]
		nums = nums[1:]

		perms := permute(nums)
		for j := 0; j < len(perms); j++ {
			perms[j] = append(perms[j], n)
			ans = append(ans, perms[j])
		}
		nums = append(nums, n)

	}

	return ans
}
