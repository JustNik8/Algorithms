package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	cnt := 0

	var dfs func(root *TreeNode, maxNum int)
	dfs = func(root *TreeNode, maxNum int) {
		if root == nil {
			return
		}

		if root.Val >= maxNum {
			cnt++
		}

		dfs(root.Left, max(root.Val, maxNum))
		dfs(root.Right, max(root.Val, maxNum))
	}

	dfs(root, root.Val)
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
