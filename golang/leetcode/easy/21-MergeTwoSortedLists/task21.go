package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	node1 := list1
	node2 := list2

	ans := new(ListNode)
	tail := ans

	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			tail.Next = node1
			node1 = node1.Next
		} else {
			tail.Next = node2
			node2 = node2.Next
		}
		tail = tail.Next
	}

	if node1 != nil {
		tail.Next = node1
	} else {
		tail.Next = node2
	}

	return ans.Next
}
