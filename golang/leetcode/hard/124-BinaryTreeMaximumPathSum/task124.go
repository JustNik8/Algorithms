package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	pathSum(root, &maxSum)

	return maxSum
}

func pathSum(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	leftPath := max(0, pathSum(node.Left, maxSum))
	rightPath := max(0, pathSum(node.Right, maxSum))

	curSum := node.Val + leftPath + rightPath

	*maxSum = max(*maxSum, curSum)

	return node.Val + max(leftPath, rightPath)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
