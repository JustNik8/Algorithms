package main

import "container/heap"

type info struct {
	word string
	freq int
}

type InfoHeap []info

func (h InfoHeap) Len() int {
	return len(h)
}

func (h InfoHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].word < h[j].word
	}
	return h[i].freq > h[j].freq
}

func (h InfoHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *InfoHeap) Push(x interface{}) {
	*h = append(*h, x.(info))
}

func (h *InfoHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	count := make(map[string]int)
	for _, w := range words {
		count[w] += 1
	}

	h := &InfoHeap{}

	heap.Init(h)

	for w, c := range count {
		heap.Push(h, info{word: w, freq: c})
	}

	ans := make([]string, k)
	for i := 0; i < k; i++ {
		ans[i] = heap.Pop(h).(info).word
	}

	return ans

}
