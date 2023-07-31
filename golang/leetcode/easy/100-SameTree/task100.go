package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	isSameLeft := isSameTree(p.Left, q.Left)
	isSameRight := isSameTree(p.Right, q.Right)

	return isSameLeft && isSameRight
}
