package main

import (
	"container/heap"
	"fmt"
)

func main() {
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	minHeap *IntHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	h := IntHeap(nums)
	heap.Init(&h)

	for len(h) > k {
		heap.Pop(&h)
	}

	return KthLargest{minHeap: &h, k: k}
}

func (this *KthLargest) Add(val int) int {
	fmt.Println("Before add", val, ":", this.minHeap)
	heap.Push(this.minHeap, val)
	fmt.Println("Middle add", val, ":", this.minHeap)

	if len(*this.minHeap) > this.k {
		heap.Pop(this.minHeap)
	}
	fmt.Println("After add", val, ":", this.minHeap)
	fmt.Println()
	return (*this.minHeap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
