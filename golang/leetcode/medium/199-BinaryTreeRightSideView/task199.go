package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	queue := make([]*TreeNode, 0)

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) != 0 {
		size := len(queue)
		var node *TreeNode

		for i := 0; i < size; i++ {
			node = queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, node.Val)
	}

	return ans
}
