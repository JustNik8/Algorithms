package main

import (
	"strconv"
	"strings"
)

func main() {

}

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
	nodes := make([]string, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			nodes = append(nodes, "n")
			return
		}

		nodes = append(nodes, strconv.Itoa(node.Val))

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return strings.Join(nodes, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")

	index := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if vals[index] == "n" {
			index++
			return nil
		}

		val, _ := strconv.Atoi(vals[index])
		node := &TreeNode{Val: val}
		index++

		node.Left = dfs()
		node.Right = dfs()

		return node
	}

	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
