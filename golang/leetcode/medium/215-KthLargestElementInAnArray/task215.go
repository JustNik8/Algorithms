package main

import "container/heap"

func main() {
}

// Solution #1. Using heap.

type IntHeap []int

func (h IntHeap) Len() int           { return (len(h)) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	maxHeap := IntHeap(nums)
	heap.Init(&maxHeap)

	for i := 0; i < k-1; i++ {
		heap.Pop(&maxHeap)
	}

	return (heap.Pop(&maxHeap)).(int)
}
