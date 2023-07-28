package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree1(root *TreeNode) *TreeNode {
	dfs(root)
	return root
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left)
	dfs(root.Right)

	left, right := root.Left, root.Right

	root.Right = left
	root.Left = right
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left, right := root.Left, root.Right

	root.Left = right
	root.Right = left

	invertTree2(root.Left)
	invertTree2(root.Right)

	return root
}
