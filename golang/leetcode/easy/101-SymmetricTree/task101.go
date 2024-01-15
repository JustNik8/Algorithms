package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}

	return l.Val == r.Val && helper(l.Left, r.Right) && helper(l.Right, r.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return helper(root.Left, root.Right)
}
