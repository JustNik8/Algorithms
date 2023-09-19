package main

func main() {

}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	subset := make([]int, 0)

	var dfs func(index int)
	dfs = func(index int) {
		if index >= len(nums) {
			dst := make([]int, len(subset))
			copy(dst, subset)

			ans = append(ans, dst)
			return
		}
		subset = append(subset, nums[index])
		dfs(index + 1)

		subset = subset[:len(subset)-1]
		dfs(index + 1)
	}

	dfs(0)
	return ans
}
