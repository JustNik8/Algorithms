package main

type HitCounter struct {
	q []int
}

func NewHitCounter() *HitCounter {
	return &HitCounter{q: make([]int, 0)}
}

func (h *HitCounter) Hit(timestamp int) {
	h.q = append(h.q, timestamp)
}

func (h *HitCounter) GetHits(timestamp int) int {
	for len(h.q) > 0 && timestamp-h.q[0] >= 300 {
		h.q = h.q[1:]
	}

	return len(h.q)
}
