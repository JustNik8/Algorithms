package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, num int) int {
	if node == nil {
		return 0
	}

	num = 10*num + node.Val
	if node.Left == nil && node.Right == nil {
		return num
	}

	left := dfs(node.Left, num)
	right := dfs(node.Right, num)

	return left + right
}
