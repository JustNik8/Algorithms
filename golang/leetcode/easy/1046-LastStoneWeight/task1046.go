package main

import "container/heap"

func main() {

}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func lastStoneWeight(stones []int) int {
	tmp := IntHeap(stones)
	maxHeap := &tmp
	heap.Init(maxHeap)

	for len(*maxHeap) > 1 {
		firstStone := heap.Pop(maxHeap).(int)
		secondStone := heap.Pop(maxHeap).(int)

		diff := firstStone - secondStone

		if diff > 0 {
			heap.Push(maxHeap, diff)
		}
	}

	if len(*maxHeap) > 0 {
		return heap.Pop(maxHeap).(int)
	} else {
		return 0
	}

}
