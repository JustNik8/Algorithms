package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}
func (h *NodeHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy

	h := make(NodeHeap, 0, len(lists))

	for _, node := range lists {
		if node != nil {
			h = append(h, node)
		}
	}

	heap.Init(&h)

	for h.Len() > 0 {
		node := (heap.Pop(&h)).(*ListNode)

		if node.Next != nil {
			heap.Push(&h, node.Next)
		}

		p.Next = node
		p = p.Next
	}

	return dummy.Next
}
