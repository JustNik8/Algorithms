package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	ans := -1
	counter := 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil || counter == k {
			return
		}

		dfs(root.Left)
		counter++
		if counter == k {
			ans = root.Val
		}
		dfs(root.Right)

	}

	dfs(root)
	return ans
}
