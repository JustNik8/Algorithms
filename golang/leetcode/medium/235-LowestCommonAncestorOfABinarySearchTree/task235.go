package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// This solution much better. Uses features of BST
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		cnt := 0
		if root == p || root == q {
			cnt++
		}

		cnt += dfs(root.Left)
		cnt += dfs(root.Right)

		if cnt == 2 && lca == nil {
			lca = root
		}
		return cnt
	}

	dfs(root)

	return lca
}
