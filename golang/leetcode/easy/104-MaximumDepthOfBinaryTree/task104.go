package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	deque := make([]*TreeNode, 0)
	deque = append(deque, root)

	for len(deque) != 0 {
		size := len(deque)
		for i := 0; i < size; i++ {
			node := deque[0]
			deque = deque[1:]

			if node.Left != nil {
				deque = append(deque, node.Left)
			}
			if node.Right != nil {
				deque = append(deque, node.Right)
			}
		}
		level++
	}

	return level
}
