package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			res = append(res, "n")
			return
		}

		numStr := strconv.Itoa(node.Val)
		res = append(res, numStr)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return strings.Join(res, ",")
}

//   2
//  / \
// 1   3

// 2 1 n n 3 n n
// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	fmt.Println(data)
	values := strings.Split(data, ",")

	i := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if i >= len(values) {
			return nil
		}
		if values[i] == "n" {
			i++
			return nil
		}

		num, _ := strconv.Atoi(values[i])
		node := &TreeNode{Val: num}

		i++
		node.Left = dfs()
		node.Right = dfs()

		return node
	}

	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
