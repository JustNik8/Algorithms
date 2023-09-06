package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return pathFunc(root, 0, targetSum)
}

func pathFunc(node *TreeNode, prevSum, targetSum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		return prevSum+node.Val == targetSum
	}

	left := pathFunc(node.Left, prevSum+node.Val, targetSum)
	right := pathFunc(node.Right, prevSum+node.Val, targetSum)

	return left || right
}
