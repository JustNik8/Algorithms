package main

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)

	var dfs func(index, total int, cur []int)
	dfs = func(index, total int, cur []int) {
		if total == target {
			dst := make([]int, len(cur))
			copy(dst, cur)
			ans = append(ans, dst)
			return
		}
		if index >= len(candidates) || total > target {
			return
		}

		num := candidates[index]

		cur = append(cur, num)
		dfs(index, total+num, cur)

		cur = cur[:len(cur)-1]
		dfs(index+1, total, cur)
	}

	dfs(0, 0, []int{})

	return ans
}
