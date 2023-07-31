package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDiameter = 0

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0

	dfs(root)

	return maxDiameter
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	if left+right > maxDiameter {
		maxDiameter = left + right
	}

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
