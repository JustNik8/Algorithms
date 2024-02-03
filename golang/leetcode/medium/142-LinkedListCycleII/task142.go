package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	slow := dummy
	fast := dummy

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			slow = dummy
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}

			return slow
		}
	}

	return nil
}
