package main

import "container/heap"

func main() {

}

type Point struct {
	index int
	dist  int
}

type PointHeap []Point

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PointHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	tmp := PointHeap{}
	minHeap := &tmp
	heap.Init(minHeap)

	for i := 0; i < len(points); i++ {
		point := points[i]
		dist := point[0]*point[0] + point[1]*point[1]

		heap.Push(minHeap, Point{index: i, dist: dist})
	}

	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		point := (heap.Pop(minHeap)).(Point)
		ans[i] = points[point.index]
	}

	return ans
}
