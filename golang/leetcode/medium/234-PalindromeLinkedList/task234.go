package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 - 2 - 2 - 2 - 1
// Time - O(n), Space - O(n)
func isPalindromeMem(head *ListNode) bool {
	size := 0

	node := head

	for node != nil {
		size++
		node = node.Next
	}

	leftSize := size / 2
	skip := size%2 == 1
	left := make([]int, leftSize)

	node = head
	for i := 0; i < leftSize; i++ {
		left[i] = node.Val
		node = node.Next
	}

	if skip {
		node = node.Next
	}

	for i := 0; i < leftSize; i++ {
		if left[leftSize-i-1] != node.Val {
			return false
		}
		node = node.Next
	}

	return true
}

//	s         f
//
// 1 -> 2 -> 2 <- 2 <- 1
//
//	s         f
//
// 1 -> 2 <- 2 <- 1
// Time - O(n), Space - O(1)
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	slow, fast := head, head

	// Поиск середины
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Развернуть вторую половину листа
	var prev *ListNode = nil
	for slow != nil {
		temp := slow.Next

		slow.Next = prev
		prev = slow
		slow = temp
	}

	// Проверка на палиндром
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}

		left = left.Next
		right = right.Next
	}

	return true
}
