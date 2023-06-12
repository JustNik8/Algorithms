package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/add-two-numbers/?envType=list&envId=oceyii8d
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	current := ans

	cnt := 0

	for l1 != nil || l2 != nil {
		var first, second int
		var nextL1, nextL2 *ListNode

		if l1 != nil {
			first = l1.Val
			nextL1 = l1.Next
		}
		if l2 != nil {
			second = l2.Val
			nextL2 = l2.Next
		}

		sum := first + second + cnt

		if sum > 9 {
			cnt = 1
		} else {
			cnt = 0
		}

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next

		l1 = nextL1
		l2 = nextL2
	}

	if cnt == 1 {
		current.Next = &ListNode{Val: 1}
	}

	return ans.Next
}
