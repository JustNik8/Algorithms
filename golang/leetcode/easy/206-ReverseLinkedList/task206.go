package main

import "fmt"

func main() {
	node3 := &ListNode{Val: 3, Next: nil}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	ans := reverseList(node1)

	for ans != nil {
		fmt.Println(ans)
		ans = ans.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		temp := curr.Next

		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}
