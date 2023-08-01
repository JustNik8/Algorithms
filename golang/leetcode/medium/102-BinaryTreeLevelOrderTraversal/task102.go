package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	queue := make([]*TreeNode, 0)

	if root != nil {
		queue = append(queue, root)
	}

	level := 0
	for len(queue) != 0 {
		size := len(queue)
		ans = append(ans, make([]int, size))

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			ans[level][i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		level++
	}

	return ans
}
