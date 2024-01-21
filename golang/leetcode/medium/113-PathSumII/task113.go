package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	curNums := make([]int, 0)
	curSum := 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		curNums = append(curNums, node.Val)
		curSum += node.Val
		if node.Left == nil && node.Right == nil {
			if curSum == targetSum {
				dst := make([]int, len(curNums))
				copy(dst, curNums)
				ans = append(ans, dst)
			}
		}

		dfs(node.Left)
		dfs(node.Right)

		curNums = curNums[:len(curNums)-1]
		curSum -= node.Val
	}

	dfs(root)
	return ans
}
