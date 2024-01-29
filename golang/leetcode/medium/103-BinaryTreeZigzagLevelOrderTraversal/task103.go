package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := make([]*TreeNode, 0)
	ans := make([][]int, 0)

	level := 0
	q = append(q, root)

	for len(q) != 0 {
		size := len(q)
		ans = append(ans, make([]int, size))

		for i := 0; i < size; i++ {
			node := q[0]

			if level%2 == 0 {
				ans[level][i] = node.Val
			} else {
				ans[level][size-i-1] = node.Val
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

			q = q[1:]
		}

		level++
	}

	return ans
}
