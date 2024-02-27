package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		cnt := 0
		if node == p || node == q {
			cnt++
		}

		cnt += dfs(node.Left)
		cnt += dfs(node.Right)

		if cnt == 2 && lca == nil {
			lca = node
		}

		return cnt
	}

	dfs(root)
	return lca
}
