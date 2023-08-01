package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if max != nil && max.Val <= root.Val {
		return false
	}
	if min != nil && min.Val >= root.Val {
		return false
	}

	isLeftValid := isValid(root.Left, min, root)
	isRightValid := isValid(root.Right, root, max)

	return isLeftValid && isRightValid
}
